package cmd

import (
	github "GORep/Repository/GitHub"
	"flag"
)

func Init() {
	var githubToken string

	flag.StringVar(&githubToken, "github_token", "", "Access Token to Github")
	flag.Parse()
	if githubToken != "" {
		github.SetGithubToken(githubToken)
	}
}
