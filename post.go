package main

import (
	"bytes"
	"html/template"
)

// Post is a generic post from different services
type Post struct {
	Title string
	URL   string
}

type PostChannel chan []Post

func (post Post) toHtml() string {
	tmpl, err := template.New("post").Parse("<div><a href=\"{{ .URL}}\">{{ .Title}}</a></div>")
	if err != nil {
		panic(err)
	}

	var str bytes.Buffer
	err = tmpl.Execute(&str, post)
	if err != nil {
		panic(err)
	}

	return str.String()
}
