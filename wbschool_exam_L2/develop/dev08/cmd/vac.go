/*
Copyright © 2022 s1ovac <ilya.pereverzev123@gmail.com>

*/

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	color "github.com/TwiN/go-color"
)

func main() {
	for {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\r%s ", color.Green, path)
		cmd, cmdArgs := readStringFromStd()
		switch cmd {
		case "cd":
			c := exec.Command("cd", cmdArgs[0])
			c.Run()
		case "pwd":

		case "echo":
			fmt.Printf("%s\n%v%s\n", color.Green, strings.Join(cmdArgs, " "), path)
		case "kill":

		case "ps":

		default:
			fmt.Printf("%s: command not found: %v\n%s\n", cmd, color.Green, path)
		}
	}
}

func readStringFromStd() (string, []string) {
	reader := bufio.NewReader(os.Stdin)
	cmd, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	cmd = strings.TrimSpace(cmd)
	if arr := strings.Split(cmd, " "); len(arr) > 1 {
		return arr[0], arr[1:]
	}
	return cmd, nil
}
