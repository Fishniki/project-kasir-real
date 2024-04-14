package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Order struct {
	Id       int
	IdBarang int
	Foto     string
	Nama     string
	Harga    int
}

type List struct {
	Id    int
	Name  string
	Harga int
	Foto  string
	Count int
}

func IndexView(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Select data menu
		menuSelect, err := db.Query("SELECT * FROM `menu` ORDER BY id DESC")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		upldMenu := List{}
		menuCard := []List{}
		for menuSelect.Next() {
			var id, harga int
			var nama, foto string

			err = menuSelect.Scan(&id, &foto, &nama, &harga)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			upldMenu.Id = id
			upldMenu.Foto = foto
			upldMenu.Name = nama
			upldMenu.Harga = harga

			menuCard = append(menuCard, upldMenu)
		}

		// Select data order
		rows, err := db.Query("SELECT * FROM `order` ORDER BY id DESC")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		upldOrder := Order{}
		orderList := []Order{}
		for rows.Next() {
			var id, idbarang, harga int
			var foto, nama string

			err = rows.Scan(&id, &idbarang, &foto, &nama, &harga)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			upldOrder.Id = id
			upldOrder.IdBarang = idbarang
			upldOrder.Foto = foto
			upldOrder.Nama = nama
			upldOrder.Harga = harga

			orderList = append(orderList, upldOrder)
		}

		// Render index.html template with both menu and order data
		fp := filepath.Join("view", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]interface{})
		data["upload"] = menuCard
		data["order"] = orderList

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
