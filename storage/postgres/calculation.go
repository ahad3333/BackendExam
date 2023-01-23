package postgres

import (
	"add/models"
	"database/sql"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	
)

type CalcRepo struct {
	db *pgxpool.Pool
}

func NewCalculationRepo(db *pgxpool.Pool) *CalcRepo {
	return &CalcRepo{
		db: db,
	}
}

func (r *CalcRepo)GetListInves(ctx context.Context, req *models.GetListInvesRequest) (*models.GetListInves, error) {

	var (
		resp   models.GetListInves
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		search = req.Search
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			name,
			price,
			updated_at
		FROM InvestorBenefit
	`
	if search != "" {
		search = fmt.Sprintf(" where name like  '%s%s' ", req.Search,"%")
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
		return &models.GetListInves{}, err
	}
	var (
		name       		 sql.NullString
		price        	 sql.NullFloat64
		updatedAt   	 sql.NullString
	)

	for rows.Next() {


		err = rows.Scan(
			&resp.Count,
			&name,
			&price,
			&updatedAt,
		)
		inves := models.InvestorBenefit{
			Name:      name.String,
			Price:     price.Float64,
			UpdatedAt: updatedAt.String,
		}
		if err != nil {
			return &models.GetListInves{}, err
		}
		
		resp.Benefits = append(resp.Benefits, &inves)


	}
	return &resp, nil
}

func (r *CalcRepo)GetListBebt(ctx context.Context, req *models.GetListBebtRequest) (*models.GetListBebt, error) {

	var (
		resp   models.GetListBebt
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		search = req.Search
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			name,
			bebt,
			updated_at 
		FROM Bebtors
	
	`
	if search != "" {
		search = fmt.Sprintf("where name like  '%s%s' ", req.Search,"%")
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
		return &models.GetListBebt{}, err
	}
	var (
		name       		 sql.NullString
		bebt        	 sql.NullFloat64
		updatedAt   	 sql.NullString
	)

	for rows.Next() {


		err = rows.Scan(
			&resp.Count,
			&name,
			&bebt,
			&updatedAt,
		)
		bebt := models.Bebtors{
			Name:      name.String,
			Bebt:      bebt.Float64,
			UpdatedAt: updatedAt.String,
		}
		if err != nil {
			return &models.GetListBebt{}, err
		}
		
		resp.Bebtors = append(resp.Bebtors, &bebt)


	}
	return &resp, nil
}