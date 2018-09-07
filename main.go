package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	fmt.Println("grd")
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(ctx, src)
	client := githubv4.NewClient(httpClient)

	var query struct {
		Repository struct {
			Description string
		} `graphql:"repository(owner: \"taxio\", name: \"grd\")"`
	}

	err = client.Query(ctx, &query, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(query.Repository.Description)
}
