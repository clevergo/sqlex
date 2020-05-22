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

func TestNamedExec(t *testing.T) {
	cases := []struct {
		ctx context.Context
		id  string
	}{
		{nil, "namedexec"},
		{context.Background(), "namedexec_context"},
	}
	for _, test := range cases {
		_, err := NamedExec(test.ctx, mockDBx, fmt.Sprintf("INSERT INTO %s(id) VALUES (:id)", mockTable), map[string]interface{}{
			"id": test.id,
		})
		assert.Nil(t, err)
		assert.True(t, isMockRowExists(test.id))
	}
}

func TestNamedQuery(t *testing.T) {
	cases := []struct {
		ctx context.Context
		id  string
	}{
		{nil, "foo"},
		{context.Background(), "bar"},
	}
	for _, test := range cases {
		_, err := NamedQuery(test.ctx, mockDBx, fmt.Sprintf("SELECT id FROM %s WHERE id=:id", mockTable), map[string]interface{}{
			"id": test.id,
		})
		assert.Nil(t, err)
	}
}
