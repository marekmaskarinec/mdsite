package main

import (
	"strings"
	"github.com/gomarkdown/markdown"
	"fmt"
)

func FindContentPlace(template string) int {
	templateLines := strings.Split(template, "\n")
	for i := range templateLines {
		if strings.Contains(templateLines[i], "##CONTENT##") {
			return i
		}
	}
	return -1
}

func ConvertToHtml(mds [][]byte) []string {
	var toReturn []string

	for i := range mds {
		toReturn = append(toReturn, string(markdown.ToHTML(mds[i], nil, nil)))
	}
	return toReturn
}

func InsertToTemplate(template string, content []string) []string {
	var toReturn, contentLines, tmp []string
	cLine := FindContentPlace(template)
	templateLines := strings.Split(template, "\n")

	for i := range content {
		fmt.Println("content")
		contentLines = strings.Split(content[i], "\n")
		tmp = append(templateLines[0:cLine], contentLines...)
		tmp = append(tmp, templateLines[cLine+1:len(templateLines)]...)
		toReturn = append(toReturn, "")
		for j := range tmp {
			toReturn[i] += tmp[j] + "\n"
		}
	}
	return toReturn
}
