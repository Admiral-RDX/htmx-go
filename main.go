package main

import (
	"context"
	"os"
    web "go-templ-dev/templ"
)

func main() {
	component := web.Hello("John")
	component.Render(context.Background(), os.Stdout)
}
