package auth

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
