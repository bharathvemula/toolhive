// SPDX-FileCopyrightText: Copyright 2026 Stacklok, Inc.
// SPDX-License-Identifier: Apache-2.0

package idjag

// ResourceClaim returns the values of an ID-JAG resource claim, which is a
// string or an array of strings (draft 3.1). A string gives a one-element
// slice. An array gives its string elements, and elements that are not strings
// are skipped. Any other value, including nil, gives nil.
func ResourceClaim(claim any) []string {
	switch v := claim.(type) {
	case string:
		return []string{v}
	case []any:
		var out []string
		for _, e := range v {
			if s, ok := e.(string); ok {
				out = append(out, s)
			}
		}
		return out
	}
	return nil
}
