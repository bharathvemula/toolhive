// SPDX-FileCopyrightText: Copyright 2026 Stacklok, Inc.
// SPDX-License-Identifier: Apache-2.0

package idjag

import "errors"

// Errors returned by [Verify]. Their messages hold no token material.
var (
	// ErrMissingJWTID means the ID-JAG has no jti claim.
	ErrMissingJWTID = errors.New("ID-JAG is missing the required jti claim")
	// ErrMissingClientID means the ID-JAG has no client_id claim.
	ErrMissingClientID = errors.New("ID-JAG is missing the required client_id claim")
	// ErrNoClient means the server has no authenticated or identified client
	// to bind the ID-JAG to.
	ErrNoClient = errors.New("no authenticated or identified client to bind the ID-JAG to")
	// ErrClientMismatch means the client_id claim names another client than
	// the one presenting the ID-JAG.
	ErrClientMismatch = errors.New("ID-JAG was not issued to the requesting client")
)

// Presented is what the redeeming server read from the ID-JAG.
type Presented struct {
	JWTID    string
	ClientID string
}

// Expected is what the server knows about the request.
type Expected struct {
	// ClientID is the ID of the authenticated or identified client. It is
	// empty when the server has no such client.
	ClientID string
}

// Verify checks the two ID-JAG rules a redeeming server applies on top of the
// generic JWT assertion rules: the jti and client_id claims are present
// (draft 3.1 makes both REQUIRED), and client_id names the client that
// presents the ID-JAG (draft 4.4.1). It returns the first failure, in this
// order: [ErrMissingJWTID], [ErrMissingClientID], [ErrNoClient],
// [ErrClientMismatch]. It returns the sentinel itself, never a wrapped error.
// An empty client_id claim can never match an empty expected client, because
// the missing client is refused first.
//
// Verify checks only these things. It does not run the generic RFC 7521 and
// RFC 7523 claim checks, and it does not verify the signature, the issuer
// trust, the audience, the resource, the tenant or replay. The caller does
// those. The jti is what makes single use possible, but this package stores
// nothing, so recording it is the caller's duty.
//
// The draft allows one exception to the client binding rule: a server that
// implements proof of possession for a client that does not authenticate
// (draft 4.4.1 and 9.11) may skip it. ToolHive has no sender-constrained
// ID-JAGs, so Verify has no way to skip it. ToolHive also still accepts a
// public client that is identified only by client_id, which draft 4.4 does
// not allow.
func Verify(p Presented, want Expected) error {
	switch {
	case p.JWTID == "":
		return ErrMissingJWTID
	case p.ClientID == "":
		return ErrMissingClientID
	case want.ClientID == "":
		return ErrNoClient
	case p.ClientID != want.ClientID:
		return ErrClientMismatch
	}
	return nil
}
