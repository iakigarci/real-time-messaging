package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/iakigarci/go-ddd-microservice-template/config"
	"github.com/iakigarci/go-ddd-microservice-template/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const (
	_defaultConnAttempts = 5
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	DB     *sqlx.DB
	logger *logger.Logger
}

func NewClient(cfg *config.Config, logger *logger.Logger) (*Postgres, error) {
	postgres, err := initPg(cfg, logger)
	if err != nil {
		logger.Error("failed to initialize postgres client", zap.Error(err))
		return nil, err
	}

	return postgres, nil
}

func initPg(cfg *config.Config, logger *logger.Logger) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  cfg.Postgres.PoolMax,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
		logger:       logger,
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	var err error
	for pg.connAttempts > 0 {
		pg.DB, err = sqlx.Open("postgres", connectionString)
		if err == nil {
			pg.DB.SetMaxOpenConns(pg.maxPoolSize)
			pg.DB.SetMaxIdleConns(pg.maxPoolSize)

			if err = pg.DB.Ping(); err == nil {
				break
			}
		}

		pg.logger.InfoAttrs("Postgres trying to connect, attempts left", map[string]string{
			"attempts": strconv.Itoa(pg.connAttempts),
		})
		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	if err != nil {
		pg.logger.ErrorAttrs("Postgres connection failed", err, map[string]string{
			"host":    cfg.Postgres.Host,
			"port":    strconv.Itoa(cfg.Postgres.Port),
			"user":    cfg.Postgres.User,
			"dbname":  cfg.Postgres.DBName,
			"sslmode": cfg.Postgres.SSLMode,
		})
		return nil, err
	}

	pg.logger.Info("Postgres connected successfully")

	return pg, nil
}

func (p *Postgres) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}

func (p *Postgres) Ping(ctx context.Context) error {
	return p.DB.PingContext(ctx)
}

func (p *Postgres) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return p.DB.BeginTx(ctx, &sql.TxOptions{})
}
