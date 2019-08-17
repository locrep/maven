package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Project struct {
	XMLName xml.Name   `xml:"project"`
	Project []Licenses `xml:"licenses"`
	Organization string   `xml:"organization"`
	ModelVersion string `xml:"modelVersion"`
	GroupId      string `xml:"groupId"`
	ArtifactId   string `xml:"artifactId"`
	Version      string `xml:"version"`
	Name         string `xml:"name"`
	Description  string `xml:"description"`
	Url          string `xml:"url"`
}

type Licenses struct {
	XMLName xml.Name `xml:"licenses"`
	License License  `xml:"license"`
}

type License struct {
	XMLName      xml.Name `xml:"license"`
	Name         string   `xml:"name"`
	Url          string   `xml:"url"`
	Distribution string   `xml:"distribution"`
}

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("smart.pom")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened pom")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array

	var project Project

	xml.Unmarshal(byteValue, &project)

	for i := 0; i < len(project.Project); i++ {

		fmt.Println("project Name " + project.Name)
		fmt.Println("project groupId =" + project.GroupId)
		fmt.Println("project artifactId = " + project.ArtifactId)
		fmt.Println("project Version =" + project.Version)
		fmt.Println("Organization =" + project.Organization)
		fmt.Println("license " + project.Project[i].License.Name)
		fmt.Println("project description =" + project.Description)
		fmt.Println("project url =" + project.Url)
	}
}
