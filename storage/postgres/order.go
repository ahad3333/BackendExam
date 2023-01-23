package postgres

import (
	"add/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
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
				INSERT INTO "Order" (
						id,
						car_id,
						client_id,
						day_count,
						updated_at
				) VALUES ($1, $2, $3, $4, now())
			`
	_, err := r.db.Exec(ctx, query,
		id,
		req.Car_id,
		req.Client_id,
		req.Day_count,
	)

	if err != nil {
		return "", err
	}

	var (
		carstatus   = req.Car_id
		km          = req.Car_id
	)

	kmUpdete := `
				UPDATE "Order" 
				SET    give_km = km
				FROM   car 
			`
	if km != "" {
		km = fmt.Sprintf(`WHERE  "Order".car_id =  '%s'%s`, req.Car_id, " AND  km IS DISTINCT FROM give_km ")
		kmUpdete += km
	}

	_, err = r.db.Exec(ctx, kmUpdete)
	if err != nil {
		return "", err
	}

	querycar := `
			UPDATE 
					car 
				SET
				status = 'booked'
	`
	if carstatus != "" {
		carstatus = fmt.Sprintf("where id =   '%s' ", req.Car_id)
		querycar += carstatus
	}

	_, err = r.db.Exec(ctx, querycar)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *OrderRepo) GetByID(ctx context.Context, req *models.OrderPrimeryKey) (*models.Order, error) {

	query := `
		SELECT
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
		FROM
			"Order"
		WHERE id = $1
	`

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

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
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

	if err != nil {
		return nil, err
	}

	order := &models.Order{
		Id:          id.String,
		Car_id:      car_id.String,
		Client_id:   client_id.String,
		Total_price: total_price.Float64,
		Paid_price:  paid_price.Float64,
		Day_count:   day_count.Float64,
		Give_km:     give_km.Float64,
		Receive_km:  receive_km.Float64,
		Status:      status.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}

	return order, nil
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
		FROM "Order"
	
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
			Car_id:      car_id.String,
			Client_id:   client_id.String,
			Total_price: total_price.Float64,
			Paid_price:  paid_price.Float64,
			Day_count:   day_count.Float64,
			Give_km:     give_km.Float64,
			Receive_km:  receive_km.Float64,
			Status:      status.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		}

		if err != nil {
			return &models.GetListOrderResponse{}, err
		}

		resp.Order = append(resp.Order, &order)

	}
	return &resp, nil
}

func (r *OrderRepo) Update(ctx context.Context, order *models.Order) error {
	query := `
		UPDATE 
			"Order" 
		SET 
			car_id = $2,
			client_id = $3,
			paid_price = $4,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		order.Id,
		order.Car_id,
		order.Client_id,
		order.Paid_price,
	)

	if err != nil {
		return err
	}

	total := `
		select 
			sum(o.day_count * c.price)
		FROM "Order" as o
		join car as c on c.id = o.car_id
		where c.id = $2 and o.id = $1
		`
		var totals sql.NullFloat64

		err = r.db.QueryRow(ctx,total,order.Id, order.Car_id).Scan(
			&totals,
		)

		order.Total_price = totals.Float64

		var(
			total_price = order.Total_price
		)

		queryTotal := `
						UPDATE 
								"Order" 
							SET
				`
		if total_price >0 {
			price := fmt.Sprintf(" total_price = %v  where id =   '%s' ",total_price, order.Id)
			queryTotal += price
		}

		_,err = r.db.Exec(ctx,queryTotal)
			if err != nil {
				return err
			}

	var(
		carstatus = order.Car_id
		orderstatus = order.Id
	)

	querycar := `
			UPDATE 
					car 
				SET
				status = 'in_use'
	`
	if carstatus != "" {
		carstatus = fmt.Sprintf("where id =   '%s' ", order.Car_id)
		querycar += carstatus
	}

	_, err = r.db.Exec(ctx, querycar)

	if err != nil {
		return  err
	}

	queryorder := `
			UPDATE 
					"Order" 
				SET
				status = 'client_took'
	`
	if orderstatus != "" {
		orderstatus = fmt.Sprintf("where id =   '%s' ", order.Id)
		queryorder += orderstatus
	}

	_, err = r.db.Exec(ctx, queryorder)

	if err != nil {
		return  err
	}

	return nil
}

