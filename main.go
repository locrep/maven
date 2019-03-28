package main

import (
	"github.com/locrep/locrep-go/config"
	"github.com/locrep/locrep-go/server"
)

func main() {
	conf := config.Config()
	port := conf.Environment.Port()

	server.NewServer(conf).Run(":" + port)
}
