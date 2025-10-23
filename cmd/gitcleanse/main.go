package main

import (
	"fmt"
	"os"

	"github.com/OliverKeefe/git-cleanse/core/ui/routes"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func main() {
	//client, err := auth.GetGitLabClient("glpat-AMtIP5ZQi-mGhJKD7TSfEW86MQp1OjcwaAk.01.0z08wjz3y")
	//if err != nil {
	//	log.Fatalf("%s", err)
	//}
	//
	//version, _, err := client.Version.GetVersion()
	//
	//if err != nil {
	//	log.Printf("Warning: could not fetch GitLab version: %v", err)
	//} else {
	//	log.Printf("Connected to GitLab API version: %s", version.Version)
	//}

	//projects, err := repositories.ListGitLabProjects(client)
	//users, err := repositories.ListGitLabUsers(client)

	//for _, u := range users {
	//	log.Printf("%s", u)
	//}

	//for _, p := range projects {
	//	log.Printf("%s", p.NameWithNamespace)
	//}

	program := tea.NewProgram(routes.NewRootModel())
	if _, err := program.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
