package pages

import (
	"fmt"
	_ "github.com/OliverKeefe/git-cleanse/core/repositories"
	"github.com/OliverKeefe/git-cleanse/core"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v6"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type RepoItem struct {
	Title       string
	Description string
}

type RepoPageModel struct {
	Cursor             int
	ListOfRepositories []RepoItem
	Repository         []*git.Repository
	Selected           string
	Width              int
	Height             int
}

func NewRepoPageModel(path string, uri string, provider string) (RepoPageModel, error) {
	var repoList []RepoItem

	switch provider {
	case:
		"GitLab"
		{
			client, err := gitlab.NewClient(token)
			if err != nil {
				return RepoPageModel{}, fmt.Errorf("failed to get GitLab client %e", err)
			}

			projects, _, err := client.Projects.ListProjects(&ListProjectsOptions{
				Membership: gitlab.Bool(true),
				Simple:     gitlab.Bool(true),
			})
			if err != nil {
				return RepoPageModel{}, fmt.Errorf("failed to get GitLab projects %e", err)
			}

			for _, r := range projects {
				repoList = append(repoList, RepoItem{
					Title:       r.Name,
					Description: r.Description,
				})
			}
		}
		return RepoPageModel{
			Cursor:             0,
			ListOfRepositories: repoList,
			Repository:         r,
		}, nil
	}
}

func (model RepoPageModel) Init() tea.Cmd { return nil }

func (model RepoPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	panic("Not implemented")
}

func (model RepoPageModel) View() string {
	panic("Not implemented")
}
