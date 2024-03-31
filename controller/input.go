package controller

import (
	"database/sql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func InsertToDb(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			r.ParseMultipartForm(10 << 20)

			formdata := r.MultipartForm

			nama := r.Form["name-item"][0]
			harga := r.Form["harga"][0]
			foto := formdata.File["foto"]


			for i := range foto {
				file, err := foto[i].Open()
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				defer file.Close()

				temp, err := ioutil.TempFile("file/", "upload-*.jpg")
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				defer temp.Close()

				fotoname := temp.Name()

				filebytes, err := ioutil.ReadAll(file)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				temp.Write(filebytes)

				_, err = db.Exec("INSERT INTO `menu` (foto, nama, harga) VALUES (?,?,?)", fotoname, nama, harga)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				log.Printf("Upload File Suksess\n")

			}

			w.Write([]byte("Sucses"))

		} else if r.Method == "GET" {

			fp := filepath.Join("view", "input.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		}
	}

}
