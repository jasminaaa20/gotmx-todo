package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world from HTMX with Go")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world from HTMX with Go ")
		io.WriteString(w, r.Method)
	})
	log.Fatal(http.ListenAndServe(":8000", nil)) //Fatal is equivalent to [Print] followed by a call to os.Exit(1)
}
