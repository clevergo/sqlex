// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package sqlex

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	cases := []struct {
		ctx context.Context
		id  string
	}{
		{nil, "exec"},
		{context.Background(), "exec_context"},
	}
	for _, test := range cases {
		_, err := Exec(test.ctx, mockDBx, fmt.Sprintf("INSERT INTO %s(id) VALUES (?)", mockTable), test.id)
		assert.Nil(t, err)
		assert.True(t, isMockRowExists(test.id))
	}
}

func TestMustExec(t *testing.T) {
	cases := []struct {
		ctx context.Context
		id  string
	}{
		{nil, "mustexec"},
		{context.Background(), "mustexec_context"},
	}
	for _, test := range cases {
		MustExec(test.ctx, mockDBx, fmt.Sprintf("INSERT INTO %s(id) VALUES (?)", mockTable), test.id)
		assert.True(t, isMockRowExists(test.id))
	}
}
