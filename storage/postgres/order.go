package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"app/models"
	"app/pkg/helper"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) Insert(ctx context.Context, order *models.CreateOrder) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO "order" (
			id,
			car_id,
			client_id,
			total_price,
			day_count,
			updated_at
		) VALUES ($1, $2, $3, (
			SELECT
				price * $4
			FROM car
			WHERE id = $2
		), $4, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		order.CarId,
		order.ClientId,
		order.DayCount,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *OrderRepo) GetByID(ctx context.Context, req *models.OrderPrimeryKey) (*models.Order, error) {

	var (
		resp       = &models.Order{}
		id         sql.NullString
		carId      sql.NullString
		clientId   sql.NullString
		totalPrice sql.NullFloat64
		paidPrice  sql.NullFloat64
		dayCount   sql.NullInt64
		giveKm     sql.NullFloat64
		receiveKm  sql.NullFloat64
		status     sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	query := `
		SELECT
			o.id,
			o.car_id,

			COALESCE(c.id::varchar, ''),
			COALESCE(c.state_number, ''),
			COALESCE(c.model, ''),

			COALESCE(cl.first_name || ' ' || cl.last_name, ''),

			o.client_id,
			o.total_price,
			o.paid_price,
			o.day_count,
			o.give_km,
			o.receive_km,
			o.status,
			o.created_at,
			o.updated_at
		FROM "order" AS o
		JOIN car AS c ON c.id = o.car_id
		JOIN client AS cl ON cl.id = o.client_id
		WHERE o.id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&carId,

		&resp.Car.Id,
		&resp.Car.StateNumber,
		&resp.Car.Model,

		&resp.ClientFullName,
		&clientId,
		&totalPrice,
		&paidPrice,
		&dayCount,
		&giveKm,
		&receiveKm,
		&status,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	resp.Id = id.String
	resp.CarId = carId.String
	resp.ClientId = clientId.String
	resp.TotalPrice = totalPrice.Float64
	resp.PaidPrice = paidPrice.Float64
	resp.DayCount = int(dayCount.Int64)
	resp.GiveKm = int(giveKm.Float64)
	resp.ReceiveKm = int(receiveKm.Float64)
	resp.Status = status.String
	resp.CreatedAt = createdAt.String
	resp.UpdatedAt = updatedAt.String

	return resp, err
}

func (r *OrderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {
	var (
		offset = "OFFSET 0"
		limit  = "LIMIT 10"
		resp   = &models.GetListOrderResponse{}
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf("OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			car_id,
			client_id,
			total_price,
			paid_price,
			day_count,
			give_km,
			receive_km,
			status,
			created_at,
			updated_at
		FROM "order"
	`

	query += offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var (
			id         sql.NullString
			carId      sql.NullString
			clientId   sql.NullString
			totalPrice sql.NullFloat64
			paidPrice  sql.NullFloat64
			dayCount   sql.NullInt64
			giveKm     sql.NullInt64
			receiveKm  sql.NullInt64
			status     sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&id,
			&carId,
			&clientId,
			&totalPrice,
			&paidPrice,
			&dayCount,
			&giveKm,
			&receiveKm,
			&status,
			&createdAt,
			&updatedAt,
		)

		resp.Orders = append(resp.Orders, &models.Order{
			Id:         id.String,
			CarId:      carId.String,
			ClientId:   clientId.String,
			TotalPrice: totalPrice.Float64,
			PaidPrice:  paidPrice.Float64,
			DayCount:   int(dayCount.Int64),
			GiveKm:     int(giveKm.Int64),
			ReceiveKm:  int(receiveKm.Int64),
			Status:     status.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return resp, err
}

func (r *OrderRepo) Update(ctx context.Context, order *models.UpdateOrder) error {

	query := `
		UPDATE
			"order"
		SET
			car_id = $2,
			client_id = $3,
			total_price = $4,
			paid_price = $5,
			day_count = $6,
			give_km = $7,
			receive_km = $8,
			status = $9
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		order.Id,
		order.CarId,
		order.ClientId,
		order.TotalPrice,
		order.PaidPrice,
		order.DayCount,
		order.GiveKm,
		order.ReceiveKm,
		order.Status,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatch) error {

	var (
		set   = " SET "
		ind   = 0
		query string
	)

	if len(req.Data) == 0 {
		return errors.New("no updates provided")
	}

	req.Data["id"] = req.Id

	for key := range req.Data {
		set += fmt.Sprintf(" %s = :%s ", key, key)
		if ind != len(req.Data)-1 {
			set += ", "
		}
		ind++
	}

	query = `
			UPDATE
				"order"
		` + set + ` , updated_at = now()
			WHERE
				id = :id
		`

	query, args := helper.ReplaceQueryParams(query, req.Data)

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) Delete(ctx context.Context, req *models.OrderPrimeryKey) error {

	query := `
	delete from "order" where id = $1
	`

	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
