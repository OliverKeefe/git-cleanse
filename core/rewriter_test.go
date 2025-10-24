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
		Header: "feat: add config and super-secret-keything23817s%$3",
		Body:   "this is a commit body message containing super-secret-keything23817s%$3 and other such things.",
		Footer: "Author: david.smith@aol.com",
	}

	var commits []types.Commit
	commits = append(commits, commit)

	results := rewriter.RewritePII(commits)

	expected := types.Commit{
		Header: "feat: add config and XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		Body:   "this is a commit body message containing XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX and other such things.",
		Footer: "Author: anon@noreply.github.com",
	}

	if results[0].Header != expected.Header {
		t.Errorf("Expected commit Header to be: %s, got: %s", expected.Header, results[0].Header)
	}
	if results[0].Body != expected.Body {
		t.Errorf("Expected commit Body to be: %s, got: %s", expected.Body, results[0].Body)
	}
	if results[0].Footer != expected.Footer {
		t.Errorf("Expected commit Footer to be: %s, got: %s", expected.Footer, results[0].Footer)
	}
}

func TestRewriter_ReverseRewrite(t *testing.T) {
	rewriter := Rewriter{
		[]string{`david.smith@aol.com`, `super-secret-keything23817s%$3`},
		[]string{`anon@noreply.github.com`, `XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`},
	}

	commit := types.Commit{
		Header: "feat: add config and super-secret-keything23817s%$3",
		Body:   "this is a commit body message containing super-secret-keything23817s%$3 and other such things.",
		Footer: "Author: david.smith@aol.com",
	}

	var commits []types.Commit
	commits = append(commits, commit)

	results := rewriter.RewritePII(commits)
	reversedResults := rewriter.ReverseRewrite(results)

	expected := types.Commit{
		Header: "feat: add config and super-secret-keything23817s%$3",
		Body:   "this is a commit body message containing super-secret-keything23817s%$3 and other such things.",
		Footer: "Author: david.smith@aol.com",
	}

	if reversedResults[0].Header != expected.Header {
		t.Errorf("Expected commit Header to be: %s, got: %s", expected.Header, reversedResults[0].Header)
	}
	if reversedResults[0].Body != expected.Body {
		t.Errorf("Expected commit Body to be: %s, got: %s", expected.Body, reversedResults[0].Body)
	}
	if reversedResults[0].Footer != expected.Footer {
		t.Errorf("Expected commit Footer to be: %s, got: %s", expected.Footer, reversedResults[0].Footer)
	}
}
