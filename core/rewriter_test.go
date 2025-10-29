package core

import (
	"github.com/OliverKeefe/git-cleanse/core/types"
	"testing"
)

func TestRewriter_RewritePII(t *testing.T) {
	rewriter := Rewriter{
		[]string{`david.smith@aol.com`, `super-secret-keything23817s%$3`},
		[]string{`anon@noreply.github.com`, `XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`},
	}

	commit := types.Commit{
		Author:    "feat: add config and super-secret-keything23817s%$3",
		Committer: "this is a commit body message containing super-secret-keything23817s%$3 and other such things.",
		Message:   "Author: david.smith@aol.com",
	}

	var commits []types.Commit
	commits = append(commits, commit)

	results := rewriter.RewritePII(commits)

	expected := types.Commit{
		Author:    "feat: add config and XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		Committer: "this is a commit body message containing XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX and other such things.",
		Message:   "Author: anon@noreply.github.com",
	}

	if results[0].Committer != expected.Committer {
		t.Errorf("Expected commit Header to be: %s, got: %s", expected.Committer, results[0].Committer)
	}
	if results[0].Committer != expected.Committer {
		t.Errorf("Expected commit Body to be: %s, got: %s", expected.Committer, results[0].Committer)
	}
	if results[0].Message != expected.Message {
		t.Errorf("Expected commit Footer to be: %s, got: %s", expected.Message, results[0].Message)
	}
}

func TestRewriter_ReverseRewrite(t *testing.T) {
	rewriter := Rewriter{
		[]string{`david.smith@aol.com`, `super-secret-keything23817s%$3`},
		[]string{`anon@noreply.github.com`, `XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`},
	}

	commit := types.Commit{
		Author:    "feat: add config and super-secret-keything23817s%$3",
		Committer: "this is a commit body message containing super-secret-keything23817s%$3 and other such things.",
		Message:   "Author: david.smith@aol.com",
	}

	var commits []types.Commit
	commits = append(commits, commit)

	results := rewriter.RewritePII(commits)
	reversedResults := rewriter.ReverseRewrite(results)

	expected := types.Commit{
		Author:    "feat: add config and super-secret-keything23817s%$3",
		Committer: "this is a commit body message containing super-secret-keything23817s%$3 and other such things.",
		Message:   "Author: david.smith@aol.com",
	}

	if reversedResults[0].Author != expected.Author {
		t.Errorf("Expected commit Header to be: %s, got: %s", expected.Author, reversedResults[0].Author)
	}
	if reversedResults[0].Committer != expected.Committer {
		t.Errorf("Expected commit Body to be: %s, got: %s", expected.Committer, reversedResults[0].Committer)
	}
	if reversedResults[0].Message != expected.Message {
		t.Errorf("Expected commit Footer to be: %s, got: %s", expected.Message, reversedResults[0].Message)
	}
}

func TestGetRepo_GetRepo(t *testing.T) {
	var uri string = "https://github.com/OliverKeefe/algostruct.git"

	repo, err := GetRepo("", uri, false)
	if err != nil {
		t.Fatalf("expected repository to be non-nil struct %e", err)
	}

	if repo == nil {
		t.Fatalf("returned a nil repository %e", err)
	}

	ref, err := repo.Head()
	if err != nil {
		t.Fatalf("repository HEAD is nil %e", err)
	}

	t.Logf("Successfully cloned repo: HEAD = %s", ref.Hash())
}
