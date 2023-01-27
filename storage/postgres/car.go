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
		dailyLimit         sql.NullInt64
		overlimit          sql.NullInt64
		investorPercentage sql.NullFloat64
		branchPercentage   sql.NullFloat64
		investorId         sql.NullString
		branchId           sql.NullString
		km                 sql.NullInt64
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
		DailyLimit:         int(dailyLimit.Int64),
		OverLimit:          int(overlimit.Int64),
		InvestorPercentage: investorPercentage.Float64,
		BranchPercentage:   branchPercentage.Float64,
		InvestorId:         investorId.String,
		BranchId:  			branchId.String,
		Km:                 int(km.Int64),
		CreatedAt:          createdAt.String,
		UpdatedAt:          updatedAt.String,
	}

	return resp, err
}

func (r *CarRepo) GetList(ctx context.Context, req *models.GetListCarRequest) (*models.GetListCarResponse, error) {
	var (
		offset = "OFFSET 0"
		limit  = "LIMIT 10"
		resp   = &models.GetListCarResponse{}
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
	`

	query += offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var (
			id                 sql.NullString
			stateNumber        sql.NullString
			model              sql.NullString
			status             sql.NullString
			price              sql.NullFloat64
			dailyLimit         sql.NullInt64
			overlimit          sql.NullInt64
			investorPercentage sql.NullFloat64
			branchPercentage   sql.NullFloat64
			investorId         sql.NullString
			branchId           sql.NullString
			km                 sql.NullInt64
			createdAt          sql.NullString
			updatedAt          sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
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

		resp.Cars = append(resp.Cars, &models.Car{
			Id:                 id.String,
			StateNumber:        stateNumber.String,
			Model:              model.String,
			Status:             status.String,
			Price:              price.Float64,
			DailyLimit:         int(dailyLimit.Int64),
			OverLimit:          int(overlimit.Int64),
			InvestorPercentage: investorPercentage.Float64,
			BranchPercentage: 	branchPercentage.Float64,
			InvestorId:         investorId.String,
			BranchId: 			branchId.String,	
			Km:                 int(km.Int64),
			CreatedAt:          createdAt.String,
			UpdatedAt:          updatedAt.String,
		})
	}

	return resp, err
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
