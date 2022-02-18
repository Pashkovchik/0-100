package main

import (
	"fmt"
)

func main() {
	var min, max, middle, number int
	fmt.Println("Загадайте число от 0 до 100: я угадаю его за 7 попыток! Чтобы продолжить, введите любое значение:")
	number = 46
	fmt.Println("Ваше число", number)
	min = 0
	max = 101
	middle = (max + min) / 2
	for i := 0; i < 7; i++ {
		if number != middle {
			switch {
			case number > middle:
				fmt.Printf("Ваше число больше %v, вычисляю ещё\n", middle)
				min = middle
				middle = (max + min) / 2
			case number < middle:
				fmt.Printf("Ваше число меньше %v, вычисляю ещё\n", middle)
				max = middle
				middle = (max + min) / 2
			}
		} else {
			fmt.Println("Ура! Я угадал за", i+1, "шагов! Твое число:", middle)
			break
		}
	}
}
