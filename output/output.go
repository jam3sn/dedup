package output

import (
	"dedup/dirscan"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func PrintStatus() {
	t := initTable("Scanning files... this may take a while.")
	renderTable(t)
}

func PrintOverview(path string, count int) {
	clearPreviousLines(2)
	t := initTable(fmt.Sprintf("Duplicates found: %d", count))
	t.AppendRow(table.Row{"Directory", path})
	renderTable(t)
}

func PrintDuplicates(duplicates dirscan.DuplicateFiles) {
	for name, _ := range duplicates {
		t := initTable(name)
		t.AppendHeader(table.Row{"Location", "Size", "Last Modified"})

		for _, file := range duplicates[name] {
			t.AppendRow(table.Row{file.Path, file.ReadableSize(), file.Modified})
		}

		renderTable(t)
	}
}

func initTable(title string) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBright)

	t.AppendHeader(table.Row{title})

	return t
}

func renderTable(t table.Writer) {
	t.Render()
	fmt.Println()
}

func clearPreviousLines(count int) {
	for i := 0; i < count; i++ {
		fmt.Print("\033[1A\033[K")
	}
}
