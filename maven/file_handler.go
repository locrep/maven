package maven

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/parnurzeal/gorequest"
)

var savePath = "./maven-repo"

const CentralRepo = "https://repo.maven.apache.org"

func getArtifact(req *ArtifactRequest) (string, []error) {
	var (
		filePath string
		errs     []error
	)
	filePath, errs = artifactExists(req)
	if filePath != EmptyString && errs == nil {
		return filePath, nil
	}
	if os.IsNotExist(errs[0]) {
		filePath, errs = fetchArtifact(req)
		if filePath == EmptyString || errs != nil {
			return EmptyString, errs
		}
		return filePath, nil
	}
	return EmptyString, errs
}

func artifactExists(req *ArtifactRequest) (string, []error) {
	var folder = fmt.Sprintf("%s/maven2/%s/%s/%s/", savePath, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version)
	var filePath = fmt.Sprintf("%s%s", folder, req.File)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return EmptyString, []error{err}
	}
	return filePath, nil
}

func fetchArtifact(req *ArtifactRequest) (string, []error) {
	var folder = fmt.Sprintf("%s/maven2/%s/%s/%s/", savePath, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version)
	var filePath = fmt.Sprintf("%s%s", folder, req.File)
	var uri = fmt.Sprintf("%s/maven2/%s/%s/%s/%s", CentralRepo, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version, req.File)

	response, body, errs := gorequest.New().Get(uri).EndBytes()
	if len(errs) > 0 || response.StatusCode != http.StatusOK {
		return EmptyString, errs
	}

	var (
		file *os.File
		err  error
	)

	if err := os.MkdirAll(folder, 0755); err != nil {
		return EmptyString, []error{err}
	}

	if file, err = os.Create(filePath); err != nil {
		return EmptyString, []error{err}
	}

	defer file.Close()

	if _, err := file.Write(body); err != nil {
		return EmptyString, []error{err}
	}

	return filePath, nil
}
