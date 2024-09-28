package main

import (
	"errors"
	"fmt"
	"strconv"
)

    const UsdToEur = 0.89 
	const UsdToRub = 91.61
    var EurToRub = UsdToRub / UsdToEur
    var EurToUsd = 1 / UsdToEur
    var RubToUsd = 1 / UsdToRub
	var RubToEur = 1 / EurToRub

func main() {
		var inputCurrency string
		var outputCurrency string
		var inputAmount float64
		var outputAmount float64
		var err error
		var amountStr string
		fmt.Println("Введите сумму конвертации: ")
		fmt.Scan(&amountStr)
		amount, err := strconv.ParseFloat(amountStr, 64)
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
			inputAmount, err = scanAmount("Введите сумму конвертации: ")
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
	
		outputAmount = convert(inputAmount, inputCurrency, outputCurrency)
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

func scanAmount(prefixPrint float64) (float64, error) {
	var amount float64
	fmt.Print("prefixPrint")
	fmt.Scanln(&amount)
	if amount < 0 {
return 0, errors.New("Введено значение меньше 0")
	}
	return amount, nil
}

func convert(inputAmount float64, inputCurrency string, outputCurrency string) float64 {
switch inputCurrency {
case "USD":
	if outputCurrency == "EUR"{
		return inputAmount * UsdToEur
	}
	if outputCurrency == "RUB"{
		return inputAmount * UsdToRub
	}
case "EUR":
	if outputCurrency == "USD"{
		return inputAmount * EurToUsd
	}
	if outputCurrency == "RUB"{
		return inputAmount * EurToRub
	}
case "RUB":
	if outputCurrency == "EUR"{
		return inputAmount * RubToEur
	}
	if outputCurrency == "USD"{
		return inputAmount * RubToUsd
	}
  }
return inputAmount
}