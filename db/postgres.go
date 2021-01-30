// Package db impls db functionality
package db

import (
	"context"

	"github.com/go-pg/pg/v10"
)

// Logger godoc
type Logger struct{}

// BeforeQuery godoc
func (d Logger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

// AfterQuery godoc
func (d Logger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	return nil
}

// NewDBConnection connects to a db
func NewDBConnection(options *pg.Options) *pg.DB {
	return pg.Connect(options)
}
