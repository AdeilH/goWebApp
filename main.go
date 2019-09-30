package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8000"
	}
	fmt.Println(port)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(port, nil)
}
