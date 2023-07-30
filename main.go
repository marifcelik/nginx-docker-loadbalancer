package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		serverName := os.Getenv("SERVER_NAME")
		fmt.Fprintf(w, "hi from %v", serverName)
	})

	http.ListenAndServe(":8080", nil)
}
