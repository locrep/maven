package maven

import "fmt"

type Artifact struct {
	GroupID    string
	ArtifactID string
	Version    string
}

func (artif Artifact) String() string {
	return fmt.Sprintf("GroupID: %s -/- ArtifactID: %s -/- Version: %s", artif.GroupID, artif.ArtifactID, artif.Version)
}
