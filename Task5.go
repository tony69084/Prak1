package main

import (
	"fmt"
)

func main() {

	var dayOfWeek, guestsCount int // Переменные дня недели, числа гостей и суммы чека
	var checkSum float64

	fmt.Println("Введите день недели:")
	fmt.Scan(&dayOfWeek) // Ввод дня недели

	fmt.Println("Введите число гостей:")
	fmt.Scan(&guestsCount) // Ввод числа гостей

	fmt.Println("Введите сумму чека:")
	fmt.Scan(&checkSum) // Ввод суммы чека

	if dayOfWeek == 1 { // Если понедельник
		discMon := checkSum * 0.1
		fmt.Println("Скидка по понедельниками:", discMon)
		checkSum -= discMon
	}

	if dayOfWeek == 5 && checkSum > 10000 { // Если пятница и сумма чека превышает 10 000 рублей
		discFri := checkSum * 0.05
		fmt.Println("Скидка по пятницам:", discFri)
		checkSum -= discFri
	}

	if guestsCount > 5 { // Если число гостей превышает пять человек
		serviceTax := checkSum * 0.1
		fmt.Println("Надбавка на обслуживание:", serviceTax)
		checkSum += serviceTax
	}

	fmt.Println("Сумма к оплате:", checkSum)

}
