package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"project-kasir/config"
	"project-kasir/controller"
)

func Routes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.IndexView(db))
	server.HandleFunc("/create", controller.InsertToDb(db))
	server.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("file"))))
}

func main() {
	db, _ := config.ConectDB()
	fmt.Println("conect DB")

	server := http.NewServeMux()
	Routes(server, db)

	http.ListenAndServe(":9000", server)

}
