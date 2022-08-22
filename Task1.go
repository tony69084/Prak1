package main

import (
	"fmt"
)

func main() {
	var exam1, exam2, exam3 int // Объявляем переменные результата каждого из 3х экзаменов

	fmt.Println("Баллы ЕГЭ.")
	fmt.Println("Введите результат первого экзамена: ")
	fmt.Scan(&exam1) // Вводим результат 1 экзамена

	fmt.Println("Введите результат второго экзамена: ")
	fmt.Scan(&exam2) // Вводим результат 2 экзамена

	fmt.Println("Введите результат третьего экзамена: ")
	fmt.Scan(&exam3) // Вводим результат 3 экзамена

	passPointsum := 275 // Объявляем переменную - Сумма проходных баллов
	fmt.Println("Сумма проходных баллов:", passPointsum)

	examSum := exam1 + exam2 + exam3 // Объявляем переменную - Кол-во набранных баллов
	fmt.Println("Количество набранных баллов:", examSum)

	if examSum >= passPointsum { // Условие проверки возможности поступления
		fmt.Println("Вы поступили")

	} else {
		fmt.Println("Вы не поступили")

	}

}
