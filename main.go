package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Wordl!"));
}

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/", HomeHandler);

	http.ListenAndServe(":3000", router);
}