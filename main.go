package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Movie struct {
	Title    string
	Director string
	Year     int
}

func main() {
	fmt.Println("Hello world from HTMX with Go")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmplt := template.Must(template.ParseFiles("index.html"))
		movies := map[string][]Movie{
			"movies": {
				{Title: "Avatar", Director: "James Cameron", Year: 2009},
				{Title: "The Avengers", Director: "Joss Whedon", Year: 2012},
				{Title: "The Wolf of Wall Street", Director: "Martin Scorsese", Year: 2013},
			},
		}
		tmplt.Execute(w, movies)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil)) //Fatal is equivalent to [Print] followed by a call to os.Exit(1)
}
