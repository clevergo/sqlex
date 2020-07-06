// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package sqlex

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var (
	// ensurces *sql.Tx and *sqlx.Tx implement DB.
	_ DB  = &sql.DB{}
	_ DBx = &sqlx.DB{}
)

// DB is an interface for sql.DB and sqlx.DB.
type DB interface {
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// DBx is an interface for sqlx.DB.
type DBx interface {
	DB
	Beginx() (*sqlx.Tx, error)
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

// Transact begins a transaction, and then invokes fn.
// Rollback automatically if receives a panic or error.
func Transact(db DB, fn func(tx *sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = fn(tx)
	return err
}

// TransactContext begins a transaction, and then invokes fn.
// Rollback automatically if receives a panic or error.
func TransactContext(ctx context.Context, opts *sql.TxOptions, db DB, fn func(ctx context.Context, tx *sql.Tx) error) (err error) {
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = fn(ctx, tx)
	return err
}

// Transactx begins a transaction, and then invokes fn.
// Rollback automatically if receives a panic or error.
func Transactx(db DBx, fn func(tx *sqlx.Tx) error) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = fn(tx)
	return err
}

// TransactContextx begins a transaction, and then invokes fn.
// Rollback automatically if receives a panic or error.
func TransactContextx(ctx context.Context, opts *sql.TxOptions, db DBx, fn func(ctx context.Context, tx *sqlx.Tx) error) (err error) {
	tx, err := db.BeginTxx(ctx, opts)
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = fn(ctx, tx)
	return err
}
