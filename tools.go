//go:build tools
// +build tools

// Package tools simply import graphql dependency
package tools

import (
	// this import is only a blank import that manages the graphql dependency
	_ "github.com/99designs/gqlgen"
)
