package config

import "os"

type Conf struct {
	Environment env
}

type env struct {
	Port      func() string
	BuildMode func() string
}

func Config() Conf {
	environment := env{
		Port: func() string {
			return os.Getenv("PORT")
		},
		BuildMode: func() string {
			return os.Getenv("BUILD_MODE")
		},
	}

	return Conf{
		Environment: environment,
	}
}
