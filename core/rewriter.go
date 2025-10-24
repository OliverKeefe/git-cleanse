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
}
