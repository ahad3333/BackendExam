package postgres

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/models"
)

type ReportRepo struct {
	db *pgxpool.Pool
}

func NewReportRepo(db *pgxpool.Pool) *ReportRepo {
	return &ReportRepo{
		db: db,
	}
}

func (r *ReportRepo) GetListDebtors(ctx context.Context) (*models.GetListDebtorResponse, error) {

	resp := &models.GetListDebtorResponse{}

	query := `
		SELECT
			COUNT(*) OVER(),
			SUM(d.price),
			cl.first_name || ' ' || cl.last_name AS full_name
		FROM debtors AS d
		JOIN client AS cl ON cl.id = d.client_id
		GROUP BY full_name
	`

	rows, err := r.db.Query(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			price    sql.NullFloat64
			fullName sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&price,
			&fullName,
		)

		if err != nil {
			return nil, err
		}

		resp.Debtors = append(resp.Debtors, &models.Debtor{
			Price:    price.Float64,
			FullName: fullName.String,
		})
	}

	return resp, nil
}

func (r *ReportRepo) GetListInvestorShare(ctx context.Context) (*models.GetListInvestorShareResponse, error) {

	resp := &models.GetListInvestorShareResponse{}

	query := `
		SELECT
			COUNT(*) OVER(),
			i.id,
			i.name,
			SUM(o.paid_price / 100 * c.investor_percentage)
		FROM "order" AS o
		JOIN car AS c ON c.id = o.car_id
		JOIN investor AS i ON i.id = c.investor_id
		GROUP BY i.id, i.name
	`

	rows, err := r.db.Query(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id       sql.NullString
			price    sql.NullFloat64
			fullName sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&id,
			&fullName,
			&price,
		)

		if err != nil {
			return nil, err
		}

		resp.Investors = append(resp.Investors, &models.InvestorShare{
			Id:       id.String,
			Price:    price.Float64,
			FullName: fullName.String,
		})
	}

	return resp, nil
}


func (r *ReportRepo) GetListBranchShare(ctx context.Context) (*models.GetListBranchShareResponse, error) {

	resp := &models.GetListBranchShareResponse{}

	query := `
		SELECT
		COUNT(*) OVER(),
			b.id,
			b.name,
			SUM(o.paid_price-(o.paid_price / 100 * c.investor_percentage))
		FROM "order" AS o
		JOIN car AS c ON c.id = o.car_id
		JOIN branch AS b ON b.id = c.branch_id
		GROUP BY b.id, b.name
	`

	rows, err := r.db.Query(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id       sql.NullString
			price    sql.NullFloat64
			name	 sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&id,
			&name,
			&price,
		)

		if err != nil {
			return nil, err
		}

		resp.Branchs = append(resp.Branchs, &models.BranchShare{
			Id:       id.String,
			Price:    price.Float64,
			Name:	  name.String,
		})
	}

	return resp, nil
}
