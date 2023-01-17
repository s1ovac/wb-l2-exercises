package server

import (
	"net/http"
	"time"
)

func RunServer(addr string) {
	server := http.Server{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Addr:         addr,
	}
	server.ListenAndServe()
	timeout := time.After(time.Second * 30)
	done := make(chan bool)
	go func(done chan<- bool) {
		select {
		case <-timeout:
			done <- true
			server.Close()
		}
	}(done)
}
