package main

import (
	"errors"
	"fmt"
)
		var inputCurrency string
		var outputCurrency string
		var inputAmount float64
		var outputAmount float64
func main() {
	fmt.Println("КОНВЕРТЕР ВАЛЮТ")
	currencyMap := map[string]float64{
		"UsdToEur": 0.89,
		"UsdToRub": 91.61,
		"EurToRub": 103.47,
		"EurToUsd": 1.12,
		"RubToUsd": 0.01,
		"RubToEur": 0.009,
	}
		fmt.Println("Введите сумму конвертации: ")
		fmt.Scan(&inputAmount)
		amount, err := scanAmount()
		if err != nil{
			fmt.Print("Ошибка: не удалось преобразовать введенное значение в число", err)
			return
		}
		fmt.Printf("Вы ввели число: %2f\n", amount)

	for {
			inputCurrency, err = scanCurrancy("Введите исходную валюту(EUR, USD, RUB): ")
		if err != nil {
			fmt.Println(err)
			continue
		}
			break
	}
	for {
			inputAmount, err = scanAmount()
		if err != nil {
			fmt.Println(err)
			continue
		}
			break
	}
	for {
			outputCurrency, err = scanCurrancy("Введите целевцю валюту(EUR, USD, RUB):  ")
		if err != nil {
			fmt.Println(err)
			continue
		}
			break
	}
	
		outputAmount = convert(inputAmount)
		fmt.Scanf("%.2f", outputAmount)
}

func scanCurrancy(prefixPrint string) (string, error) {
	var currency string
	fmt.Print("prefixPrint")
	fmt.Scanln(&currency)
	if currency != "USD" && currency != "EUR" && currency != "RUB"{
		return "",errors.New("WRONG_CURRENCY")
	}
	return currency, nil
}

func scanAmount() (float64, error) {
	var amount float64
	fmt.Print("Введите сумму конвертации")
	fmt.Scanln(&amount)
	if amount < 0 {
return 0, errors.New("Введено значение меньше 0. Введите положительное значение")
	}
	return amount, nil
}

func convert(currencyMap map[string]float64, inputCurrency string, outputCurrency string, inputAmount float64) float64 {
switch inputCurrency {
case "USD":
	if outputCurrency == "EUR"{
		return inputAmount * currencyMap["UsdToEur"]
	}
	if outputCurrency == "RUB"{
		return inputAmount * currencyMap["UsdToEur"]
	}
case "EUR":
	if outputCurrency == "USD"{
		return inputAmount * currencyMap["EurToUsd"]
	}
	if outputCurrency == "RUB"{
		return inputAmount * currencyMap["EurToRub"]
	}
case "RUB":
	if outputCurrency == "EUR"{
		return inputAmount * currencyMap["RubToEur"]
	}
	if outputCurrency == "USD"{
		return inputAmount * currencyMap["RubToUsd"]
	}
  }
return inputAmount
}