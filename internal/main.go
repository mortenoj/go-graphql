// Package internal is the internal package
package internal

import (
	"github.com/mortenoj/go-graphql-template/internal/config"
	"github.com/mortenoj/go-graphql-template/internal/orm"
	"github.com/mortenoj/go-graphql-template/internal/server"
	"github.com/mortenoj/go-graphql-template/pkg/utils/errutils"
)

// Start the application
func Start(cfg *config.Config) {
	orm, err := orm.Factory(cfg)
	errutils.Must(err)

	server.Run(cfg, orm)
}
