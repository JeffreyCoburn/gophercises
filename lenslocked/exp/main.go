package main

import (
	"html/template"
	"os"
	"time"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name    string
		Weekday string
		Age     int
		Flavors []string
		Prices  map[string]float64
	}{"John Smith", time.Now().Weekday().String(), 1, []string{"chocoloate", "vanilla"}, map[string]float64{"butter": 2.19, "milk": 2.59}}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
