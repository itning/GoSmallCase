package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/b", index)
	mux.HandleFunc("/a", a)
	e := http.ListenAndServe(":8080", mux)
	if e != nil {
		fmt.Println(e.Error())
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, e := writer.Write([]byte("哈哈"))
	if e != nil {
		fmt.Println(e.Error())
	}
}

func a(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, e := writer.Write([]byte("a"))
	if e != nil {
		fmt.Println(e.Error())
	}
}
