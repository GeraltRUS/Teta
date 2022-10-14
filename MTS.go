package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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
	return regexp.MustCompile(`^[A-Za-zА-Яа-я]*$`).MatchString(s)
}

var conc = []concurrences{}

func main() {
	var str2 string

	//Считываем строку целиком со всеми символами и проверяем на условие:

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
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
