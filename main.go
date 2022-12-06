package main

import (
	"dedup/dedup"
	"dedup/dirscan"
	"dedup/output"
	"fmt"
	"os"
)

func main() {
	basePath := getBasePath()
	duplicates, count := searchForDuplicates(basePath)

	output.PrintOverview(basePath, count)
	if count > 0 {
		output.PrintDuplicates(duplicates)
	}
}

func getBasePath() string {
	if len(os.Args) == 1 {
		fmt.Println("Argument for path required, e.g. dedup ./some-dir")
		os.Exit(1)
	}

	return os.Args[1]
}

func searchForDuplicates(path string) (dirscan.DuplicateFiles, int) {
	output.PrintStatus()

	files, err := dirscan.GetFiles(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return dedup.FilterDuplicateFiles(files)
}
