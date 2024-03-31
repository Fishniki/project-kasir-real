package controller

import (
	"net/http"
)

func Text() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Wordcuy"))
	}
}