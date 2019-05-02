package maven

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/locrep-go/config"
	"github.com/parnurzeal/gorequest"
	"log"
	"os"
	"strings"
)

type handler struct {
	config config.Conf
}

func (h handler) Handle(ctx *gin.Context) {
	for _, repo := range config.Config().MavenRepos {
		log.Println(repo + ctx.Request.URL.String())

		filePath := "./maven-repo" + ctx.Request.URL.String()

		if _, err := os.Stat(filePath); os.IsNotExist(err) {

			_, body, _ := gorequest.New().Get(repo + ctx.Request.URL.String()).EndBytes()
			//todo error handling
			//todo status code handling

			var (
				file *os.File
				err  error
			)

			paths := strings.Split(filePath, "/")
			fileName := paths[len(paths)-1]
			folder := filePath[0 : len(filePath)-len(fileName)]

			if err := os.MkdirAll(folder, 0777); err!=nil{
				log.Println(err.Error())
			}

			if file, err = os.Create(filePath); err != nil {
				log.Println(err.Error())
			}
			defer file.Close()

			if _, err := file.Write(body); err != nil {
				log.Println(err.Error())
			}
			//todo error handling
			//todo: check sha and md5
		}

		ctx.File(filePath)
	}

}

func NewHandler(config config.Conf) handler {
	return handler{config: config}
}
