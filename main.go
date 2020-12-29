package main

func run() {
	content := ConvertToHtml(LoadMd(ScanDir("content")))
	template := LoadTemplate()
	SaveHtml(InsertToTemplate(template, content), ScanDir("content"))
}

func main() {
	run()
}
