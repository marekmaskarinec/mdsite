package main

import (
	"io/ioutil"
	"strings"
)

/*Scans directory for markdown files*/
func ScanDir() []string {
	var toReturn []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic("looking for md files: " + err.Error())
	}
	
	for i := range files {
		if strings.Contains(files[i].Name(), ".md") {
			toReturn = append(toReturn, files[i].Name())		
		}
	}
	return toReturn
}

/*loads markdown files*/
func LoadMd(files []string) [][]byte {
	var toReturn [][]byte
	var dat []byte
	var err error

	for i := range files {
		dat, err = ioutil.ReadFile(path+files[i])
		if err != nil {
			panic("loading md files: " + err.Error())
		} else {
			toReturn = append(toReturn, dat)
		}
	}
	return toReturn
}

func LoadTemplate() string {
	dat, err := ioutil.ReadFile(path + "template.html")
	if err != nil {
		panic("loading template: " + err.Error())
	}
	return string(dat)
}

func SaveHtml(toSave, names []string) {
	var filename string
	var err error
	
	for i := range toSave {
		filename = strings.Split(names[i], ".")[0] + ".html"
		err = ioutil.WriteFile(filename, []byte(toSave[i]), 0644)
		if err != nil {
			panic("saving " + filename + ": " + err.Error())
		}
	}
}
