/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, daeuniverse Organization <team@v2raya.org>
 */

package config

import (
	daeConfig "github.com/daeuniverse/dae/config"
	"github.com/graph-gophers/graphql-go"
	"github.com/letsdecrypt/dae-wing/common"
	"github.com/letsdecrypt/dae-wing/db"
	"github.com/letsdecrypt/dae-wing/graphql/service/config/global"
)

type Resolver struct {
	DaeGlobal *daeConfig.Global
	Model     *db.Config
}

func (r *Resolver) ID() graphql.ID {
	return common.EncodeCursor(r.Model.ID)
}

func (r *Resolver) Name() string {
	return r.Model.Name
}

func (r *Resolver) Global() *global.Resolver {
	return &global.Resolver{
		Global: r.DaeGlobal,
	}
}

func (r *Resolver) Selected() bool {
	return r.Model.Selected
}
