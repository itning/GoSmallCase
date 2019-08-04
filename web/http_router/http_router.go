package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprint(w, "Welcome!\n")
	if err != nil {
		log.Print(err.Error())
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	if err != nil {
		log.Print(err.Error())
	}
}
