package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hang(w http.ResponseWriter, r *http.Request) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello astaxie!")
	}
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", hang)
	timeout := 5 * time.Second
	s := &http.Server{
		ReadTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
		Addr:              ":8080",
		Handler:           m,
	}
	log.Fatal(s.ListenAndServe())
}
