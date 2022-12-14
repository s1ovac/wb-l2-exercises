package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	result, err := decodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func unpackString(s string) (string, error) {
	var (
		sb   strings.Builder
		temp string
	)
	for _, v := range s {
		if n, err := strconv.ParseInt(string(v), 0, 64); err == nil {
			for i := int64(0); i < n-1; i++ {
				sb.WriteString(temp)
				continue
			}
			continue
		}
		temp = string(v)
		sb.WriteRune(v)
	}
	if len(sb.String()) == 0 && len(s) != 0 {
		fmt.Println("Некорректная строка")
		return "", fmt.Errorf("Error occure while input string is incorrect")
	}
	return sb.String(), nil
}
