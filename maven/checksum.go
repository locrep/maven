package maven

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/parnurzeal/gorequest"
)

//md5File is a function that calculates and returns the md5 checksum of the file at the given path.
func md5File(path string) (string, error) {
	f, err := os.Open(path)

	if err != nil {
		return EmptyString, err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return EmptyString, err
	}

	checksum := fmt.Sprintf("%x", h.Sum(nil))
	return checksum, nil

}

//md5File is a function that calculates and returns the sha1 checksum of the file at the given path.
func sha1File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return EmptyString, err
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return EmptyString, err
	}

	checksum := fmt.Sprintf("%x", h.Sum(nil))
	return checksum, nil
}

func verifyMd5(req *ArtifactRequest) (bool, []error) {
	var folder = fmt.Sprintf("%s/maven2/%s/%s/%s/", savePath, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version)
	var filePath = fmt.Sprintf("%s%s", folder, req.File)
	uri := fmt.Sprintf("%s/maven2/%s/%s/%s/%s.md5", CentralRepo, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version, req.File)

	response, body, errs := gorequest.New().Get(uri).EndBytes()
	if len(errs) > 0 || response.StatusCode != http.StatusOK {
		return false, errs
	}

	fetchedChecksum := string(body)

	calculatedChecksum, err := md5File(filePath)
	if err != nil {
		return false, []error{err}
	}

	if calculatedChecksum == fetchedChecksum {
		return true, nil
	}

	return false, nil

}

func verifySha1(req *ArtifactRequest) (bool, []error) {
	var folder = fmt.Sprintf("%s/maven2/%s/%s/%s/", savePath, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version)
	var filePath = fmt.Sprintf("%s%s", folder, req.File)
	uri := fmt.Sprintf("%s/maven2/%s/%s/%s/%s.sha1", CentralRepo, strings.ReplaceAll(req.Artifact.GroupID, ".", "/"), req.Artifact.ArtifactID, req.Artifact.Version, req.File)

	response, body, errs := gorequest.New().Get(uri).EndBytes()
	if len(errs) > 0 || response.StatusCode != http.StatusOK {
		return false, errs
	}

	fetchedChecksum := string(body)

	calculatedChecksum, err := sha1File(filePath)
	if err != nil {
		return false, []error{err}
	}

	if calculatedChecksum == fetchedChecksum {
		return true, nil
	}

	return false, nil

}
