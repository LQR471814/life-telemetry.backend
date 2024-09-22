package eczema

import (
	"context"
	"database/sql"
	eczemav1 "life-telemetry/proto/services/eczema/v1"
	eczemadb "life-telemetry/services/eczema/db"
	"log/slog"

	"connectrpc.com/connect"

	_ "github.com/mattn/go-sqlite3"
)

var log = slog.With(slog.String("service", "eczema.v1"))

type Service struct {
	db  *sql.DB
	qry *eczemadb.Queries
}

func NewService(database *sql.DB) Service {
	return Service{
		db:  database,
		qry: eczemadb.New(database),
	}
}

func (s Service) Push(ctx context.Context, req *connect.Request[eczemav1.PushRequest]) (*connect.Response[eczemav1.PushResponse], error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	txqry := s.qry.WithTx(tx)

	for _, e := range req.Msg.GetEvents() {
		err := txqry.CreateEvent(ctx, eczemadb.CreateEventParams{
			Time:     int64(e.GetTime()),
			Duration: int64(e.GetDuration()),

			State:    float64(e.GetState()),
			Variant:  int64(e.GetVariant()),
			Location: int64(e.GetLocation()),
			Action:   int64(e.GetAction()),
		})
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &connect.Response[eczemav1.PushResponse]{Msg: &eczemav1.PushResponse{}}, nil
}
