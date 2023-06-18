package Server

import (
	s "GoMango/src/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

func HostHTML(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/", s.Index).Methods("GET")
	r.HandleFunc("/", s.Submit).Methods("POST")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))

	log.Println("HTML server listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func HostPHP(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("php", "list.php")
		output, err := cmd.Output()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(output)
	})

	log.Println("PHP server listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
