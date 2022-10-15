package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Проверка условия задачи на ввод, только на символы алфавитов.
func strValid(s string) bool {
	return regexp.MustCompile(`^[A-Za-zА-Яа-я]*$`).MatchString(s)
}

// Функция сворачивания строки
func strSort(s string) (ss string) {
	sr := []rune(s)
	for sym := 'A'; sym <= 'я'; sym++ {
		if sym == '{' { //в unicode "{" идёт сразу после "z"
			sym = 'А' //русское А!
		}
		for i := range sr {
			if sr[i] == sym {
				count := strconv.Itoa(strings.Count(s, string(sym)))
				ss += string(sr[i]) + count
				break
			}
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	valid := strValid(str)
	if !valid {
		fmt.Println("Не соответствует условию задачи, вводить нужно только кириллицу или латиницу без пробелов")
	} else {
		fmt.Println(strSort(str))
	}
}
