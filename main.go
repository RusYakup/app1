package main

import (
	_ "go/ast"
	"html/template"
	"log"
	"net/http"
	"os"
)


func main()  {
	http.HandleFunc("/", HelloWorldHandler)

	port := ":9090"
	println("server listen port", port)
	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Fatal("App failed with error:", err)
	}
}
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	filepath := path + "/index.html"

	tmpl, err := template.ParseFiles(filepath)
	if err!= nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err:= tmpl.Execute(w, nil); err !=nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
