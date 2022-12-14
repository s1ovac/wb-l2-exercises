package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	column    int
	number    bool
	reverse   bool
	unique    bool
	month     bool
	backSpace bool
	check     bool
	numberH   bool
)

func initFlags() {
	flag.IntVar(&column, "k", 0, "choose column to sort (default value is 0)")
	flag.BoolVar(&number, "n", false, "sort by number")
	flag.BoolVar(&reverse, "r", false, "reverse sort")
	flag.BoolVar(&unique, "u", false, "doesn't print recurring strings")
	flag.BoolVar(&month, "M", false, "sort by month")
	flag.BoolVar(&backSpace, "b", false, "ignore spaces at the back")
	flag.BoolVar(&check, "c", false, "check if data is sorted")
	flag.BoolVar(&numberH, "h", false, "sort by numbers with suffix")
	flag.Parse()
}

func main() {
	initFlags()
	result, err := readFile("/home/s1ovac/github.com/wb-l2-exercises/wbschool_exam_L2/develop/dev03/text.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func readFile(filename string) (result []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if column > 0 {
			strArr := strings.Split(scanner.Text(), " ")
			result = append(result, strArr[column])
		} else {
			result = append(result, scanner.Text())
		}
	}
	sort.Strings(result)
	return
}
