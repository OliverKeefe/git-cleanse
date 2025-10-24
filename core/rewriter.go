package core

import (
	"regexp"

	"github.com/OliverKeefe/git-cleanse/core/types"
)

type ToRewrite []string
type RewriteWith []string

type Rewriter struct {
	rewrite     ToRewrite
	rewriteWith RewriteWith
}

func NewRewriter(toRewrite []string, rewriteWith []string) Rewriter {
	return Rewriter{
		rewrite:     toRewrite,
		rewriteWith: rewriteWith,
	}
}

func (rewriter Rewriter) RewritePII(commits []types.Commit) []types.Commit {
	return rewriter.RewriteHelper(commits, rewriter.rewriteWith, rewriter.rewrite)
}

func (rewriter Rewriter) ReverseRewrite(commits []types.Commit) []types.Commit {
	return rewriter.RewriteHelper(commits, rewriter.rewrite, rewriter.rewriteWith)
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
