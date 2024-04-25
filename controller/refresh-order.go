package controller

import (
	"database/sql"
	"log"
	"net/http"
)

func DeletOrder(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := db.Exec("DELETE FROM `order`")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			panic(err)
		}
		log.Println("Pesanan berhasil di Reresh👌")
		http.Redirect(w,r, "/", http.StatusMovedPermanently)

	}
}