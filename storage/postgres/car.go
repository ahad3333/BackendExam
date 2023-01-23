package postgres

import (
	"add/models"
	"database/sql"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	
)

type CarRepo struct {
	db *pgxpool.Pool
}

func NewCarRepo(db *pgxpool.Pool) *CarRepo {
	return &CarRepo{
		db: db,
	}
}

func (r *CarRepo) Insert(ctx context.Context, req *models.CreateCar) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
	INSERT INTO Car (
			id,
			state_number,
			model,
			price,
			daily_limit,
			over_limit,
			investor_id,
			km,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())
		`

	_, err := r.db.Exec(ctx, query,
		id,
		req.State_number,
		req.Model,
		req.Price,
		req.Daily_limit,
		req.Over_limit,
		req.Investor_id,
		req.Km,
	)

	if err != nil {
		return "", err
	}


	return id, nil
}

func (r *CarRepo) GetByID(ctx context.Context, req *models.CarPrimeryKey) (*models.Car, error) {

	query := `
		SELECT
				id,
				state_number,
				model,
				status,
				price,
				daily_limit,
				over_limit,
				investor_percentage,
				investor_id,
				km,
				created_at,
				updated_at
		FROM
			Car
		WHERE id = $1
		`

	var (
		id          	   sql.NullString
		state_number       sql.NullString
		model       	   sql.NullString
		status             sql.NullString
		price  			   sql.NullFloat64
		daily_limit        sql.NullFloat64
		over_limit     	   sql.NullFloat64
		investor_percentage  sql.NullFloat64
		investor_id        sql.NullString
		km  			   sql.NullFloat64
		createdAt   	   sql.NullString
		updatedAt   	   sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
				&id,
				&state_number,
				&model,
				&status,
				&price,
				&daily_limit,
				&over_limit,
				&investor_percentage,
				&investor_id,
				&km,
				&createdAt,
				&updatedAt,
		)

		if err != nil {
			return nil, err
		}

	car := &models.Car{
			Id:                  id.String,
			State_number:        state_number.String,
			Model:               model.String,
			Status:              status.String,
			Price:               price.Float64,
			Daily_limit:         daily_limit.Float64,
			Over_limit:          over_limit.Float64,
			Investor_percentage: investor_percentage.Float64,
			Investor_id:         investor_id.String,
			Km:                  km.Float64,
			CreatedAt:           createdAt.String,
			UpdatedAt:           updatedAt.String,
		}
	return car,nil
}
func (r *CarRepo)GetList(ctx context.Context, req *models.GetListCarRequest) (*models.GetListCarResponse, error) {

	var (
		resp   models.GetListCarResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		search = req.Search
	)

	query := `
		SELECT
			COUNT(*) OVER(),
				id,
				state_number,
				model,
				status,
				price,
				daily_limit,
				over_limit,
				investor_percentage,
				investor_id,
				km,
				created_at,
				updated_at 
		FROM Car
	
	`
	if search != "" {
		search = fmt.Sprintf("where model like  '%s%s' ", req.Search,"%")
		query += search
	}
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit
	
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return &models.GetListCarResponse{}, err
	}
	var (
		id          	   sql.NullString
		state_number       sql.NullString
		model       	   sql.NullString
		status             sql.NullString
		price  			   sql.NullFloat64
		daily_limit        sql.NullFloat64
		over_limit     	   sql.NullFloat64
		investor_percentage  sql.NullFloat64
		investor_id        sql.NullString
		km  			   sql.NullFloat64
		createdAt   	   sql.NullString
		updatedAt   	   sql.NullString
	)

	for rows.Next() {


		err = rows.Scan(
				&resp.Count,
				&id,
				&state_number,
				&model,
				&status,
				&price,
				&daily_limit,
				&over_limit,
				&investor_percentage,
				&investor_id,
				&km,
				&createdAt,
				&updatedAt,
		)
		
		car := &models.Car{
			Id:                  id.String,
			State_number:        state_number.String,
			Model:               model.String,
			Status:              status.String,
			Price:               price.Float64,
			Daily_limit:         daily_limit.Float64,
			Over_limit:          over_limit.Float64,
			Investor_percentage: investor_percentage.Float64,
			Investor_id:         investor_id.String,
			Km:                  km.Float64,
			CreatedAt:           createdAt.String,
			UpdatedAt:           updatedAt.String,
		}
		if err != nil {
			return &models.GetListCarResponse{}, err
		}
		
		resp.Cars = append(resp.Cars, car)


	}
	return &resp, nil
}


func (r *CarRepo)Update(ctx context.Context, car *models.UpdateCar) error {
	query := `
		UPDATE 
			Car 
		SET 
			state_number = $2,
			model = $3,
			price = $4,
			daily_limit = $5,
			over_limit = $6,
			investor_id = $7,
			km = $8,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx,query,
		car.Id,
		car.State_number,
		car.Model,
		car.Price,
		car.Daily_limit,
		car.Over_limit,
		car.Investor_id,
		car.Km,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *CarRepo)Delete(ctx context.Context, req *models.CarPrimeryKey) error {

	_, err := r.db.Exec(ctx,"DELETE FROM Car WHERE id  = $1 ", req.Id)
		if err != nil {
			return err
		}

	return nil
}