func (r *OrderRepo) Delete(ctx context.Context, req *models.OrderPrimeryKey) error {

	_, err := r.db.Exec(ctx, `DELETE FROM "Order" WHERE id  = $1 `, req.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) Return(ctx context.Context, order *models.Order) error {
	query := `
		UPDATE 
			"Order" 
		SET 
			car_id = $2,
			client_id = $3,
			receive_km = $4,
			updated_at = now()
		WHERE id = $1
	`
	

	_, err := r.db.Exec(ctx, query,
		order.Id,
		order.Car_id,
		order.Client_id,
		order.Receive_km,
	)

	if err != nil {
		return err
	}

	var (
		investor models.InvestorBenefit
		client models.Bebtors
	)

		queryResult :=`
				select 
				i.name,
				sum(paid_price/100*70),
				cl.first_name, 
                o.total_price - paid_price
			FROM "Order" as o
			join Client as cl on cl.id = o.client_id
			join car as c on o.car_id = c.id
			join Investor as i on investor_id = c.investor_id
			where c.id = $2
			and cl.id = $3 and o.id = $1
			GROUP BY i.name, cl.first_name,o.total_price, paid_price
		`
		var(
			bebt sql.NullFloat64
		)

		err = r.db.QueryRow(ctx, queryResult,order.Id,order.Car_id,order.Client_id).Scan(
			&investor.Name,
			&investor.Price,
			&client.Name,
			&bebt,
		)

		client.Bebt = bebt.Float64

		if err != nil {
			return err
		}

		var(
			InsertName = investor.Name
		)

		queryBenefit :=`	
				INSERT INTO InvestorBenefit (
					name,
					price,
					updated_at
				) VALUES (
	`
		if InsertName != "" {
			InsertName = fmt.Sprintf(" '%s', %v, now() %s", investor.Name, investor.Price,");")
			queryBenefit += InsertName 
		}

		_,err = r.db.Exec(ctx,queryBenefit)

			if err != nil {
				return err
			}

			var(
				clientName = client.Name
			)

			queryBebtors :=`	
					INSERT INTO Bebtors (
						name,
						bebt,
						updated_at
					) VALUES (
		`
			if clientName != "" {
				clientName = fmt.Sprintf(" '%s', %v, now() %s", client.Name, client.Bebt,")")
				queryBebtors += clientName 
			}
			_,err = r.db.Exec(ctx,queryBebtors)

				if err != nil {
					return err
				}

				var(
					carstatus = order.Car_id
					orderstatus = order.Id
				)

				querycar := `
						UPDATE 
								car 
							SET
							status = 'in_stock'
				`
				if carstatus != "" {
					carstatus = fmt.Sprintf("where id =   '%s' ", order.Car_id)
					querycar += carstatus
				}
			
				_, err = r.db.Exec(ctx, querycar)
				if err != nil {
					return  err
				}
			
				queryorder := `
						UPDATE 
								"Order" 
							SET
							status = 'client_returned'
				`
				if orderstatus != "" {
					orderstatus = fmt.Sprintf("where id =   '%s' ", order.Id)
					queryorder += orderstatus
				}
			
				_, err = r.db.Exec(ctx, queryorder)

				if err != nil {
					return  err
				}

				var (
					km   = order.Car_id
				)

				kmUpdete := `
							UPDATE car 
							SET    km = receive_km
							FROM  "Order" 
						`
				if km != "" {
					km = fmt.Sprintf(`WHERE  car.id =  '%s'%s`, order.Car_id, " AND  km IS DISTINCT FROM receive_km")
					kmUpdete += km
				}
			
				_, err = r.db.Exec(ctx, kmUpdete)
				
				if err != nil {
					return  err
				}				
		
	return nil
}
