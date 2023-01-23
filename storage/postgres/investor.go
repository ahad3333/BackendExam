package postgres

import (
	"add/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)
type InvestorRepo struct {
	db *pgxpool.Pool
}

func NewInvestorRepo(db *pgxpool.Pool) *InvestorRepo {
	return &InvestorRepo{
		db: db,
	}

}

func (r *InvestorRepo) Insert(ctx context.Context, req *models.CreateInvestor) (string, error) {
	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO investor (
			id,
			name,
			updated_at
		) VALUES ($1, $2, now())
		`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *InvestorRepo) GetByID(ctx context.Context, req *models.InvestorPrimeryKey) (*models.Investor, error) {
	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM
			Investor
		WHERE id = $1
	`
	var (
		id          sql.NullString
		name        sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&name,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

	investor := &models.Investor{
		Id:          id.String,
		Name:        name.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}


  return investor, nil
}


func (r *InvestorRepo)GetList(ctx context.Context, req *models.GetListInvestorRequest) (*models.GetListInvestorResponse, error) {

	var (
		resp   models.GetListInvestorResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		search = req.Search
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			created_at,
			updated_at
		FROM Investor
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

	rows, err := r.db.Query(ctx,query)

	if err != nil {
		return &models.GetListInvestorResponse{}, err
	}

	var (
		id          sql.NullString
		name        sql.NullString
		createdAt 	sql.NullString
		updatedAt 	sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(
			&resp.Count,
			&id,
			&name,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return &models.GetListInvestorResponse{}, err
		}

		investor := models.CreateInvestor{
			Id:          id.String,
			Name:        name.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		}

		resp.Investors = append(resp.Investors, &investor)
	}


	return &resp, nil
}

func (r *InvestorRepo)Update(ctx context.Context, investor *models.UpdateInvestor) error {

	query := `
		UPDATE 
		Investor
		SET 
			name = $2
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx,query,
		investor.Id,
		investor.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *InvestorRepo)Delete(ctx context.Context, req *models.InvestorPrimeryKey) error {
	
	_, err := r.db.Exec(ctx,"DELETE FROM Investor WHERE id = $1 ", req.Id)

	if err != nil {
		return err
	}

	return nil
}
