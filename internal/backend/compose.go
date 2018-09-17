package backend

//import (
//	"github.com/docker/libcompose/docker"
//	"github.com/docker/libcompose/docker/ctx"
//	"github.com/docker/libcompose/project"
//	"github.com/docker/libcompose/project/options"
//	"golang.org/x/net/context"
//	"log"
//)
//
//func RunComposeFile() {
//	composeFile := "test/resources/docker-compose-ubuntu.yml"
//	project, err := docker.NewProject(&ctx.Context{
//		Context: project.Context{
//			ComposeFiles: []string{composeFile},
//			ProjectName:  "my-compose",
//		},
//	}, nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = project.Up(context.Background(), options.Up{})
//
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func Kill() {
//	// kill started compose container for project.
//}
//
//func GenerateProjectId() {
//	// create a reproducible project id based on
//	// - camunda version
//	// - server and server version
//	// - database and database version
//}
