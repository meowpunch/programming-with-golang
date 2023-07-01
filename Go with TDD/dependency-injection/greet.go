package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {

	Greet(os.Stdout, "jinhoon")

	http.HandleFunc("/", GreetHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
