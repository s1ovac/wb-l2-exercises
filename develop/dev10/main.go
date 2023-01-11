/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	timeout time.Duration
	address string = ":8080"
)

func main() {
	initFlags()
	mux := http.NewServeMux()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	start(ctx, mux, address)
	go func(chan os.Signal) {
		osCall := <-c
		log.Printf("system call:%+v", osCall)
		cancel()
	}(c)
}

func initFlags() {
	flag.DurationVar(&timeout, "timeout", time.Second*5, "timeout for connecting to server")
	flag.Parse()
	if len(os.Args) >= 3 {
		address = ":" + os.Args[3]
	}
}

func start(ctx context.Context, mux *http.ServeMux, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Handler:      mux,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	go control(ctx, server)
	log.Fatal(server.Serve(listener))
}

func control(ctx context.Context, server *http.Server) {
	select {
	case <-ctx.Done():
		log.Printf("server stopped\n")
		server.Shutdown(ctx)
	default:
		fmt.Println("Hi")
	}
}
