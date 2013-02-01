package main

import (
	"code.google.com/p/couch-go"
	"fmt"
	"github.com/bmizerany/pat"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello "+req.URL.Query().Get(":name")+"!\n")
}

func SingleFriends(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "antes de conectar ")
	db, err := couch.NewDatabase("localhost", "443", "mb_profiles")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "inserte json aqui "+db.DBURL())
}

func main() {
	m := pat.New()
	m.Get("/singleFriends", http.HandlerFunc(SingleFriends))
	m.Get("/:name", http.HandlerFunc(hello))

	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
