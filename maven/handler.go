package maven

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/locrep/go/config"
)

const Repo = "./maven-repo"

type handler struct {
	config config.Conf
}

func UrlResolve(param string) (*ArtifactRequest, error) {
	// TODO: Optimize the function (use variables for indexes, dont recalculate)
	var slashLastIndex = strings.LastIndex(param, "/")
	if slashLastIndex == -1 {
		return nil, errors.New("slashIndex failed")
	}
	var fileName = param[slashLastIndex+1:]
	var dashIndex = strings.LastIndex(fileName, "-")
	if dashIndex == -1 {
		return nil, errors.New("dashIndex failed")
	}
	var dotIndex = strings.LastIndex(fileName, ".")
	if dotIndex == -1 {
		return nil, errors.New("dotIndex failed")
	}
	var trimmed = param[:slashLastIndex]
	trimmed = trimmed[:strings.LastIndex(trimmed, fileName[dashIndex+1:dotIndex])]
	var grid = trimmed[:strings.LastIndex(trimmed, fileName[:dashIndex])-1]
	res := new(ArtifactRequest)
	res.Artifact = new(Artifact)
	res.Artifact.Version = fileName[dashIndex+1 : dotIndex]
	res.Artifact.ArtifactID = fileName[:dashIndex]
	res.Artifact.GroupID = grid
	res.File = fileName
	return res, nil
}

func NewHandler(config config.Conf) handler {
	return handler{config: config}
}

func (h handler) Handle(ctx *gin.Context) {
	// TODO: Fix the error handling and update the url resolver to serve index of files
	url := ctx.Request.URL.String()
	log.Println(url[:7])
	if url[:7] == "/maven2" {
		res, err := UrlResolve(url[8:])
		if err != nil {
			log.Println("Error resolve")
			ctx.JSON(404, DependencyFetchError(err))
		} else {
			filePath, errs := getArtifact(res)
			if len(errs) == 0 {
				ctx.File(filePath)
			} else {
				log.Println("Error fetch")
				ctx.JSON(404, DependencyFetchError(errs[0]))
			}

		}
	}

}
