package main

import (
	"github.com/locrep/maven/config"
	"github.com/locrep/maven/server"
)

func main() {
	conf := config.Config()
	port := conf.Environment.Port()

	server.NewServer(conf).Run(":" + port)
}
