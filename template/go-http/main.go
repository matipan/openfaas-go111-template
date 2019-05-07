package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"handler/function"
)

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8082),
		ReadTimeout:    getDuration("write_timeout", 3*time.Second),
		WriteTimeout:   getDuration("read_timeout", 3*time.Second),
		MaxHeaderBytes: 1 << 20, // Max header of 1MB
	}

	http.HandleFunc("/", function.Handle)
	log.Fatal(s.ListenAndServe())
}

func getDuration(key string, defaultValue time.Duration) time.Duration {
	result := defaultValue
	if val := os.Getenv(key); val != "" {
		parsed, err := time.ParseDuration(val)
		if err != nil {
			return result
		}
		result = parsed
	}
	return result
}