// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: movie.sql

package datastore

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

const createMovie = `-- name: CreateMovie :execresult
INSERT INTO movie (movie_id, extl_id, title, rated, released, run_time, director, writer,
                   create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
`

type CreateMovieParams struct {
	MovieID         uuid.UUID
	ExtlID          string
	Title           string
	Rated           sql.NullString
	Released        sql.NullTime
	RunTime         sql.NullInt32
	Director        sql.NullString
	Writer          sql.NullString
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, createMovie,
		arg.MovieID,
		arg.ExtlID,
		arg.Title,
		arg.Rated,
		arg.Released,
		arg.RunTime,
		arg.Director,
		arg.Writer,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
}

const deleteMovie = `-- name: DeleteMovie :execrows
DELETE FROM movie
WHERE movie_id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, movieID uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deleteMovie, movieID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findMovieByExternalID = `-- name: FindMovieByExternalID :one
SELECT m.movie_id, m.extl_id, m.title, m.rated, m.released, m.run_time, m.director, m.writer, m.create_app_id, m.create_user_id, m.create_timestamp, m.update_app_id, m.update_user_id, m.update_timestamp
FROM movie m
WHERE m.extl_id = $1
`

func (q *Queries) FindMovieByExternalID(ctx context.Context, extlID string) (Movie, error) {
	row := q.db.QueryRow(ctx, findMovieByExternalID, extlID)
	var i Movie
	err := row.Scan(
		&i.MovieID,
		&i.ExtlID,
		&i.Title,
		&i.Rated,
		&i.Released,
		&i.RunTime,
		&i.Director,
		&i.Writer,
		&i.CreateAppID,
		&i.CreateUserID,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateUserID,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findMovieByExternalIDWithAudit = `-- name: FindMovieByExternalIDWithAudit :one
SELECT m.movie_id,
       m.extl_id,
       m.title,
       m.rated,
       m.released,
       m.run_time,
       m.director,
       m.writer,
       m.create_app_id,
       ca.org_id          create_app_org_id,
       ca.app_extl_id     create_app_extl_id,
       ca.app_name        create_app_name,
       ca.app_description create_app_description,
       m.create_user_id,
       cu.first_name     create_user_first_name,
       cu.last_name      create_user_last_name,
       m.create_timestamp,
       m.update_app_id,
       ua.org_id          update_app_org_id,
       ua.app_extl_id     update_app_extl_id,
       ua.app_name        update_app_name,
       ua.app_description update_app_description,
       m.update_user_id,
       uu.first_name     update_user_first_name,
       uu.last_name      update_user_last_name,
       m.update_timestamp
FROM movie m
         INNER JOIN app ca on ca.app_id = m.create_app_id
         INNER JOIN app ua on ua.app_id = m.update_app_id
         LEFT JOIN users cu on cu.user_id = m.create_user_id
         LEFT JOIN users uu on uu.user_id = m.update_user_id
WHERE m.extl_id = $1
`

type FindMovieByExternalIDWithAuditRow struct {
	MovieID              uuid.UUID
	ExtlID               string
	Title                string
	Rated                sql.NullString
	Released             sql.NullTime
	RunTime              sql.NullInt32
	Director             sql.NullString
	Writer               sql.NullString
	CreateAppID          uuid.UUID
	CreateAppOrgID       uuid.UUID
	CreateAppExtlID      string
	CreateAppName        string
	CreateAppDescription string
	CreateUserID         uuid.NullUUID
	CreateUserFirstName  sql.NullString
	CreateUserLastName   sql.NullString
	CreateTimestamp      time.Time
	UpdateAppID          uuid.UUID
	UpdateAppOrgID       uuid.UUID
	UpdateAppExtlID      string
	UpdateAppName        string
	UpdateAppDescription string
	UpdateUserID         uuid.NullUUID
	UpdateUserFirstName  sql.NullString
	UpdateUserLastName   sql.NullString
	UpdateTimestamp      time.Time
}

func (q *Queries) FindMovieByExternalIDWithAudit(ctx context.Context, extlID string) (FindMovieByExternalIDWithAuditRow, error) {
	row := q.db.QueryRow(ctx, findMovieByExternalIDWithAudit, extlID)
	var i FindMovieByExternalIDWithAuditRow
	err := row.Scan(
		&i.MovieID,
		&i.ExtlID,
		&i.Title,
		&i.Rated,
		&i.Released,
		&i.RunTime,
		&i.Director,
		&i.Writer,
		&i.CreateAppID,
		&i.CreateAppOrgID,
		&i.CreateAppExtlID,
		&i.CreateAppName,
		&i.CreateAppDescription,
		&i.CreateUserID,
		&i.CreateUserFirstName,
		&i.CreateUserLastName,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateAppOrgID,
		&i.UpdateAppExtlID,
		&i.UpdateAppName,
		&i.UpdateAppDescription,
		&i.UpdateUserID,
		&i.UpdateUserFirstName,
		&i.UpdateUserLastName,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findMovies = `-- name: FindMovies :many
