package controller

import (
	"database/sql"
	"log"
	"net/http"
)

func DeletPesanan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		emp := r.URL.Query().Get("id")
		_, err := db.Exec("DELETE FROM `order` WHERE id=?", emp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println("Delete Sucsess")
		http.Redirect(w,r, "/", http.StatusMovedPermanently)

	}
}