package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Структура для сохранения результатов
type concurrences struct {
	number rune
	count  int
}

// Функция для подсчета элемента
func inc(num rune) {
	changed := false
	for i := 0; i < len(conc); i++ {
		if conc[i].number == num {
			conc[i].count++
			changed = true
			break
		}
	}
	if !changed {
		conc = append(conc, concurrences{num, 1})
	}
}

// Проверка условия задачи на ввод, только на символы алфавитов.

func LetterOnly(s string) bool {
	SymbLet := regexp.MustCompile("[a-zA-Zа-яА-Я]{1,}")
	OrigStr := SymbLet.FindAllString(s, -1)
	NewString := strings.Join(OrigStr, "")
	if strings.Compare(s, NewString) == 0 {
		return true
	} else {
		return false
	}
}

var conc = []concurrences{}

func main() {
	var str2 string

	//Считываем строку целиком со всеми символами и проверяем на условие:
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Каким-то образом приехала не строка")
	}
	str := strings.TrimSuffix(s, "\r\n")
	if LetterOnly(str) == false {
		fmt.Println("Не соответствует условию задачи, вводить нужно только кириллицу или латиницу без пробелов")
	} else {
		arr := []rune(str)
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

		//Сохраняем отсортированную строку и конфертим в массив рун для передачи в структуру:
		str2 = string(arr)
		arr2 := []rune(str2)

		for i := 0; i < len(arr2); i++ {
			inc(arr2[i])
		}
		//Выводим форматированно из своей структуры:
		for i := 0; i < len(conc); i++ {
			if conc[i].count >= 1 {
				fmt.Printf("%s%d", string(conc[i].number), conc[i].count)
			}
		}
	}
}
