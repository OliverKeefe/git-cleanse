package types

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
	Header string
	Body   string
	Footer string
}
type RepoPath string
