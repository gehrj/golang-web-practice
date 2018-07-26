package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type Item struct {
	Name        string
	Description string
	Price       float32
}

type Current struct {
	Meal  string
	Items []Item
}

type Menu []Current

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := Menu{
		Current{
			Meal: "Breakfast",
			Items: []Item{
				Item{
					Name:        "Eggs",
					Description: "Hard Boiled in fancy pan!",
					Price:       7.95,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
