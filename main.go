package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/display", display)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t, _ := template.ParseFiles("home.html")
		t.Execute(w, nil)
		return
	}

	http.Redirect(w, r, "/display?info="+r.FormValue("info"), http.StatusSeeOther)
}

func display(w http.ResponseWriter, r *http.Request) {
	info := r.URL.Query().Get("info")

	if info == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	t, _ := template.ParseFiles("display.html")
	t.Execute(w, info)
}
