// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type ApiKey struct {
	ID      int32         `json:"id"`
	Key     string        `json:"key"`
	Owner   int32         `json:"owner"`
	Enabled sql.NullBool  `json:"enabled"`
	Hits    sql.NullInt32 `json:"hits"`
	Errors  sql.NullInt32 `json:"errors"`
}

type Storage struct {
	ID        int32         `json:"id"`
	Key       string        `json:"key"`
	Value     string        `json:"value"`
	Available sql.NullBool  `json:"available"`
	CreatedBy int32         `json:"created_by"`
	Downloads sql.NullInt32 `json:"downloads"`
	Errors    sql.NullInt32 `json:"errors"`
}

type User struct {
	ID int32 `json:"id"`
	// username
	Username string `json:"username"`
	// Name of the User
	Name string `json:"name"`
}