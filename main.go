package main

import (
	"database/sql"
	"log"
	"net/http"
	"project-kasir/config"
	"project-kasir/controller"
)

func Routes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.IndexView(db))
	server.HandleFunc("/create", controller.InsertToDb(db))
	server.HandleFunc("/create/order", controller.InputToOrder(db))
	server.HandleFunc("/create/order/delete", controller.DeletPesanan(db))
	server.HandleFunc("/create/order/refresh", controller.DeletOrder(db))
	server.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("file"))))
}

func main() {
	db, _ := config.ConectDB()
	log.Println("Server Berjalan di localhost:9000")

	server := http.NewServeMux()
	Routes(server, db)

	http.ListenAndServe(":9000", server)

}
