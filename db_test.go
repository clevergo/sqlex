// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package sqlex

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB    *sql.DB
	mockDBx   *sqlx.DB
	mockTable = "sqlex_test"
)

func TestMain(m *testing.M) {
	var dbPassword string
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		dbPassword = ":" + v
	}
	dbName := "test"
	if v := os.Getenv("DB_NAME"); v != "" {
		dbName = v
	}
	mockDBx = sqlx.MustOpen("mysql", fmt.Sprintf("root%s@/%s", dbPassword, dbName))
	defer mockDBx.Close()
	mockDBx.MustExec(fmt.Sprintf("DROP TABLE IF EXISTS %s", mockTable))
	mockDBx.MustExec(fmt.Sprintf("CREATE TABLE %s(id varchar(36) not null, PRIMARY KEY(id))", mockTable))
	mockDB = mockDBx.DB

	m.Run()
}

func insertMockRow(tx *sql.Tx, id string) error {
	_, err := tx.Exec(fmt.Sprintf("INSERT INTO %s VALUES(?)", mockTable), id)
	return err
}

func insertMockRowx(tx *sqlx.Tx, id string) error {
	_, err := tx.Exec(fmt.Sprintf("INSERT INTO %s VALUES(?)", mockTable), id)
	return err
}

func isMockRowExists(id string) bool {
	var v string
	mockDBx.Get(&v, fmt.Sprintf("SELECT id FROM %s WHERE id=?", mockTable), id)
	return v == id
}

func TestTransact(t *testing.T) {
	err := Transact(mockDB, func(tx *sql.Tx) error {
		assert.Nil(t, insertMockRow(tx, "transact"))
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, isMockRowExists("transact"))

	expectedErr := errors.New("error from transaction")
	err = Transact(mockDB, func(tx *sql.Tx) error {
		assert.Nil(t, insertMockRow(tx, "transact_error"))
		return expectedErr
	})
	assert.Equal(t, expectedErr, err)
	assert.False(t, isMockRowExists("transact_error"))

	// panic
	assert.Panics(t, func() {
		Transact(mockDB, func(tx *sql.Tx) error {
			assert.Nil(t, insertMockRow(tx, "transact_panic"))
			panic("transact_panic")
		})
	})
	assert.False(t, isMockRowExists("transact_panic"))
}

func TestTransactContext(t *testing.T) {
	err := TransactContext(context.Background(), nil, mockDB, func(ctx context.Context, tx *sql.Tx) error {
		assert.Nil(t, insertMockRow(tx, "transact_context"))
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, isMockRowExists("transact_context"))

	expectedErr := errors.New("error from transaction")
	err = TransactContext(context.Background(), nil, mockDB, func(ctx context.Context, tx *sql.Tx) error {
		assert.Nil(t, insertMockRow(tx, "transact_context_error"))
		return expectedErr
	})
	assert.Equal(t, expectedErr, err)
	assert.False(t, isMockRowExists("transact_context_error"))

	// panic
	assert.Panics(t, func() {
		TransactContext(context.Background(), nil, mockDB, func(ctx context.Context, tx *sql.Tx) error {
			assert.Nil(t, insertMockRow(tx, "transact_context_panic"))
			panic("transact_context_panic")
		})
	})
	assert.False(t, isMockRowExists("transact_context_panic"))
}

func TestTransactx(t *testing.T) {
	err := Transactx(mockDBx, func(tx *sqlx.Tx) error {
		assert.Nil(t, insertMockRowx(tx, "transactx"))
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, isMockRowExists("transactx"))

	expectedErr := errors.New("error from transaction")
	err = Transactx(mockDBx, func(tx *sqlx.Tx) error {
		assert.Nil(t, insertMockRowx(tx, "transactx_error"))
		return expectedErr
	})
	assert.Equal(t, expectedErr, err)
	assert.False(t, isMockRowExists("transactx_error"))

	// panic
	assert.Panics(t, func() {
		Transactx(mockDBx, func(tx *sqlx.Tx) error {
			assert.Nil(t, insertMockRowx(tx, "transactx_panic"))
			panic("transactx_panic")
		})
	})
	assert.False(t, isMockRowExists("transactx_panic"))
}

func TestTransactContextx(t *testing.T) {
	err := TransactContextx(context.Background(), nil, mockDBx, func(ctx context.Context, tx *sqlx.Tx) error {
		assert.Nil(t, insertMockRowx(tx, "transact_contextx"))
		return nil
	})
	assert.Nil(t, err)
	assert.True(t, isMockRowExists("transact_contextx"))

	expectedErr := errors.New("error from transaction")
	err = TransactContextx(context.Background(), nil, mockDBx, func(ctx context.Context, tx *sqlx.Tx) error {
		assert.Nil(t, insertMockRowx(tx, "transact_contextx_error"))
		return expectedErr
	})
	assert.Equal(t, expectedErr, err)
	assert.False(t, isMockRowExists("transact_contextx_error"))

	// panic
	assert.Panics(t, func() {
		TransactContextx(context.Background(), nil, mockDBx, func(ctx context.Context, tx *sqlx.Tx) error {
			assert.Nil(t, insertMockRowx(tx, "transact_contextx_panic"))
			panic("transact_contextx_panic")
		})
	})
	assert.False(t, isMockRowExists("transact_contextx_panic"))
}
