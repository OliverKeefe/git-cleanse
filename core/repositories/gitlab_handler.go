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

// ListGitLabUsers TODO: Change this because currently it gets every single user in the tenant.
func ListGitLabUsers(client *gitlab.Client) ([]*gitlab.User, error) {
	options := &gitlab.ListUsersOptions{
		ListOptions: gitlab.ListOptions{
			Pagination: "",
			PerPage:    10,
			Page:       1,
			PageToken:  "",
			OrderBy:    "",
			Sort:       "",
		},
		Active:               nil,
		Blocked:              nil,
		Humans:               nil,
		ExcludeInternal:      nil,
		ExcludeActive:        nil,
		ExcludeExternal:      nil,
		ExcludeHumans:        nil,
		PublicEmail:          nil,
		Search:               nil,
		Username:             nil,
		ExternalUID:          nil,
		Provider:             nil,
		CreatedBefore:        nil,
		CreatedAfter:         nil,
		OrderBy:              nil,
		Sort:                 nil,
		TwoFactor:            nil,
		Admins:               nil,
		External:             nil,
		WithoutProjects:      nil,
		WithCustomAttributes: nil,
		WithoutProjectBots:   nil,
	}
	var allUsers []*gitlab.User
	for {
		users, response, err := client.Users.ListUsers(options)
		if err != nil {
			return nil, err
		}

		allUsers = append(allUsers, users...)
		if response.CurrentPage >= response.TotalPages {
			break
		}
		options.Page = response.NextPage
	}
	log.Printf("Fetched %d Users from GitLab", len(allUsers))
	return allUsers, nil
}
