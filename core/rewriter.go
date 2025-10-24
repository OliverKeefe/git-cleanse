package core

type ToReplace []string
type DontReplace []string

type Rewriter struct {
	PII    ToReplace
	notPII DontReplace
}

func NewRewriter(replace []string, dontReplace []string) Rewriter {
	return Rewriter{
		PII:    replace,
		notPII: dontReplace,
	}
func (rewriter Rewriter) RewriteHelper(commits []types.Commit, to []string, from []string) []types.Commit {
	for i, commit := range commits {
		for j, pattern := range from {
			re := regexp.MustCompile(regexp.QuoteMeta(pattern))

			if j < len(to) {
				commit.Header = re.ReplaceAllString(commit.Header, to[j])
				commit.Body = re.ReplaceAllString(commit.Body, to[j])
				commit.Footer = re.ReplaceAllString(commit.Footer, to[j])
			}
		}
		commits[i] = commit
	}
	return commits
}
