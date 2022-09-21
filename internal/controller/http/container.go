package http

import "github.com/jmoiron/sqlx"

type container struct {
}

func NewContainer(db *sqlx.DB) *container {
	return &container{}
}
