package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/bye", byeHandler)
	http.ListenAndServe(":7777", nil)
}

func rootHandler(rw http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "data read error", http.StatusBadRequest)
	}
	// rw.Write([]byte("Hello"))
	rw.Write([]byte(fmt.Sprintf("Hello %s", data)))
}

func byeHandler(rw http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "data read error", http.StatusBadRequest)
	}
	rw.Write([]byte(fmt.Sprintf("Bye %s", data)))
}
