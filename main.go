package main

import (
	"fmt"
	"go/server/api"
	"go/server/data"
	"html/template"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	err := http.ListenAndServe(":3333", server)
	if err == nil {
		fmt.Println("Error while opening the server")
	}
}

func handleTemplate (w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte("Internal Server Error")))
	}
	html.Execute(w, data.GetAll())
}