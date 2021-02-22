package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello docker")

}

func main() {
	http.HandleFunc("/", indexHandler)
	_ = http.ListenAndServe(":8080", nil)
}
