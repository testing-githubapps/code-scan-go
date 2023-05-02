package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

var (
	organization = flag.String("org", "", "organization name")
)

func main() {
	required := []string{"org"}
	flag.Parse()
	for _, req := range required {
		if flag.Lookup(req).Value.String() == "" {
			log.Fatalf("missing required -%s argument/flag\n", req)
		}
	}

	opts := api.ClientOptions{
		AuthToken:   os.Getenv("GITHUB_TOKEN"),
		EnableCache: true,
		Timeout:     5 * time.Second,
		Log:         os.Stdout,
	}
	client, err := api.NewGraphQLClient(opts)
	if err != nil {
		log.Fatal(err)
	}
	var query struct {
		Repository struct {
			Refs struct {
				Nodes []struct {
					Name string
				}
			} `graphql:"refs(refPrefix: $refPrefix, last: $last)"`
			Languages struct {
				Nodes []struct {
					Name string
				}
			} `graphql:"languages(first: $first)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	variables := map[string]interface{}{
		"refPrefix": graphql.String("refs/heads/"),
		"first":     graphql.Int(10),
		"last":      graphql.Int(30),
		"owner":     graphql.String(*organization),
		"name":      graphql.String("importer-labs"),
	}
	err = client.Query("Repository", &query, variables)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(query)
}
