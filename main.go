package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	path := fmt.Sprintf("http://127.0.0.1:%s/%s", port, os.Getenv("HEALTHPATH"))
	fmt.Println("now hitting path for " + path)
	_, err := http.Get(path)
	if err != nil {
		os.Exit(1)
	}
}
