package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fields               int
	delimiter, separated bool
)

func main() {
	var (
		text string
		err  error
	)
	initFlags()
	for text != "\n" {
		text, err = readStdin()
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case delimiter:
		case separated:
		default:
			if fields > 0 {
				fmt.Printf("%s\n", cut(text, fields))
				continue
			} else if fields < 0 {
				fmt.Printf("cut: fields are numbered from 1; your fields: %d\n", fields)
				os.Exit(1)
			}
			fmt.Printf("cut: you must specify a list of bytes, characters, or fields\n")
			os.Exit(1)
		}
	}

}

func initFlags() {
	flag.IntVar(&fields, "f", 0, "select only these fields also print any line that contains no delimiter character, unless the -s option is specified")
	flag.BoolVar(&delimiter, "d", false, "use DELIM instead of TAB for field delimiter")
	flag.BoolVar(&separated, "s", false, "only strings with separated")
	flag.Parse()
}

func readStdin() (string, error) {
	var text strings.Builder
	reader := bufio.NewReader(os.Stdin)
	text1, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text.WriteString(text1)
	return text.String(), nil
}

func cut(text string, fields int) []string {
	strs := strings.Split(text, "\t")
	result := make([]string, 1)
	if len(strs) > 1 {
		for _, v := range strs {
			if strings.Contains(v, "\n") {
				result = append(result, strs[fields])
			}
		}
	}
	return result
}
