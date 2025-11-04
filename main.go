// The Swiss Army Knife for examining text files.
//
// Copyright 2025 Christian Uhsat. All rights reserved.
// Use of this source code is governed by the GPL-3.0
// license that can be found in the LICENSE.md file.
//
// For more information, please consult:
//
//	forensic-examiner.eu
package main

import (
	"github.com/inconshreveable/mousetrap"

	"github.com/cuhsat/fox/v3/internal/cmd"
	"github.com/cuhsat/fox/v3/internal/pkg/sys"
)

// Main start and catch.
func main() {
	defer sys.Recover()

	if mousetrap.StartedByExplorer() {
		sys.Trap()
	} else {
		sys.Init()
		_ = cmd.Fox.Execute()
	}
}
