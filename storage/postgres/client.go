package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/models"
	"app/pkg/helper"
)

type ClientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *ClientRepo {
	return &ClientRepo{
		db: db,
	}
}

func (r *ClientRepo) Insert(ctx context.Context, client *models.CreateClient) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO client (
			id,
			first_name,
			last_name,
			address,
			phone_number,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		client.FirstName,
		client.LastName,
		client.Address,
		client.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *ClientRepo) GetByID(ctx context.Context, req *models.ClientPrimeryKey) (*models.Client, error) {

	var (
		id          sql.NullString
		firstName   sql.NullString
		lastName    sql.NullString
		address     sql.NullString
		phoneNumber sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query := `
		SELECT
			id,
			first_name,
			last_name,
			address,
			phone_number,
			created_at,
			updated_at
		FROM client
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&firstName,
		&lastName,
		&address,
		&phoneNumber,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	resp := &models.Client{
		Id:          id.String,
		FirstName:   firstName.String,
		LastName:    lastName.String,
		Address:     address.String,
		PhoneNumber: phoneNumber.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}

	return resp, err
}

func (r *ClientRepo) GetList(ctx context.Context, req *models.GetListClientRequest) (*models.GetListClientResponse, error) {
	var (
		offset = "OFFSET 0"
		limit  = "LIMIT 10"
		resp   = &models.GetListClientResponse{}
	)

	

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name,
			last_name,
			address,
			phone_number,
			created_at,
			updated_at
		FROM client
	`
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	query += offset +" "+ limit

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	var (
		id          sql.NullString
		firstName   sql.NullString
		lastName    sql.NullString
		address     sql.NullString
		phoneNumber sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&address,
			&phoneNumber,
			&createdAt,
			&updatedAt,
		)
		fmt.Println(resp)


		resp.Clients = append(resp.Clients, &models.Client{
			Id:          id.String,
			FirstName:   firstName.String,
			LastName:    lastName.String,
			Address:     address.String,
			PhoneNumber: phoneNumber.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, err
}

func (r *ClientRepo) Update(ctx context.Context, client *models.UpdateClient) (int64, error) {

	var (
		params map[string]interface{}
		query  string
	)

	query = `
		UPDATE
			client
		SET
			first_name = :first_name,
			last_name = :last_name,
			address = :address,
			phone_number = :phone_number,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"first_name":   client.FirstName,
		"last_name":    client.LastName,
		"address":      client.Address,
		"phone_number": client.PhoneNumber,
		"id": client.Id,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (r *ClientRepo) Delete(ctx context.Context, req *models.ClientPrimeryKey) error {

	_, err := r.db.Exec(ctx, "delete from client where id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
