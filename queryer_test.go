// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package sqlex

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	cases := []struct {
		ctx context.Context
	}{
		{nil},
		{context.Background()},
	}
	for _, test := range cases {
		_, err := Query(test.ctx, mockDBx, fmt.Sprintf("SELECT id FROM %s WHERE id=?", mockTable), "foo")
		assert.Nil(t, err)
	}
}

func TestQueryx(t *testing.T) {
	cases := []struct {
		ctx context.Context
	}{
		{nil},
		{context.Background()},
	}
	for _, test := range cases {
		_, err := Queryx(test.ctx, mockDBx, fmt.Sprintf("SELECT id FROM %s WHERE id=?", mockTable), "foo")
		assert.Nil(t, err)
	}
}

func TestQueryRowx(t *testing.T) {
	cases := []struct {
		ctx context.Context
	}{
		{nil},
		{context.Background()},
	}
	for _, test := range cases {
		var id string
		row := QueryRowx(test.ctx, mockDBx, fmt.Sprintf("SELECT id FROM %s WHERE id=?", mockTable), "foo")
		row.Scan(&id)
		assert.Equal(t, "foo", id)
	}
}

func TestQueryGet(t *testing.T) {
	cases := []struct {
		ctx context.Context
	}{
		{nil},
		{context.Background()},
	}
	for _, test := range cases {
		var id string
		Get(test.ctx, mockDBx, &id, fmt.Sprintf("SELECT id FROM %s WHERE id=?", mockTable), "foo")
		assert.Equal(t, "foo", id)
	}
}

func TestQuerySelect(t *testing.T) {
	cases := []struct {
		ctx context.Context
	}{
		{nil},
		{context.Background()},
	}
	for _, test := range cases {
		var rows []struct {
			ID string `db:"id"`
		}
		Select(test.ctx, mockDBx, &rows, fmt.Sprintf("SELECT id FROM %s WHERE id IN (?, ?)", mockTable), "foo", "bar")
		assert.Len(t, rows, 2)
		for _, row := range rows {
			assert.True(t, row.ID == "foo" || row.ID == "bar")
		}
	}
}
