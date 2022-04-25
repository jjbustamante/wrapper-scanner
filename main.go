package main

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"os/exec"
	"strings"
)

func main()  {

	repo, err := name.NewRepository("buildpacksio/lifecycle")
	if err != nil {
		panic(err)
	}

	allTags, err := remote.List(repo)
	if err != nil {
		panic(err)
	}

	versionTags := []string{}
	for _, tag := range allTags {
		if validTag(tag) {
			versionTags = append(versionTags, tag)
		}
	}

	for _, version := range versionTags {
		image := fmt.Sprintf("buildpacksio/lifecycle:%s", version)
		fmt.Println("****************************************\n")
		fmt.Printf("Scanning Lifecycle Image: %s\n", image)
		cmd := exec.Command("trivy", "image", image)

		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(stdout))
		fmt.Println("****************************************\n")
	}
}

func validTag(tag string) bool {
	return strings.HasPrefix(tag, "0.") && !strings.Contains(tag, "linux") && !strings.Contains(tag, "windows")
}
