package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"add/config"
	"add/storage"
)

type Store struct {
	db       *pgxpool.Pool
	car      *CarRepo
	order    *OrderRepo
	client   *ClientRepo
	investor *InvestorRepo
	calc 	 *CalcRepo
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
		client:   NewClientRepo(pool),
		car:      NewCarRepo(pool),
		order:    NewOrderRepo(pool),
		calc: NewCalculationRepo(pool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Client() storage.ClientRepoI {

	if s.client == nil {
		s.client = NewClientRepo(s.db)
	}

	return s.client
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

func (s *Store) Order() storage.OrderRepoI {

	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}

	return s.order
}

func (s *Store) Calc() storage.CalcRepoI {

	if s.calc == nil {
		s.calc = NewCalculationRepo(s.db)
	}

	return s.calc
}
