package controller

import (
	"database/sql"
	"log"
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

			// Periksa apakah idBarang sudah ada dalam database
			var count int
			row := db.QueryRow("SELECT COUNT(*) FROM `order` WHERE id_barang = ?", idBarang)
			err := row.Scan(&count)	
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("Error while checking duplicate id:", err)
				return
			}

			// Jika count > 0, idBarang sudah 	ada dalam database
			if count > 0 {
				log.Println("Data dengan idBarang yang sama sudah ada dalam database")
				// Tambahkan tanggapan sesuai kebutuhan, misalnya:
				http.Redirect(w,r,"/",http.StatusMovedPermanently)
				return
				
			}

			// Jika idBarang belum ada dalam database, masukkan data ke dalam database
			_, err = db.Exec("INSERT INTO `order` (id_barang, foto, nama, harga) VALUES (?,?,?,?)", idBarang, foto, nama, harga)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("Error while inserting data into database:", err)
				return
			}

			log.Println("Menu berhasil masuk ke pesanan")
			http.Redirect(w, r, "/", http.StatusMovedPermanently)        
		}
	}
}
