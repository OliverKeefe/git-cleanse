package pages

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-git/go-git/v6"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type RepoItem struct {
	Title       string
	Description string
}

type GitLabPageModel struct {
	Cursor             int
	Client             *gitlab.Client
	list               list.Model
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
func (m GitLabPageModel) NewGitLabModel() (tea.Model, tea.Cmd) {
	return GitLabPageModel{}, nil
}

//func GetGitLabRepos() ([]RepoItem, error) {
//	var projectList []RepoItem
//
//	client := &gitlab.Client
//
//	projects, err := repos.ListGitLabProjects(client)
//
//	for _, p := range projects {
//		projectList = append(projectList, RepoItem{
//			Title:       p.Name,
//			Description: p.Description,
//		})
//	}
//	return GitLabPageModel{
//		Cursor:             0,
//		ListOfRepositories: projectList,
//	}, nil
//}

func (model GitLabPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmds []tea.Cmd

	return model, nil
}

func (model GitLabPageModel) View() string {
	return appStyle.Render(model.list.View())
}

func (model GitLabPageModel) Init() tea.Cmd {
	return nil
}

func GitLabBool(v bool) *bool {
	return &v
}
