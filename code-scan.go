package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v52/github"
)

func main() {
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "testing-githubapps", opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(repos)
}
