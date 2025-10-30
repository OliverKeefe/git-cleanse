package pages

import (
	"github.com/OliverKeefe/git-cleanse/core"
	tea "github.com/charmbracelet/bubbletea"
)

type RepoItem struct {
	title       string
	description string
}

type RepoPageModel struct {
	cursor       int
	repositories []RepoItem
	selected     string
	width        int
	height       int
}

func NewRepoPageModel() RepoPageModel {
	return RepoPageModel{
		cursor:       0,
		repositories: core.GetRepo(),
	}
}

func (model RepoPageModel) Init() tea.Cmd { return nil }

func (model RepoPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

}

func (model RepoPageModel) View() string {

}
