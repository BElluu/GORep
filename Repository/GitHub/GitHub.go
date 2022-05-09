package github

type GitHubAccess struct {
	TOKEN string
}

type GitHubResponse struct {
	Title       string
	Repository  string
	Description string
	Link        string
	Author      string
	Reviewers   string
}

var MyGithubAccess = GitHubAccess{}
var myGithubResponse = GitHubResponse{}

func SetGithubToken(token string) *GitHubAccess {
	MyGithubAccess.TOKEN = token
	return &MyGithubAccess
}

func (g GitHubResponse) GetResponseData() *GitHubResponse {
	return &myGithubResponse
}
