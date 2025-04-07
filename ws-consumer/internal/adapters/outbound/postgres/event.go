package postgres

import (
	"context"
	"encoding/json"
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
	dataJSON, err := json.Marshal(event.Data)
	if err != nil {
		return err
	}

	query := NewQueryBuilder().
		Query(BASE_EVENT_QUERY).
		AddArgs(event.ID, userID, dataJSON)

	err = db.db.GetContext(ctx, event, query.Build(), query.GetArgs()...)
	if err != nil {
		return err
	}

	return nil
}
