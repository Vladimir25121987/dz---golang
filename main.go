package main

import "fmt"

	const UsdToEur = 0.89
	const UsdToRub = 92.71
	type courses map[string]map[string]float64

func main() {
	fmt.Println("КОНВЕРТЕР ВАЛЮТ")
	courses = courses{
		"USD": {
				"EUR": UsdToEur,
				"RUB": UsdToRub,
		},
		"EUR": {
				"USD": 1 / UsdToEur,
				"RUB": UsdToEur / UsdToRub,
		},
		"RUB": {
				"USD": 1 / UsdToRub,
				"EUR": UsdToEur / UsdToRub,
		},
	}
	inputCurrency := getinputСurrency()
	outputCurrency := getoutputСurrency()
	inputAmount := getinputAmount()
	convert(inputAmount, inputCurrency, outputCurrency)
	fmt.Println()
}
func getinputСurrency() string {
	var incur string
	fmt.Println("Введите исходную валюту(USD, EUR, RUB): ")
	for {
		fmt.Scan(&incur)
			if incur != "USD" || incur != "EUR" || incur != "RUB"{
		fmt.Println("Вы неправильно ввели исходную валюту. Повторите ввод.")
	} 
	return incur
	}	
}

func getoutputСurrency(inputCurrency string) string {
	var outcur string
	fmt.Println("Выберите желаемую валюту(USD, EUR, RUB): ")
	for {
		fmt.Scan(&outcur)
		if outcur != "USD" || outcur != "EUR" || outcur != "RUB"{
			fmt.Println("Вы неправильно ввели желаемую валюту. Повторите ввод.")
		} 
		if outcur == inputCurrency{
			fmt.Println("Исходная и желаемая валюты совпадают. Повторите ввод.")
		}
		return outcur
	}
}

func getinputAmount() float64 {
	var sum float64
	fmt.Println("Выберите сумму конвертации: ")
	for {
		fmt.Scan(&sum)
		if sum < 0 {
			fmt.Println("Ошибка. Введите положительное число")
		}
		return sum
	}	
}
	
func convert(courses courses , inputAmount float64, inputCurrency string, outputCurrency string) float64 {
	result = inputAmount * courses[inputCurrency][outputCurrency]
return result
}