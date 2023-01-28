package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"app/models"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
)

type CarRepo struct {
	db *pgxpool.Pool
}

func NewCarRepo(db *pgxpool.Pool) *CarRepo {
	return &CarRepo{
		db: db,
	}
}

func (r *CarRepo) Insert(ctx context.Context, car *models.CreateCar) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO car (
			id,
			state_number,
			model,
			price,
			daily_limit,
			over_limit,
			investor_percentage,
			branch_percentage,
			investor_id,
			branch_id,
			km,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		car.StateNumber,
		car.Model,
		car.Price,
		car.DailyLimit,
		car.OverLimit,
		car.InvestorPercentage,
		car.BranchPercentage,
		car.InvestorId,
		car.BranchId,
		car.Km,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *CarRepo) GetByID(ctx context.Context, req *models.CarPrimeryKey) (*models.Car, error) {

	var (
		id                 sql.NullString
		stateNumber        sql.NullString
		model              sql.NullString
		status             sql.NullString
		price              sql.NullFloat64
		dailyLimit         sql.NullFloat64
		overlimit          sql.NullFloat64
		investorPercentage sql.NullFloat64
		branchPercentage   sql.NullFloat64
		investorId         sql.NullString
		branchId           sql.NullString
		km                 sql.NullFloat64
		createdAt          sql.NullString
		updatedAt          sql.NullString
	)

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
			branch_percentage,
			investor_id,
			branch_id,
			km,
			created_at,
			updated_at
		FROM car
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&stateNumber,
		&model,
		&status,
		&price,
		&dailyLimit,
		&overlimit,
		&investorPercentage,
		&branchPercentage,
		&investorId,
		&branchId,
		&km,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	resp := &models.Car{
		Id:                 id.String,
		StateNumber:        stateNumber.String,
		Model:              model.String,
		Status:             status.String,
		Price:              price.Float64,
		DailyLimit:         dailyLimit.Float64,
		OverLimit:          overlimit.Float64,
		InvestorPercentage: investorPercentage.Float64,
		BranchPercentage:   branchPercentage.Float64,
		InvestorId:         investorId.String,
		BranchId:  			branchId.String,
		Km:                 km.Float64,
		CreatedAt:          createdAt.String,
		UpdatedAt:          updatedAt.String,
	}

	return resp, err
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
			StateNumber:        state_number.String,
			Model:               model.String,
			Status:              status.String,
			Price:               price.Float64,
			DailyLimit:         daily_limit.Float64,
			OverLimit:          over_limit.Float64,
			InvestorPercentage: investor_percentage.Float64,
			InvestorId:         investor_id.String,
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

func (r *CarRepo) Update(ctx context.Context, car *models.UpdateCar) (int64, error) {
	query := `
		UPDATE
			car
		SET
			state_number = $2,
			model = $3,
			price = $4,
			daily_limit = $5,
			over_limit = $6,
			investor_percentage = $7,
			branch_percentage = $8,
			investor_id = $9,
			branch_id = $10,
			km = $11,
			updated_at = now()
		WHERE id = $1
	`

	rowsAffacted, err := r.db.Exec(ctx, query,
		car.Id,
		car.StateNumber,
		car.Model,
		car.Price,
		car.DailyLimit,
		car.OverLimit,
		car.InvestorPercentage,
		car.BranchPercentage,
		car.InvestorId,
		car.BranchId,
		car.Km,
	)

	if err != nil {
		return 0, err
	}

	return rowsAffacted.RowsAffected(), nil
}

func (r *CarRepo) Delete(ctx context.Context, req *models.CarPrimeryKey) error {

	_, err := r.db.Exec(ctx, "delete from car where id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
