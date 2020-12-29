# mdsite
Static website generator

# setup
1. Clone this repository.

2. Change into the directory: `cd mdsite`.

3. Install using `go install -i`. Note: you have to add $GOPATH/bin to your path. If you don't want to do that, you can use `go build` and move the binary to `/bin`.

# usage
Create folder `content` in your project folder. This folder holds all files needed to generate the website.

## template
Template is a html file. It is, what will be used to insert your md files into. You need to create line containing `###CONTENT###`. This line will be replaced with content.

## content files
After making template, you can start making content for your website in markdown.

## building website
To build the website run mdsite in parent directory of the content directory. Html files will now be generated. Names of the html files will be same as markdown files, so `about_me.md` will generate `about_me.html` etc.
If you want to add CSS to your site, it againg should be in parent of the content folder.
