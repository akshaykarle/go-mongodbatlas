package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	dac "github.com/akshaykarle/go-http-digest-auth-client"
	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
)

func main() {
	username := os.Args[1]
	password := os.Args[2]
	orgID := os.Args[3]
	t := dac.NewTransport(username, password)
	httpClient := &http.Client{Transport: &t}
	client := ma.NewClient(httpClient)

	// Projects.List example
	projects, _, err := client.Projects.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("projects list: \n%v\n\n", projects)

	// Projects.Get example
	project, _, err := client.Projects.Get(projects[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("project get: \n%v\n\n", project)

	// Projects.Create example
	params := &ma.Project{
		OrgID: orgID,
		Name:  "test",
	}
	project, _, err = client.Projects.Create(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("project created: \n%v\n\n", project)

	// Projects.GetByName example
	project, _, err = client.Projects.GetByName("test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("project get: \n%v\n\n", project)

	// Projects.Delete example
	_, err = client.Projects.Delete(project.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("project deleted: \n%v\n\n", project)
}
