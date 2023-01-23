package postgres

import (
	"add/models"
	"database/sql"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	
)

type ClientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *ClientRepo {
	return &ClientRepo{
		db: db,
	}
}

func (r *ClientRepo) Insert(ctx context.Context, req *models.CreateClient) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO Client (
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
		req.First_name,
		req.Last_name,
		req.Address,
		req.Phone_number,
	)

	if err != nil {
		return "", err
	}


	return id, nil
}

func (r *ClientRepo) GetByID(ctx context.Context, req *models.ClientPrimeryKey) (*models.Client, error) {

	query := `
		SELECT
			id,
			first_name,
			last_name,
			address,
			phone_number,
			created_at,
			updated_at
		FROM
			Client
		WHERE id = $1
	`

	var (
		id          	 sql.NullString
		first_name       sql.NullString
		last_name        sql.NullString
		address          sql.NullString
		phone_number     sql.NullString
		createdAt   	 sql.NullString
		updatedAt   	 sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&first_name,
			&last_name,
			&address,
			&phone_number,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	client := &models.Client{
			Id:         	id.String,
			First_name: 	first_name.String,
			Last_name:  	last_name.String,
			Address:    	address.String,
			Phone_number: 	phone_number.String,
			CreatedAt:  	createdAt.String,
			UpdatedAt:  	updatedAt.String,
	 }
	return client,nil
}

func (r *ClientRepo)GetList(ctx context.Context, req *models.GetListClientRequest) (*models.GetListClientResponse, error) {

	var (
		resp   models.GetListClientResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		search = req.Search
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
		FROM Client
	
	`
	if search != "" {
		search = fmt.Sprintf("where first_name like  '%s%s' ", req.Search,"%")
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
		return &models.GetListClientResponse{}, err
	}

	var (
		id          	 sql.NullString
		first_name       sql.NullString
		last_name        sql.NullString
		address          sql.NullString
		phone_number     sql.NullString
		createdAt   	 sql.NullString
		updatedAt   	 sql.NullString
	)

	for rows.Next() {


		err = rows.Scan(
			&resp.Count,
			&id,
			&first_name,
			&last_name,
			&address,
			&phone_number,
			&createdAt,
			&updatedAt,
		)

		client := models.CreateClient{
			Id:         id.String,
			First_name: first_name.String,
			Last_name:  last_name.String,
			Address:    address.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		}
		if err != nil {
			return &models.GetListClientResponse{}, err
		}
		
		resp.Client = append(resp.Client, &client)


	}
	return &resp, nil
}


func (r *ClientRepo)Update(ctx context.Context, client *models.UpdateClient) error {

	query := `
		UPDATE 
			Client 
		SET 
			first_name = $2,
			last_name = $3,
			address = $4,
			phone_number = $5,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx,query,
		client.Id,
		client.First_name,
		client.Last_name,
		client.Address,
		client.Phone_number,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ClientRepo)Delete(ctx context.Context, req *models.ClientPrimeryKey) error {

	_, err := r.db.Exec(ctx,"DELETE FROM Client WHERE id  = $1 ", req.Id)
		if err != nil {
			return err
		}

	return nil
}
