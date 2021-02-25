package main

import (
	"fmt"
	"os"
)

var path string

func run() {
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path = "content/"
	}

	fmt.Println("Scanning for markdown files")
	files := ScanDir()
	fmt.Println("Loading markdown files")
	mds := LoadMd(files)
	fmt.Println("Converting markdown to html")
	content := ConvertToHtml(mds)
	fmt.Println("Loading template")
	template := LoadTemplate()
	fmt.Println("Inserting into template")
	final := InsertToTemplate(template, content)
	fmt.Println("Saving html files")
	SaveHtml(final, ScanDir())
}

func main() {
	run()
}
