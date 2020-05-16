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
	// ensure sqlx.DB and sqlx.Tx implement Execer interface.
	_ Execer = &sqlx.DB{}
	_ Execer = &sqlx.Tx{}
)

// Execer is an interface that wraps sqlx.Execer and sqlx.ExecerContext.
type Execer interface {
	sqlx.Execer
	sqlx.ExecerContext
}

// Exec equals to sqlx.Execer.Exec and sqlx.ExecerContext.ExecContext.
func Exec(ctx context.Context, e Execer, query string, args ...interface{}) (sql.Result, error) {
	if ctx == nil {
		return e.Exec(query, args...)
	}
	return e.ExecContext(ctx, query, args...)
}

// MustExec equals to sqlx.MustExec and sqlx.MustExecContext.
func MustExec(ctx context.Context, e Execer, query string, args ...interface{}) sql.Result {
	if ctx == nil {
		return sqlx.MustExec(e, query, args...)
	}

	return sqlx.MustExecContext(ctx, e, query, args...)
}
