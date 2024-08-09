package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world from HTMX with Go")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplt := template.Must(template.ParseFiles("index.html"))
		tmplt.Execute(w, nil)
	})
	log.Fatal(http.ListenAndServe(":8000", nil)) //Fatal is equivalent to [Print] followed by a call to os.Exit(1)
}
