package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	user := User{Id: 1, Name: "哈哈", Pwd: "123"}
	bytes, e := json.Marshal(user)
	if e != nil {
		_, e := w.Write([]byte("出错"))
		if e != nil {
			log.Print(e.Error())
		}
		return
	}
	_, err := w.Write(bytes)
	//_, err := fmt.Fprint(w, "Welcome!\n")
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
