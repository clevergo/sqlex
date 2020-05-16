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
	// ensure sqlx.DB and sqlx.Tx implement Ext interface.
	_ Ext = &sqlx.DB{}
	_ Ext = &sqlx.Tx{}
)

// Ext is an interface that wraps sqlx.Ext and sqlx.ExtContext.
type Ext interface {
	Execer
	Queryer
	DriverName() string
	Rebind(string) string
	BindNamed(string, interface{}) (string, []interface{}, error)
}

// NamedExec equals to sqlx.NamedExec and sqlx.NamedExecContext.
func NamedExec(ctx context.Context, e Ext, query string, arg interface{}) (sql.Result, error) {
	if ctx == nil {
		return sqlx.NamedExec(e, query, arg)
	}
	return sqlx.NamedExecContext(ctx, e, query, arg)
}

// NamedQuery equals to sqlx.NamedQuery and sqlx.NamedQueryContext.
func NamedQuery(ctx context.Context, e Ext, query string, arg interface{}) (*sqlx.Rows, error) {
	if ctx == nil {
		return sqlx.NamedQuery(e, query, arg)
	}
	return sqlx.NamedQueryContext(ctx, e, query, arg)
}
