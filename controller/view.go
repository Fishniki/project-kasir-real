package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type List struct {
	Id    int
	Name  string
	Harga int
	Foto  string
	Count int
}

func IndexView(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT id, nama, harga, foto FROM `menu`")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		upld := List{}
		menu := []List{}
		for rows.Next() {

			var id, harga int
			var name, foto string

			err = rows.Scan(&id, &name, &harga, &foto)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			upld.Id = id
			upld.Name = name
			upld.Harga = harga
			upld.Foto = foto

			menu = append(menu, upld)

		}

		upld.Count = len(menu)
		if upld.Count > 0 {
			// tmpl.ExecuteTemplate(w, "view.html", res)
			fp := filepath.Join("view", "index.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["upload"] = menu

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}
		db.Close()

	}
}
