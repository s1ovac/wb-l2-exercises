/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"flag"
	"github.com/s1ovac/go-telnet/internal/client"
	"github.com/s1ovac/go-telnet/internal/server"
	"os"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные из сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	timeout time.Duration
	host    string = "localhost"
	port    string = ":8080"
)

func main() {
	initFlags()
	go server.RunServer(host + port)
	client := client.NewTelnet(timeout, host, port)
	client.Start()
}

func initFlags() {
	flag.DurationVar(&timeout, "timeout", time.Second*15, "timeout for connecting to server")
	flag.Parse()
	if len(os.Args) == 1 {
		host = flag.Arg(1)
	}
	if len(os.Args) == 2 {
		port = ":" + flag.Arg(2)
	}
}
