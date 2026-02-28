package tree_sitter_acsl_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_acsl "github.com/luc65r/tree-sitter-acsl/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_acsl.Language())
	if language == nil {
		t.Errorf("Error loading ANSI/ISO C Specification Language grammar")
	}
}
