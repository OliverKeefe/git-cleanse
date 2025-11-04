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

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

// NewGitLabPageModel constructs a new Gitlab Page Model.
// param: token string
func NewGitLabPageModel(token string, uri string, baseUrl string) (GitLabPageModel, error) {
	var projectList []RepoItem

	client, err := gitlab.NewClient(token)
	if err != nil {
		return GitLabPageModel{}, fmt.Errorf("unable to get gitlab client %e", err)
	}

	projects, _, err := client.Projects.ListProjects(&gitlab.ListProjectsOptions{
		Membership: GitLabBool(true),
		Simple:     GitLabBool(true),
	})
	if err != nil {
		return GitLabPageModel{}, fmt.Errorf("failed to get gitlab projects %e", err)
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

func (model GitLabPageModel) View() string {
	return appStyle.Render(model.list.View())
}

func GitLabBool(v bool) *bool {
	return &v
}
