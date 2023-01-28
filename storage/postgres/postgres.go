package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/config"
	"app/storage"
)

type Store struct {
	db       *pgxpool.Pool
	investor *InvestorRepo
	car      *CarRepo
	client   *ClientRepo
	order    *OrderRepo
	report   *ReportRepo
	branch   *BranchRepo
	user     *UserRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))

	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConn

	pool, err := pgxpool.ConnectConfig(ctx, config)

	if err != nil {
		return nil, err
	}

	return &Store{
		db:       pool,
		investor: NewInvestorRepo(pool),
		car:      NewCarRepo(pool),
		client:   NewClientRepo(pool),
		order:    NewOrderRepo(pool),
		report:   NewReportRepo(pool),
		branch:   NewBranchRepo(pool),
		user:  NewUserRepo(pool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Investor() storage.InvestorRepoI {

	if s.investor == nil {
		s.investor = NewInvestorRepo(s.db)
	}

	return s.investor
}

func (s *Store) Car() storage.CarRepoI {

	if s.car == nil {
		s.car = NewCarRepo(s.db)
	}

	return s.car
}

func (s *Store) Client() storage.ClientRepoI {

	if s.client == nil {
		s.client = NewClientRepo(s.db)
	}

	return s.client
}

func (s *Store) Order() storage.OrderRepoI {

	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}

	return s.order
}

func (s *Store) Report() storage.ReportRepoI {

	if s.report == nil {
		s.report = NewReportRepo(s.db)
	}

	return s.report
}

func (s *Store) Branch() storage.BranchRepoI {

	if s.branch == nil {
		s.branch =NewBranchRepo(s.db)
	}

	return s.branch
}

func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user =NewUserRepo(s.db)
	}

	return s.user
}