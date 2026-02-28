/**
 * @file ANSI/ISO C Specification Language grammar for tree-sitter
 * @author Lucas Ransan <lucas@ransan.fr>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "acsl",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
