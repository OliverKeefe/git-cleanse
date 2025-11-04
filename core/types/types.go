package types

import "github.com/go-git/go-git/v6/plumbing/hash"

type PageID string
type NavigateMsg struct {
	To PageID
}
type AuthSubmittedMsg struct {
	Platform string
	Email    string
	Token    string
}

type Commit struct {
	Author       string
	Committer    string
	Message      string
	TreeHash     hash.Hash
	ParentHashes []hash.Hash
}
type RepoPath string

type BasePath string
