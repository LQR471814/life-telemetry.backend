// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package eczemadb

import (
	"context"
)

const createEvent = `-- name: CreateEvent :exec
insert into Event (time, duration, state, variant, location, action) values (?, ?, ?, ?, ?, ?)
`

type CreateEventParams struct {
	Time     int64
	Duration int64
	State    float64
	Variant  int64
	Location int64
	Action   int64
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) error {
	_, err := q.db.ExecContext(ctx, createEvent,
		arg.Time,
		arg.Duration,
		arg.State,
		arg.Variant,
		arg.Location,
		arg.Action,
	)
	return err
}
