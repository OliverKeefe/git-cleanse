package auth

import (
	"context"
	"log"

	"github.com/google/go-github/v61/github"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/oauth2"
)

type Accounts struct {
	Github    []string
	GitLab    []string
	BitBucket []string
	LocalGit  []string
}

func GetGitHubClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func GetGitLabClient(token string) (*gitlab.Client, error) {
	git, err := gitlab.NewClient(token, gitlab.WithBaseURL("https://cseegit.essex.ac.uk/"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	} else {
		log.Println("Success!")
	}

	return git, err
}
