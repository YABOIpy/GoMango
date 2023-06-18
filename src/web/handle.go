package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

	SqlProject := Project{
		Name:        r.FormValue("projectName"),
		Description: r.FormValue("description"),
		StartDate:   r.FormValue("startDate"),
		EndDate:     r.FormValue("endDate"),
	}
	fmt.Println("Received Data: ", SqlProject)
	db, err := b.Connect()
	if err != nil {
		return
	}

	stmt, err := db.Prepare(`
		INSERT INTO project (name, description, sdate, edate)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		SqlProject.Name,
		SqlProject.Description,
		SqlProject.StartDate,
		SqlProject.EndDate,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Added Project To Database")

	http.Redirect(w, r, "", http.StatusSeeOther)
}
