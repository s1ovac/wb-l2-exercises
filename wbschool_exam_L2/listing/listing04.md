Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Вывод: fatal error: all goroutines are asleep - deadlock!
	Это произошло из-за того, что мы не закрыли канал, и продолжили читать из пустого канала.
	Следовательно, для исправления ошибки нужно закрыть канал, после добавления значений в цикле, и цикл range автоматически завершит чтение.
```
Правильный код:
```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for n := range ch {
		println(n)
	}
}
```
