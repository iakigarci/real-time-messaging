package postgres

import (
	"context"
	"database/sql"
	"errors"
	"real-time-messaging/consumer/internal/domain/entities"
	port "real-time-messaging/consumer/internal/domain/ports"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) port.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (db *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := &entities.User{}

	query := NewQueryBuilder().
		Query(BASE_USER_QUERY).
		Where("email = $1").
		AddArgs(email)

	err := db.db.GetContext(ctx, user, query.Build(), query.GetArgs()...)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
