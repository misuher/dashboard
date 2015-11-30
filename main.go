package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	port := ":5000"
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/login", login)
	fmt.Println("Servidor de goTwitter corriendo en localhost: %x", port)
	http.ListenAndServe(port, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["usermail"])
		fmt.Println("password:", r.Form["passwd"])
		t, _ := template.ParseFiles("static/index.html")
		t.Execute(w, nil)
	}
}
