package main

import (
	"os"
	"testing"
)

func TestFooo(t *testing.T) {
	s, err := os.Stat("frontend/node_modules/.pnpm/@jridgewell+gen-mapping@0.3.3/node_modules/@jridgewell/set-array")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s.IsDir(), s.Name())
}

func xTestFooo(t *testing.T) {
	s := server{tmp: "tmp"}
	ss := []string{"dirs.go", "frontend/", "go.mod", "main.go", "main_test.go", "README.md", "server/", "server.go", "tmp/", "zip.go"}
	name, err := s.zipDirs(ss...)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(name)
}
