package core

import (
	"fmt"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
	"io"
	"log"
	"os"
	"regexp"
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

func (rewriter Rewriter) RewritePII(commits []*object.Commit) []*object.Commit {
	return rewriter.RewriteHelper(commits, rewriter.rewriteWith, rewriter.rewrite)
}

func (rewriter Rewriter) ReverseRewrite(commits []*object.Commit) []*object.Commit {
	return rewriter.RewriteHelper(commits, rewriter.rewrite, rewriter.rewriteWith)
}

func (rewriter Rewriter) RewriteHelper(commits []*object.Commit, to []string, from []string) []*object.Commit {
	for i, commit := range commits {
		for j, pattern := range from {
			re := regexp.MustCompile(regexp.QuoteMeta(pattern))

			if j < len(to) {
				commit.Author.Email = re.ReplaceAllString(commit.Author.Email, to[j])
				commit.Author.Name = re.ReplaceAllString(commit.Author.Name, to[j])
				commit.Committer.Email = re.ReplaceAllString(commit.Committer.Email, to[j])
				commit.Committer.Name = re.ReplaceAllString(commit.Committer.Name, to[j])
				commit.Message = re.ReplaceAllString(commit.Message, to[j])
			}
		}
		commits[i] = commit
	}
	return commits
}

// TODO: func (rewriter Rewriter) ScanFor

func GetRepo(path string, uri string, isLocal bool) (*git.Repository, error) {
	if !isLocal {
		repo, err := git.PlainOpen(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open repo: %w", err)
		}
		return repo, nil
	} else {
		dir, err := os.MkdirTemp("", "repo-*")
		if err != nil {
			return nil, fmt.Errorf("failed to create temp dir to clone repo to: %w", err)
		}

		repo, err := git.PlainClone(dir, &git.CloneOptions{
			URL:      uri,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatal(err)
		}
		return repo, nil
	}
}

// TODO: Get commits stored in slice in memory for iteration.
func GetRepoCommits(repo *git.Repository) ([]*object.Commit, error) {
	var commits []*object.Commit

	iter, err := repo.CommitObjects()
	if err != nil {
		return nil, fmt.Errorf("failed to get commit iterator: %w", err)
	}
	defer iter.Close()

	for {
		commit, err := iter.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read commit object: %w", err)
		} else {
			commits = append(commits, commit)
		}
	}
	return commits, nil
}
