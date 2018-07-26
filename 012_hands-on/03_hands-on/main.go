package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type region struct {
	Region string
	Hotels []hotel
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := region{
		Region: "Northern",
		Hotels: []hotel{
			hotel{
				"Palace",
				"123 druy lane",
				"San Fran",
				"60598",
			},
			hotel{
				"motel on roids",
				"muhahahah",
				"the bay",
				"60484",
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
