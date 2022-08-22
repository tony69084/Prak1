package main

import (
	"fmt"
)

func main() {

	var num1, num2, num3, numCountMoreFive int // Объявляем переменные 3х чисел и кол-ва чисел больше пяти

	fmt.Println("Три числа.")

	fmt.Println("Введите первое число: ")
	fmt.Scan(&num1) // Ввод первого числа

	fmt.Println("Введите второе число: ")
	fmt.Scan(&num2) // Ввод второго числа

	fmt.Println("Введите третье число: ")
	fmt.Scan(&num3) // Ввод третьего числа

	if num1 >= 5 { // Проверка первого числа на то, что оно больше пяти
		numCountMoreFive++
	}

	if num2 >= 5 { // Проверка второго числа на то, что оно больше пяти
		numCountMoreFive++
	}

	if num3 >= 5 { // Проверка третьего числа на то, что оно больше пяти
		numCountMoreFive++
	}

	fmt.Println("Среди введённых чисел", numCountMoreFive, "больше или равны 5.")

}
