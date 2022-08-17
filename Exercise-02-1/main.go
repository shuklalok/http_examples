package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shuklalok/http_examples/Exercise-02-1/byehandler"
	"github.com/shuklalok/http_examples/Exercise-02-1/hellohandler"
)

func main() {
	l := log.New(os.Stdout, "greeting-code", log.LstdFlags)
	hh := hellohandler.NewHello(l)
	bh := byehandler.NewBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", bh)
	http.ListenAndServe(":7777", sm)
}
