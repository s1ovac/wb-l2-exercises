Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: error
```
Обьяснение:
```
 	Ситуация аналогичная с интерфейсами.
	Поле data интерфейса error nil, а поле с метаинформацией не пустое, 
	следовательно нужно сравнивать поле data c nil, а не весь тип.

```
