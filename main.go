package main

import (
	"github.com/locrep/go/config"
	"github.com/locrep/go/server"
)

func main() {
	conf := config.Config()
	port := conf.Environment.Port()

	server.NewServer(conf).Run(":" + port)
}
