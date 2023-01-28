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

func (r *OrderRepo) Insert(ctx context.Context, req *models.CreateOrder) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
				INSERT INTO "order" (
						id,
						car_id,
						client_id,
						day_count,
						updated_at
				) VALUES ($1, $2, $3, $4, now())
			`
	_, err := r.db.Exec(ctx, query,
		id,
		req.CarId,
		req.ClientId,
		req.DayCount,
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
		dayCount   sql.NullFloat64
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
	resp.DayCount = dayCount.Float64
	resp.GiveKm = giveKm.Float64
	resp.ReceiveKm = receiveKm.Float64
	resp.Status = status.String
	resp.CreatedAt = createdAt.String
	resp.UpdatedAt = updatedAt.String

	return resp, err
}

func (r *OrderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {

	var (
		resp   models.GetListOrderResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

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

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return &models.GetListOrderResponse{}, err
	}

	var (
		id          sql.NullString
		car_id      sql.NullString
		client_id   sql.NullString
		total_price sql.NullFloat64
		paid_price  sql.NullFloat64
		day_count   sql.NullFloat64
		give_km     sql.NullFloat64
		receive_km  sql.NullFloat64
		status      sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	for rows.Next() {

		err = rows.Scan(
			&resp.Count,
			&id,
			&car_id,
			&client_id,
			&total_price,
			&paid_price,
			&day_count,
			&give_km,
			&receive_km,
			&status,
			&createdAt,
			&updatedAt,
		)

		order := models.Order{
			Id:          id.String,
			CarId:      car_id.String,
			ClientId:   client_id.String,
			TotalPrice: total_price.Float64,
			PaidPrice:  paid_price.Float64,
			DayCount:   day_count.Float64,
			GiveKm:     give_km.Float64,
			ReceiveKm:  receive_km.Float64,
			Status:      status.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		}

		if err != nil {
			return &models.GetListOrderResponse{}, err
		}
		resp.Orders = append(resp.Orders, &order)

	}
	return &resp, nil
}

func (r *OrderRepo) Update(ctx context.Context, order *models.UpdateOrder) error {
	query := `
		UPDATE 
			"order" 
		SET 
			car_id = $2,
			client_id = $3,
			paid_price = $4,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		order.Id,
		order.CarId,
		order.ClientId,
		order.PaidPrice,
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

