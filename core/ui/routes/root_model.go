package routes

import (
	"fmt"

	"github.com/OliverKeefe/git-cleanse/core/auth"
	t "github.com/OliverKeefe/git-cleanse/core/types"
	"github.com/OliverKeefe/git-cleanse/core/ui/pages"
	tea "github.com/charmbracelet/bubbletea"
)

type RootModel struct {
	currentPage t.PageID
	pages       map[t.PageID]tea.Model
}

const (
	PageStartMenu  t.PageID = "StartMenu"
	PageAuthGitLab t.PageID = "GitLab"
	PageListGitLab t.PageID = "GitLabList"
)

func NewRootModel() RootModel {
	return RootModel{
		currentPage: PageStartMenu,
		pages: map[t.PageID]tea.Model{
			PageStartMenu:  pages.NewMenuModel(),
			PageAuthGitLab: pages.NewAuthPage("GitLab"),
			PageListGitLab: pages.NewAuthPage("GitLabList"),
		},
	}
}

func (rootModel RootModel) Init() tea.Cmd { return nil }

func (rootModel RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case t.NavigateMsg:
		rootModel.currentPage = msg.To
		return rootModel, nil

	case t.AuthSubmittedMsg:
		switch msg.Platform {
		case "GitLab":
			client, err := auth.GetGitLabClient(msg.Token)
			if err != nil {
				fmt.Println("Failed to create GitLab client:", err)
			} else {
				fmt.Println("GITLAB CLIENT:", client)
			}
		}
	}

	current := rootModel.pages[rootModel.currentPage]
	newModel, cmd := current.Update(msg)
	rootModel.pages[rootModel.currentPage] = newModel
	return rootModel, cmd
}

func (rootModel RootModel) View() string {
	return rootModel.pages[rootModel.currentPage].View()
}
