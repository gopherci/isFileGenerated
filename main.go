package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shurcooL/go/generated"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%v <root directory> <path to file>\n", os.Args[0])
		os.Exit(2)
	}

	generated, err := generated.ParseFile(filepath.Join(os.Args[1], os.Args[2]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error checking if %q has generated comment: %v\n", os.Args[2], err)
		os.Exit(3)
	}

	if strings.HasPrefix(os.Args[2], "vendor/") || strings.Contains(os.Args[2], "/vendor/") ||
		strings.HasPrefix(os.Args[2], "Godeps/") ||
		strings.HasPrefix(os.Args[2], "testdata/") || strings.Contains(os.Args[2], "/testdata/") {
		generated = true
	}

	if generated {
		fmt.Printf("%v was generated\n", os.Args[2])
		os.Exit(0)
	}
	fmt.Printf("%v was not generated\n", os.Args[2])
	os.Exit(1)
}
