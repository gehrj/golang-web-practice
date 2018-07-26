package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	home := Page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}

// the big thing to take away here is that using html/template is build on top of text/template so you can do all the same things.
// the big advantage is that using html/template it will escape characters that could be dangerous for cross site scripting
// as you can see in this example the javascript code does not run the alert but it did in the text/template example
// essentially just use html/template so dont have to worry about cross site scripting.
