-- name: CreateEvent :exec
insert into Event (time, duration, state, variant, location, action) values (?, ?, ?, ?, ?, ?);
