package maven

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/locrep/go/config"
	"github.com/parnurzeal/gorequest"
)

var savePath = "./maven-repo"

func getArtifact(req *ArtifactRequest) (string, []error) {
	var (
		filePath string
		errs     []error
	)
	filePath, errs = artifactExists(req)
	if filePath != "" && errs == nil {
		return filePath, nil
	}
	if os.IsNotExist(errs[0]) {
		filePath, errs = fetchArtifact(req)
		if filePath == "" || errs != nil {
			return "", errs
		}
		return filePath, nil
	}
	return "", errs
}

func artifactExists(req *ArtifactRequest) (string, []error) {
	var folder = fmt.Sprintf("%s/maven2/%s/%s/%s/", savePath, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version)
	var filePath = fmt.Sprintf("%s%s", folder, req.File)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", []error{err}
	}
	return filePath, nil
}

func fetchArtifact(req *ArtifactRequest) (string, []error) {
	var folder = fmt.Sprintf("%s/maven2/%s/%s/%s/", savePath, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version)
	var filePath = fmt.Sprintf("%s%s", folder, req.File)
	for _, repo := range config.Config().MavenRepos {
		var uri = fmt.Sprintf("%s/maven2/%s/%s/%s", repo, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version)

		response, body, errs := gorequest.New().Get(uri).EndBytes()
		if len(errs) > 0 || response.StatusCode != http.StatusOK {
			return "", errs
		}

		var (
			file *os.File
			err  error
		)

		if err := os.MkdirAll(folder, 0755); err != nil {
			return "", []error{err}
		}

		if file, err = os.Create(filePath); err != nil {
			return "", []error{err}
		}

		defer file.Close()

		if _, err := file.Write(body); err != nil {
			return "", []error{err}
		}

	}
	return filePath, nil
}
