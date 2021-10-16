package main

import (
	_ "go/ast"
	"html/template"
	"log"
	"net/http"
)
func main()  {
	http.HandleFunc("/", mainPage)

	port := ":9090"
	println("server listent port", port)
	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Fatal("listenand server", err)
	}
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err!= nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err:= tmpl.Execute(w, nil); err !=nil {
		http.Error(w, err.Error(), 400)
		return
	}
}