package main

import (
	internal "github.com/mortenoj/reko-ring-backend/internal"
	"github.com/mortenoj/reko-ring-backend/internal/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	internal.Start(cfg)
}
