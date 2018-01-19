package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	dac "github.com/akshaykarle/go-http-digest-auth-client"
	"github.com/akshaykarle/mongodb-atlas-go/mongodb"
)

func main() {
	username := os.Args[1]
	password := os.Args[2]
	// orgId := os.Args[3]
	t := dac.NewTransport(username, password)
	httpClient := &http.Client{Transport: &t}
	client := mongodb.NewClient(httpClient)

	// Projects.List example
	projects, _, err := client.Projects.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("projects list: %v\n", projects)

	// Projects.Get example
	project, _, err := client.Projects.Get(projects[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("project get: %v\n", project)

	// Projects.Create example
	params := &mongodb.Project{
		OrgID: orgId,
		Name:  "test",
	}
	project, _, err = client.Projects.Create(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("project created: %v\n", project)

	// Projects.GetByName example
	project, _, err = client.Projects.GetByName("test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("project get: %v\n", project)
}
