package controller

import (
	"database/sql"
	// "html/template"
	"net/http"
	// "path/filepath"
)

func InputToOrder(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

				id := r.FormValue("pesan")
				idBarang := r.FormValue("id_" + id) 
				foto := r.FormValue("foto-menu_" + id)
				nama := r.FormValue("name-menu_" + id)
				harga := r.FormValue("harga-menu_" + id)
				_, err := db.Exec("INSERT INTO `order` (id_barang, foto, nama, harga) VALUES (?,?,?,?)", idBarang, foto, nama, harga)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					panic(err)
				}
				http.Redirect(w, r, "/", http.StatusMovedPermanently)        
		}
		
	}
}