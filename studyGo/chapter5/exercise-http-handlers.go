package main

import (
	"fmt"
	"log"
	"net/http"
)

type Str string

type Hello struct {
	Greeting string
	Punct string
	Who string
}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){
	fmt.Fprint(w, h.Greeting, h.Punct, h.Who)
}

func (s Str) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){
	fmt.Fprint(w, s)
}

func main() {
	http.Handle("/str", Str("I am a frayed knot."))	
	http.Handle("/hello", &Hello{"Hello",":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000",nil))
}
