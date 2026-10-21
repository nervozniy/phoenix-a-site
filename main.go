package main

import (
	"fmt"
)

func main() {
	const usdToRub = 95.50 // рублей в одном долларе
	const eurToUsd = 1.08  // долларов в одном евро

	// Конвертация из рублей в доллары
	fmt.Println("=== Конвертер валют ===")
	fmt.Print("Введите сумму в рублях (RUB): ")
	
	var rub float64 // Объявляем переменную для рублей
	_, err := fmt.Scanln(&rub)
	if err != nil || rub < 0 {
		fmt.Println("Ошибка: введено некорректное число")
		return
	}

	usdFromRub := rub / usdToRub
	fmt.Printf("%.2f RUB = %.2f USD\n\n", rub, usdFromRub)

	// Конвертация из долларов в евро
	var usd float64
	fmt.Print("Введите сумму в долларах (USD): ")
	
	_, err = fmt.Scanln(&usd)
	if err != nil || usd < 0 {
		fmt.Println("Ошибка: введено некорректное число")
		return
	}

	eurFromUsd := usd / eurToUsd
	fmt.Printf("%.2f USD = %.2f EUR\n", usd, eurFromUsd)
}
