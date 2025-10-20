package repositories

import (
	"log"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

func ListGitLabProjects(client *gitlab.Client) ([]*gitlab.Project, error) {
	// Pagination
	options := &gitlab.ListProjectsOptions{
		Membership: gitlab.Ptr(true),
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 50,
		},
	}

	var allProjects []*gitlab.Project

	for {
		projects, response, err := client.Projects.ListProjects(options)
		if err != nil {
			return nil, err
		}

		allProjects = append(allProjects, projects...)

		if response.CurrentPage >= response.TotalPages {
			break
		}

		options.Page = response.NextPage
	}

	log.Printf("Fetched %d projects from GitLab", len(allProjects))
	return allProjects, nil
}
