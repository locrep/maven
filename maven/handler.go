package maven

import (
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/locrep/maven/config"
)

const Repo = "./maven-repo"

type handler struct {
	config config.Conf
}

// EmptyString is a constant value used to improve readability when a function needs to return empty strings.
const EmptyString = ""

func UrlResolve(param string) (*ArtifactRequest, error) {
	r, _ := regexp.Compile(`([A-z0-9-_/.]+)\/([A-z0-9-_.]+)\/([0-9.]+[A-z0-9-_.]*)\/([A-z0-9-_.]+)`)
	match := r.FindStringSubmatch(param)

	res := new(ArtifactRequest)
	res.Artifact = new(Artifact)
	res.Artifact.Version = match[3]
	res.Artifact.ArtifactID = match[2]
	res.Artifact.GroupID = match[1]
	res.File = match[4]

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
