package main

import (
	"fmt"
)

func run() {
	fmt.Println("Scanning for markdown files")
	files := ScanDir("content")
	fmt.Println("Loading markdown files")
	mds := LoadMd(files)
	fmt.Println("Converting markdown to html")
	content := ConvertToHtml(mds)
	fmt.Println("Loading template")
	template := LoadTemplate()
	fmt.Println("Inserting into template")
	final := InsertToTemplate(template, content)
	fmt.Println("Saving html files")
	SaveHtml(final, ScanDir("content"))
}

func main() {
	run()
}
