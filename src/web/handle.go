package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	b Data
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func Submit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("projectName")
	desc := r.FormValue("description")
	sdate := r.FormValue("startDate")
	edate := r.FormValue("endDate")

	db, err := b.Connect()
	if err != nil {
		return
	}
	b.Write(db, fmt.Sprintf("INSERT INTO project (%s, %s, %s, %s) VALUES (:name, :description, :sdata, :edate)",
		name,
		desc,
		sdate,
		edate,
	))
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, "", http.StatusSeeOther)
}
