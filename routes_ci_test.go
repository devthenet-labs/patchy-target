//go:build ci

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// TestRoutesDocumented scans all production route registrations, including
// those in main(), which the test process never calls. A new handler must be
// added to ROUTES.md for the CI-only test to pass.
func TestRoutesDocumented(t *testing.T) {
	registered := registeredRoutes(t)
	documented := documentedRoutes(t)
	for route := range registered {
		if !documented[route] {
			t.Errorf("registered route %q is missing from ROUTES.md", route)
		}
	}
	for route := range documented {
		if !registered[route] {
			t.Errorf("ROUTES.md lists unregistered route %q", route)
		}
	}
}

func registeredRoutes(t *testing.T) map[string]bool {
	t.Helper()
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	routes := make(map[string]bool)
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(token.NewFileSet(), name, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		ast.Inspect(file, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			selector, ok := call.Fun.(*ast.SelectorExpr)
			if !ok || (selector.Sel.Name != "Handle" && selector.Sel.Name != "HandleFunc") {
				return true
			}
			if len(call.Args) == 0 {
				t.Errorf("%s: route registration has no pattern", name)
				return true
			}
			literal, ok := call.Args[0].(*ast.BasicLit)
			if !ok || literal.Kind != token.STRING {
				t.Errorf("%s: route registration must use a literal pattern", name)
				return true
			}
			pattern, err := strconv.Unquote(literal.Value)
			if err != nil {
				t.Errorf("%s: invalid route pattern: %v", name, err)
				return true
			}
			routes[pattern] = true
			return true
		})
	}
	if len(routes) == 0 {
		t.Fatal("no production routes found")
	}
	return routes
}

func documentedRoutes(t *testing.T) map[string]bool {
	t.Helper()
	data, err := os.ReadFile(filepath.Clean("ROUTES.md"))
	if err != nil {
		t.Fatal(err)
	}
	routes := make(map[string]bool)
	for _, line := range strings.Split(string(data), "\n") {
		if !strings.HasPrefix(line, "- `") {
			continue
		}
		pattern := strings.TrimSuffix(strings.TrimPrefix(line, "- `"), "`")
		if pattern == "" || !strings.HasSuffix(line, "`") {
			t.Errorf("malformed route entry %q", line)
			continue
		}
		if routes[pattern] {
			t.Errorf("duplicate route %q in ROUTES.md", pattern)
		}
		routes[pattern] = true
	}
	if len(routes) == 0 {
		t.Fatal("ROUTES.md has no routes")
	}
	return routes
}

func TestRouteInventoryIsSorted(t *testing.T) {
	data, err := os.ReadFile("ROUTES.md")
	if err != nil {
		t.Fatal(err)
	}
	var listed []string
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "- `") {
			listed = append(listed, line)
		}
	}
	if !sort.StringsAreSorted(listed) {
		t.Fatal("ROUTES.md routes must be sorted")
	}
}
