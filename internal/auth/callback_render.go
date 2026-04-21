package auth

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates/*.html templates/*.svg
var callbackFS embed.FS

var callbackTemplates = template.Must(template.ParseFS(callbackFS, "templates/*.html", "templates/*.svg"))

func renderSuccess(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := callbackTemplates.ExecuteTemplate(w, "success.html", nil); err != nil {
		fmt.Fprintf(w, "Signed in. You can close this window and return to your terminal.")
	}
}

func renderError(w http.ResponseWriter, title, detail string, link *errorLink) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := struct {
		Title  string
		Detail string
		Link   *errorLink
	}{Title: title, Detail: detail, Link: link}
	if err := callbackTemplates.ExecuteTemplate(w, "error.html", data); err != nil {
		fmt.Fprintf(w, "%s. %s", title, detail)
	}
}

type errorLink struct {
	URL  string
	Text string
}
