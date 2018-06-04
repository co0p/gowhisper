package gowhisper

import (
	"net/http"
	"text/template"
)

type StatusPage struct {
	template *template.Template
	clients  *[]Client
}

type StatusEntry struct {
	Label  string
	URL    string
	Online bool
}

func NewStatusPage(clients *[]Client) (StatusPage, error) {
	t, err := template.ParseFiles("index.html")
	return StatusPage{template: t, clients: clients}, err
}

func (s *StatusPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	entries := []StatusEntry{}
	for _, v := range *s.clients {
		e := StatusEntry{Label: v.Label, URL: v.URL, Online: v.Online}
		entries = append(entries, e)
	}

	vm := struct {
		Entries []StatusEntry
	}{
		Entries: entries,
	}

	s.template.Execute(w, vm)
}
