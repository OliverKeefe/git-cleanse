package pages

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v6"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type RepoItem struct {
	Title       string
	Description string
}

type GitLabPageModel struct {
	Cursor             int
	ListOfRepositories []RepoItem
	Repository         []*git.Repository
	Selected           string
	Width              int
	Height             int
}

func NewGitLabPageModel(token string, uri string, baseUrl string) (GitLabPageModel, error) {
	var projectList []RepoItem

	client, err := gitlab.NewClient(token)
	if err != nil {
		return GitLabPageModel{}, fmt.Errorf("unable to get gitlab client %e", err)
	}

	projects, _, err := client.Projects.ListProjects(&gitlab.ListProjectsOptions{
		Membership: gitlab.BoolValue(true),
		Simple:     gitlab.BoolValue(true),
	})
	if err != nil {
		return RepoPageModel{}, fmt.Errorf("failed to get gitlab projects %e", err)
	}

	for _, p := range projects {
		projectList = append(projectList, RepoItem{
			Title:       p.Name,
			Description: p.Description,
		})
	}
	return GitLabPageModel{
		Cursor:             0,
		ListOfRepositories: projectList,
	}, nil
}

func (model GitLabPageModel) Init() tea.Cmd {
	return nil
}

func (model GitLabPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	panic("Not implemented yet.")
}
