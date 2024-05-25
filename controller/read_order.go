package controller

import (
    "database/sql"
    "html/template"
    "net/http"
    "path/filepath"
    "strconv"
)

type StrukBelanja struct {
    Id          int
    Namamenu    string
    Harga       int
    Namauser    string
    Jumlah      string
    MPembayaran string
}

func StrukPDF(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

        r.ParseForm()
        nameUser := r.Form["nama-user"][0]
        jumlah := r.Form["jumlah"][0]
        MPembayaran := r.FormValue("metode-pembayaran")

		jumlasstr, err := strconv.Atoi(jumlah)
		if err != nil {
			panic(err)
		}

        Struk, err := db.Query("SELECT id, nama, harga FROM `order` ORDER BY id DESC")
        if err != nil {
            w.Write([]byte(err.Error()))
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        StrukModel := []StrukBelanja{}

        for Struk.Next() {
            var namabarang string
            var harga, id int

            err = Struk.Scan(&id, &namabarang, &harga)
            if err != nil {
                w.Write([]byte(err.Error()))
                w.WriteHeader(http.StatusInternalServerError)
                return
            }

            hargaTotal := harga * jumlasstr

            StrukList := StrukBelanja{
                Id:          id,
                Namamenu:    namabarang,
                Harga:       hargaTotal,
                Namauser:    nameUser,
                Jumlah:      jumlah,
                MPembayaran: MPembayaran,
            }

            StrukModel = append(StrukModel, StrukList)
        }

        fp := filepath.Join("view", "struk.html")
        tmpl, err := template.ParseFiles(fp)
        if err != nil {
            w.Write([]byte(err.Error()))
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        data := map[string]interface{}{
            "struk": StrukModel,
        }

        err = tmpl.Execute(w, data)

        if err != nil {
            w.Write([]byte(err.Error()))
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    }
}
