package maven

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/locrep-go/config"
	"github.com/parnurzeal/gorequest"
)

type handler struct {
	config config.Conf
}

func (h handler) Handle(ctx *gin.Context) {
	for _, repo := range config.Config().MavenRepos {
		println(repo + ctx.Request.URL.String())
		gorequest.New().Get(repo + ctx.Request.URL.String()).End()
		//todo: check sha and md5
		//todo: return expected maven response
		//todo: download artifacts
	}

}

func NewHandler(config config.Conf) handler {
	return handler{config: config}
}
