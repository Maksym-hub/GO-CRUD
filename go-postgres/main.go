package main


import (
"fmt"
"go-postgres/routers"
"log"
"net/http"
)

func main() {
	r := routers.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
