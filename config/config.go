package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const MavenReposYaml string = "./config/maven_repos.yaml"

type Conf struct {
	Environment env
	MavenRepos  []string
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

	mvnRepos := make([]string, 0)
	mavenReposYml, err := ioutil.ReadFile(MavenReposYaml)
	if err != nil {
		log.Printf("Reading "+MavenReposYaml+" failed: #%v ", err)
	}
	err = yaml.Unmarshal(mavenReposYml, &mvnRepos)
	if err != nil {
		log.Fatalf("Unmarshalling "+MavenReposYaml+" failed: #%v ", err)
	}

	return Conf{
		Environment: environment,
		MavenRepos:  mvnRepos,
	}
}
