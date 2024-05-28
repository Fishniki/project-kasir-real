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

type Pemesan struct {
    Nama            string
    MetodePembayaran string
}

func StrukPDF(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        nameUser := r.FormValue("nama-user")
        jumlah := r.FormValue("jumlah")
        MPembayaran := r.FormValue("metode-pembayaran")

        jumlasstr, err := strconv.Atoi(jumlah)
        if err != nil {
            http.Error(w, "Invalid jumlah", http.StatusBadRequest)
            return
        }

        pemesanModel := []Pemesan{
            {
                Nama:            nameUser,
                MetodePembayaran: MPembayaran,
            },
        }

        Struk, err := db.Query("SELECT id, nama, harga FROM `order` ORDER BY id DESC")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer Struk.Close()

        StrukModel := []StrukBelanja{}
        totalHarga := 0

        for Struk.Next() {
            var id, harga int
            var namabarang string

            err = Struk.Scan(&id, &namabarang, &harga)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            hargaTotal := harga * jumlasstr
            totalHarga += hargaTotal

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

        if err = Struk.Err(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fp := filepath.Join("view", "struk.html")
        tmpl, err := template.ParseFiles(fp)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        data := map[string]interface{}{
            "struk":      StrukModel,
            "user":       pemesanModel,
            "totalHarga": totalHarga,
        }

        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}
