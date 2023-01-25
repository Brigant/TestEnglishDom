package main

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
import (
	"flag"
	"fmt"
	"os"
)

func main() {
	info := "Specify at least one argument. The value of argument should be a path. ex.: ./ or ./my_dir1 or ./my_dir1 /my_dir2"

	flag.Parse()

	sPath := flag.Args()

	// fmt.Printf("%#v", sPath)

	if len(sPath) == 0 {
		fmt.Println(info)
		return
	}

	result := make(map[string]string)

	for _, path := range sPath {
		sFiles, _ := os.ReadDir(path)
		for _, f := range sFiles {
			result[f.Name()] = f.Name()
		}
	}

	for path := range result {
		fmt.Println(path)
	}
}
