package core

import (
	"fmt"

	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

type ScanFor []string

// Open repo func

//// Search repo func - call rewrite
//
//func (scan ScanFor) ScanRepo(repoPath string) (bool, error) {
//	repo, err := git.PlainOpen(repoPath)
//	if err != nil {
//		return fmt.Errorf("failed to open repo: %w", err)
//	}
//
//	worktree, err := repo.Worktree()
//	if err != nil {
//		return err
//	}
//}
