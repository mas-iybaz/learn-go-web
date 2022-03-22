package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

var counter = 5

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Letakkan proses utama pada goroutine, dan letakkan kode deteksi di luar
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("Request Canceled")
			} else {
				log.Println("Unknown error:", err.Error())
			}
		}
	case <-done:
		log.Println("Done")
		counter -= 1
	}

	log.Println("Request remaining:", counter)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}
