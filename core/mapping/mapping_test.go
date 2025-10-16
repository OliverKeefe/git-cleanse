package core_test

import (
	"testing"

	core "github.com/OliverKeefe/git-cleanse/core/mapping"
)

func TestMappingListApply(t *testing.T) {
	mappings := []core.Mapping[string]{
		{Old: "alice.smith@aol.com", New: "anon.no-reply@github.com"},
		{Old: "steveo@bigcorp.co.uk", New: "anon@bigcorp.noreply.co.uk"},
	}

	mappingList := core.NewMappingList(mappings)

	tests := []struct {
		input    string
		expected string
	}{
		{"alice.smith@aol.com", "anon.no-reply@github.com"},
		{"steveo@bigcorp.co.uk", "anon@bigcorp.noreply.co.uk"},
	}

	for _, test := range tests {
		result := mappingList.Apply(test.input)
		if result != test.expected {
			t.Errorf("Apply(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
