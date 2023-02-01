package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"app/models"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (f *UserRepo) Create(ctx context.Context, user *models.CreateUser) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO users (
			id,
			first_name,
			last_name,
			login,
			typeu,
			password,
			phone_number,
			updated_at
		) VALUES ( $1, $2, $3, $4, $5, $6,$7, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		user.FirstName,
		user.LastName,
		user.Login,
		user.TypeU,
		user.Password,
		user.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *UserRepo) GetByPKey(ctx context.Context, pkey *models.UserPrimarKey) (*models.User, error) {

	var (
		id          sql.NullString
		firstName   sql.NullString
		lastName    sql.NullString
		login       sql.NullString
		typeu      sql.NullString
		password    sql.NullString
		phoneNumber sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
		typeuser     sql.NullString
	)

	if len(pkey.Login) > 0 {

		err := f.db.QueryRow(ctx, "SELECT id, typeu FROM users WHERE login = $1 and password = $2", pkey.Login,pkey.Password).
			Scan(&pkey.Id,&typeuser)
			fmt.Println(typeuser)
			ab:=&models.UserPrimarKey{
				TypeU: typeuser.String,
			}
			pkey.TypeU=ab.TypeU
		if err != nil {
			return nil, err
		}
	}

	query := `
		SELECT
			id,
			first_name,
			last_name,
			login,
			typeu,
			password,
			phone_number,
			created_at,
			updated_at
		FROM
			users
		WHERE id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&firstName,
			&lastName,
			&login,
			&typeu,
			&password,
			&phoneNumber,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:          id.String,
		FirstName:   firstName.String,
		LastName:    lastName.String,
		Login:       login.String,
		TypeU: 		 typeu.String,
		Password:    password.String,
		PhoneNumber: phoneNumber.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (f *UserRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {

	var (
		resp   = models.GetListUserResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 5"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name,
			last_name,
			login,
			password,
			phone_number,
			created_at,
			updated_at
		FROM
			users
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id          sql.NullString
			firstName   sql.NullString
			lastName    sql.NullString
			login       sql.NullString
			password    sql.NullString
			phoneNumber sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&login,
			&password,
			&phoneNumber,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:          id.String,
			FirstName:   firstName.String,
			LastName:    lastName.String,
			Login:       login.String,
			Password:    password.String,
			PhoneNumber: phoneNumber.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})

	}

	return &resp, err
}

