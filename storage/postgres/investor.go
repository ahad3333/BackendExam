package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"app/models"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
)

type InvestorRepo struct {
	db *pgxpool.Pool
}

func NewInvestorRepo(db *pgxpool.Pool) *InvestorRepo {
	return &InvestorRepo{
		db: db,
	}
}

func (r *InvestorRepo) Insert(ctx context.Context, Investor *models.CreateInvestor) (string, error) {

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
		Investor.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *InvestorRepo) GetByID(ctx context.Context, req *models.InvestorPrimeryKey) (*models.Investor, error) {

	var (
		id        sql.NullString
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM investor
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Investor{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, err
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

		investor := models.Investor{
			Id:          id.String,
			Name:        name.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		}

		resp.Investors = append(resp.Investors, &investor)
	}


	return &resp, nil
}

func (r *InvestorRepo) Update(ctx context.Context, Investor *models.UpdateInvestor) error {
	query := `
		UPDATE
			investor
		SET
			name = $2,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		Investor.Id,
		Investor.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *InvestorRepo) Delete(ctx context.Context, req *models.InvestorPrimeryKey) error {

	_, err := r.db.Exec(ctx, "delete from investor where id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
