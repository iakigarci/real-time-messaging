package postgres

import (
	"context"
	"real-time-messaging/consumer/internal/domain/events"
	port "real-time-messaging/consumer/internal/domain/ports"

	"github.com/jmoiron/sqlx"
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) port.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (db *EventRepository) CreateEvent(ctx context.Context, event *events.BaseEvent, userID string) error {
	query := NewQueryBuilder().
		Query(BASE_EVENT_QUERY).
		AddArgs(event.ID, userID, event.Data)

	err := db.db.GetContext(ctx, event, query.Build(), query.GetArgs()...)
	if err != nil {
		return err
	}

	return nil
}
