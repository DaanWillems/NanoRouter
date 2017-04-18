package main

import (
	"NanoRouter/nano"
	"fmt"
	"net/http"
)

func main() {
	router := nano.NewRouter()
	router.NewRoute("/home", home)
	router.SetNotFoundRoute(notfound)
	http.Handle("/", router)
	http.ListenAndServe(":8380", nil)
}

func notfound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unknown page")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, nano.Vars["a"])
}
