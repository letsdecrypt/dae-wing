/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, daeuniverse Organization <team@v2raya.org>
 */

package main

import (
	"github.com/json-iterator/go/extra"
	"github.com/letsdecrypt/dae-wing/cmd"
	"os"
)

import (
	_ "github.com/daeuniverse/dae/component/outbound"
)

func main() {
	extra.RegisterFuzzyDecoders()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
