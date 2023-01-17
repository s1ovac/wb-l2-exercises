package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)
*/

// функция которая будет объединять один или более done каналов в single канал
func or(channels ...<-chan interface{}) <-chan interface{} {
	all := make(chan interface{})      //общий канал для всех горутин сигнализирующий что один из каналов закрылся
	for i, channel := range channels { //слушаем все каналы
		go func(ch <-chan interface{}, closer chan interface{}) {
			select {
			case <-ch: //если один из каналов закрылся сигнализируем остальным
				fmt.Printf("Один из каналов - %d закрылся!\n", i)
				close(all)
			case <-closer:

			}
		}(channel, all)
	}
	return all
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		time.Sleep(after)
		defer close(c)
	}()
	return c
}

func main() {
	start := time.Now()

	<-or(
		sig(2*time.Second),
		sig(3*time.Second),
		sig(1*time.Second),
	)
	fmt.Printf("fone after %v", time.Since(start))
}