SELECT m.movie_id,
       m.extl_id,
       m.title,
       m.rated,
       m.released,
       m.run_time,
       m.director,
       m.writer,
       m.create_app_id,
       ca.org_id          create_app_org_id,
       ca.app_extl_id     create_app_extl_id,
       ca.app_name        create_app_name,
       ca.app_description create_app_description,
       m.create_user_id,
       cu.first_name     create_user_first_name,
       cu.last_name      create_user_last_name,
       m.create_timestamp,
       m.update_app_id,
       ua.org_id          update_app_org_id,
       ua.app_extl_id     update_app_extl_id,
       ua.app_name        update_app_name,
       ua.app_description update_app_description,
       m.update_user_id,
       uu.first_name     update_user_first_name,
       uu.last_name      update_user_last_name,
       m.update_timestamp
FROM movie m
         INNER JOIN app ca on ca.app_id = m.create_app_id
         INNER JOIN app ua on ua.app_id = m.update_app_id
         LEFT JOIN users cu on cu.user_id = m.create_user_id
         LEFT JOIN users uu on uu.user_id = m.update_user_id
`

type FindMoviesRow struct {
	MovieID              uuid.UUID
	ExtlID               string
	Title                string
	Rated                sql.NullString
	Released             sql.NullTime
	RunTime              sql.NullInt32
	Director             sql.NullString
	Writer               sql.NullString
	CreateAppID          uuid.UUID
	CreateAppOrgID       uuid.UUID
	CreateAppExtlID      string
	CreateAppName        string
	CreateAppDescription string
	CreateUserID         uuid.NullUUID
	CreateUserFirstName  sql.NullString
	CreateUserLastName   sql.NullString
	CreateTimestamp      time.Time
	UpdateAppID          uuid.UUID
	UpdateAppOrgID       uuid.UUID
	UpdateAppExtlID      string
	UpdateAppName        string
	UpdateAppDescription string
	UpdateUserID         uuid.NullUUID
	UpdateUserFirstName  sql.NullString
	UpdateUserLastName   sql.NullString
	UpdateTimestamp      time.Time
}

func (q *Queries) FindMovies(ctx context.Context) ([]FindMoviesRow, error) {
	rows, err := q.db.Query(ctx, findMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindMoviesRow
	for rows.Next() {
		var i FindMoviesRow
		if err := rows.Scan(
			&i.MovieID,
			&i.ExtlID,
			&i.Title,
			&i.Rated,
			&i.Released,
			&i.RunTime,
			&i.Director,
			&i.Writer,
			&i.CreateAppID,
			&i.CreateAppOrgID,
			&i.CreateAppExtlID,
			&i.CreateAppName,
			&i.CreateAppDescription,
			&i.CreateUserID,
			&i.CreateUserFirstName,
			&i.CreateUserLastName,
			&i.CreateTimestamp,
			&i.UpdateAppID,
			&i.UpdateAppOrgID,
			&i.UpdateAppExtlID,
			&i.UpdateAppName,
			&i.UpdateAppDescription,
			&i.UpdateUserID,
			&i.UpdateUserFirstName,
			&i.UpdateUserLastName,
			&i.UpdateTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findMoviesByTitle = `-- name: FindMoviesByTitle :many
SELECT m.movie_id, m.extl_id, m.title, m.rated, m.released, m.run_time, m.director, m.writer, m.create_app_id, m.create_user_id, m.create_timestamp, m.update_app_id, m.update_user_id, m.update_timestamp
FROM movie m
WHERE m.title = $1
`

func (q *Queries) FindMoviesByTitle(ctx context.Context, title string) ([]Movie, error) {
	rows, err := q.db.Query(ctx, findMoviesByTitle, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.MovieID,
			&i.ExtlID,
			&i.Title,
			&i.Rated,
			&i.Released,
			&i.RunTime,
			&i.Director,
			&i.Writer,
			&i.CreateAppID,
			&i.CreateUserID,
			&i.CreateTimestamp,
			&i.UpdateAppID,
			&i.UpdateUserID,
			&i.UpdateTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovie = `-- name: UpdateMovie :exec
UPDATE movie
SET title            = $1,
    rated            = $2,
    released         = $3,
    run_time         = $4,
    director         = $5,
    writer           = $6,
    update_app_id    = $7,
    update_user_id   = $8,
    update_timestamp = $9
WHERE movie_id = $10
`

type UpdateMovieParams struct {
	Title           string
	Rated           sql.NullString
	Released        sql.NullTime
	RunTime         sql.NullInt32
	Director        sql.NullString
	Writer          sql.NullString
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
	MovieID         uuid.UUID
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) error {
	_, err := q.db.Exec(ctx, updateMovie,
		arg.Title,
		arg.Rated,
		arg.Released,
		arg.RunTime,
		arg.Director,
		arg.Writer,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
		arg.MovieID,
	)
	return err
}
