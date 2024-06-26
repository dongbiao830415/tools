// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aliases

import (
	"go/token"
	"go/types"
)

// Package aliases defines backward compatible shims
// for the types.Alias type representation added in 1.22.
// This defines placeholders for x/tools until 1.26.

// NewAlias creates a new TypeName in Package pkg that
// is an alias for the type rhs.
//
// When GoVersion>=1.22 and GODEBUG=gotypesalias=1 (or unset),
// the Type() of the return value is a *types.Alias.
func NewAlias(pos token.Pos, pkg *types.Package, name string, rhs types.Type) *types.TypeName {
	if enabled() {
		tname := types.NewTypeName(pos, pkg, name, nil)
		newAlias(tname, rhs)
		return tname
	}
	return types.NewTypeName(pos, pkg, name, rhs)
}
