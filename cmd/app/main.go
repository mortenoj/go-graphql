package main

import (
	internal "github.com/mortenoj/go-graphql-template/internal"
	"github.com/mortenoj/go-graphql-template/internal/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	internal.Start(cfg)
}
