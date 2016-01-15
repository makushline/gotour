package main

import (
    "log"
    "net/http"
    "io"
    "fmt"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, h.Greeting, h.Punct, h.Who )
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, string(s))
}

func main() {
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    log.Fatal(http.ListenAndServe("localhost:4001", nil))
}
