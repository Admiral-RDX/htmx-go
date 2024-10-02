package main

import (
	"fmt"
	base "go-templ-dev/templ/base"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := base.Hello("John")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
