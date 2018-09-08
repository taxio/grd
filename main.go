package main

import (
	"github.com/spf13/cobra"
	"github.com/taxio/grd/cmd"
)

var Version string = "v0.0.1"
var showCmd = &cobra.Command{}

func main() {
	cmd.Execute()
	// fmt.Println("grd", Version)
	// ctx := context.Background()
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// src := oauth2.StaticTokenSource(
	// 	&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	// )
	// httpClient := oauth2.NewClient(ctx, src)
	// client := githubv4.NewClient(httpClient)

	// var query struct {
	// 	Repository struct {
	// 		Description string
	// 	} `graphql:"repository(owner: \"taxio\", name: \"grd\")"`
	// }

	// err = client.Query(ctx, &query, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(query.Repository.Description)
}
