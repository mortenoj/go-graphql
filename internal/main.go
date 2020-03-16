// Package internal is the internal package
package internal

import (
	"github.com/mortenoj/reko-ring-backend/internal/config"
	"github.com/mortenoj/reko-ring-backend/internal/orm"
	"github.com/mortenoj/reko-ring-backend/internal/server"
	"github.com/mortenoj/reko-ring-backend/pkg/utils/errutils"
)

// Start the application
func Start(cfg *config.Config) {
	orm, err := orm.Factory(cfg)
	errutils.Must(err)

	server.Run(cfg, orm)
}
