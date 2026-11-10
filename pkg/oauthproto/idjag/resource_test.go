// SPDX-FileCopyrightText: Copyright 2026 Stacklok, Inc.
// SPDX-License-Identifier: Apache-2.0

package idjag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceClaim(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		claim any
		want  []string
	}{
		{name: "string", claim: "https://a.example.com", want: []string{"https://a.example.com"}},
		{
			name:  "array",
			claim: []any{"https://a.example.com", "https://b.example.com"},
			want:  []string{"https://a.example.com", "https://b.example.com"},
		},
		{
			name:  "array with a non-string element skips it",
			claim: []any{"https://a.example.com", 1},
			want:  []string{"https://a.example.com"},
		},
		{name: "empty array", claim: []any{}},
		{name: "nil", claim: nil},
		{name: "number", claim: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, ResourceClaim(tt.claim))
		})
	}
}
