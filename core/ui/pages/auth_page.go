package pages

import (
	"github.com/OliverKeefe/git-cleanse/core/repos"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type AuthMsg struct {
	Success      bool
	Platform     string
	GitlabClient *gitlab.Client
	User         *repos.GitlabUser
	Error        error
}
