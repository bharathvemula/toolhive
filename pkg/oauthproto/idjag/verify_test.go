// SPDX-FileCopyrightText: Copyright 2026 Stacklok, Inc.
// SPDX-License-Identifier: Apache-2.0

package idjag

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerify(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		p       Presented
		want    Expected
		wantErr error
		wantMsg string
	}{
		{
			name: "ok",
			p:    Presented{JWTID: "jti-1", ClientID: "client-a"},
			want: Expected{ClientID: "client-a"},
		},
		{
			name:    "jti and client_id missing reports jti first",
			p:       Presented{},
			want:    Expected{ClientID: "client-a"},
			wantErr: ErrMissingJWTID,
			wantMsg: "ID-JAG is missing the required jti claim",
		},
		{
			name:    "jti missing with no client reports jti first",
			p:       Presented{ClientID: "client-a"},
			want:    Expected{},
			wantErr: ErrMissingJWTID,
			wantMsg: "ID-JAG is missing the required jti claim",
		},
		{
			name:    "client_id missing",
			p:       Presented{JWTID: "jti-1"},
			want:    Expected{ClientID: "client-a"},
			wantErr: ErrMissingClientID,
			wantMsg: "ID-JAG is missing the required client_id claim",
		},
		{
			name:    "client_id missing with no client reports client_id first",
			p:       Presented{JWTID: "jti-1"},
			want:    Expected{},
			wantErr: ErrMissingClientID,
			wantMsg: "ID-JAG is missing the required client_id claim",
		},
		{
			name:    "no client",
			p:       Presented{JWTID: "jti-1", ClientID: "client-a"},
			want:    Expected{},
			wantErr: ErrNoClient,
			wantMsg: "no authenticated or identified client to bind the ID-JAG to",
		},
		{
			name:    "mismatch",
			p:       Presented{JWTID: "jti-1", ClientID: "client-a"},
			want:    Expected{ClientID: "client-b"},
			wantErr: ErrClientMismatch,
			wantMsg: "ID-JAG was not issued to the requesting client",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := Verify(tt.p, tt.want)
			if tt.wantErr == nil {
				require.NoError(t, err)
				return
			}
			require.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.wantMsg, err.Error())
		})
	}
}
