package main

import (
	"fmt"
	sdk "github.com/gaia-pipeline/gosdk"
	"log"
	"os"
	"os/exec"
)

const (
	BuildCMD = "mkdir -p %s/bin && go version && go build -o %s/bin/traffic_service %s/cmd/traffic_service.go"
)

func Build(args sdk.Arguments) error {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	command := fmt.Sprintf(BuildCMD, pwd, pwd, pwd)
	log.Println(command)
	cmd := exec.Command(command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("result =>", string(output))
	return nil
}

func main() {
	jobs := sdk.Jobs{
		sdk.Job{
			Handler:     Build,
			Title:       "Build TMZ",
			Description: "Build TMZ.",
		},
	}

	// Serve
	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
