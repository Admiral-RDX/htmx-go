package main

import (
	"fmt"
	base "go-templ-dev/templ/base"
	calendar "go-templ-dev/templ/calendar"
	dashboard "go-templ-dev/templ/dashboard"
	login "go-templ-dev/templ/login"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	// Template definition
	loginPage := base.HTML(login.Login().Render)
	dashboardPage := base.HTML(dashboard.Dashboard().Render)
	calendarPage := base.HTML(calendar.Calendar().Render)

	// Static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.Handle("/", templ.Handler(loginPage))
	http.Handle("/dashboard", templ.Handler(dashboardPage))
	http.Handle("/calendar", templ.Handler(calendarPage))

	// HTTP
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
