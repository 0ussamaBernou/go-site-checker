package main

import (
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("You")

	go http.Handle("/", templ.Handler(component))
	// go http.Handle("/foo", fooHandler)

	go http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Method))
	})

	port := "1323"
	fmt.Println("Listening on port: ", port)
	http.ListenAndServe(":"+port, nil)
}
func check() string {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}
