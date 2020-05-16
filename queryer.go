// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package sqlex

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var (
	// ensure sqlx.DB and sqlx.Tx implement Queryer interface.
	_ Queryer = &sqlx.DB{}
	_ Queryer = &sqlx.Tx{}
)

// Queryer is an interface that wraps sqlx.Queryer and sqlx.QueryerContext.
type Queryer interface {
	sqlx.Queryer
	sqlx.QueryerContext
}

// Query equals to sqlx.Queryer.Query and sqlx.QueryerContext.QueryContext.
func Query(ctx context.Context, q Queryer, query string, args ...interface{}) (*sql.Rows, error) {
	if ctx == nil {
		return q.Query(query, args...)
	}
	return q.QueryContext(ctx, query, args...)
}

// Queryx equals to sqlx.Queryer.Queryx and sqlx.QueryerContext.QueryxContext.
func Queryx(ctx context.Context, q Queryer, query string, args ...interface{}) (*sqlx.Rows, error) {
	if ctx == nil {
		return q.Queryx(query, args...)
	}
	return q.QueryxContext(ctx, query, args...)
}

// QueryRowx equals to sqlx.Queryer.QueryRowx and sqlx.QueryerContext.QueryRowxContext.
func QueryRowx(ctx context.Context, q Queryer, query string, args ...interface{}) *sqlx.Row {
	if ctx == nil {
		return q.QueryRowx(query, args...)
	}
	return q.QueryRowxContext(ctx, query, args...)
}

// Get equals to sqlx.Get and sqlx.GetContext.
func Get(ctx context.Context, q Queryer, dest interface{}, query string, args ...interface{}) error {
	if ctx == nil {
		return sqlx.Get(q, dest, query, args...)
	}

	return sqlx.GetContext(ctx, q, dest, query, args...)
}

// Select equals to sqlx.Select and sqlx.SelectContext.
func Select(ctx context.Context, q Queryer, dest interface{}, query string, args ...interface{}) error {
	if ctx == nil {
		return sqlx.Select(q, dest, query, args...)
	}

	return sqlx.SelectContext(ctx, q, dest, query, args...)
}
