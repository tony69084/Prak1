package main

import (
	"fmt"
)

func main() {

	var num1, num2, num3 int // Объявляем переменные 3х чисел

	fmt.Println("Три числа.")

	fmt.Println("Введите первое число: ")
	fmt.Scan(&num1) // Ввод первого числа

	fmt.Println("Введите второе число: ")
	fmt.Scan(&num2) // Ввод второго числа

	fmt.Println("Введите третье число: ")
	fmt.Scan(&num3) // Ввод третьего числа

	if num1 > 5 || num2 > 5 || num3 > 5 { // Условие того,что есть ли среди введенных чисел, числа больше пяти
		fmt.Println("Среди введённых чисел есть число больше 5.")
	} else {
		fmt.Println("Среди введённых чисел нет числа больше 5.")
	}

}
