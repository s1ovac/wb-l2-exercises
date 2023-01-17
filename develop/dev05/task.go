package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===
Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и go-lint.
*/

var (
	after   int
	before  int
	context int
	count   int
	ignore  bool
	invert  bool
	fixed   bool
	num     bool
)

func main() {
	initFlags()
	elems, err := readFile(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Grep(elems, os.Args[len(os.Args)-2]))
}

func Grep(elems []string, substr string) string {
	var str strings.Builder
	for i, v := range elems {
		switch true {
		case ignore:
			strings.ToLower(v)
			fallthrough
		case fixed:
			if strings.Compare(v, substr) == 0 {
				str.WriteString(v)
			}
		case num:
			if strings.Contains(v, substr) {
				str.WriteString(strconv.Itoa(i) + " " + substr)
			}
		default:
			if strings.Contains(v, substr) {
				str.WriteString(substr)
			}
		}
	}
	return str.String()
}

func readFile(filename string) ([]string, error) {
	var (
		found []string
	)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		found = append(found, scanner.Text())
	}
	return found, nil
}

func initFlags() {
	flag.IntVar(&after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.IntVar(&count, "c", 0, "количество строк")
	flag.BoolVar(&ignore, "i", false, "игнорировать регистр")
	flag.BoolVar(&invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&fixed, "F", false, "точное совпадение со строкой")
	flag.BoolVar(&num, "n", false, "печатать номер строки")
	flag.Parse()
}
