package main

import "fmt"

const UsdToEur = 0.9 
const UsdToRub = 93.36

type courses map[string]map[string]float64

func main() { 
fmt.Println("КОНВЕРТЕР ВАЛЮТ")
courses := courses {  
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

inputCurrency := getInputCurrency() 
outputCurrency := getOutputCurrency(inputCurrency) 
inputAmount := getInputAmount() 
convertedAmount := convert(courses, inputAmount, inputCurrency, outputCurrency) 
fmt.Printf("%.2f %s = %.2f %s\n", inputAmount, inputCurrency, convertedAmount, outputCurrency) 
}

func getInputCurrency() string {
	 var incur string 
	fmt.Println("Введите исходную валюту(USD, EUR, RUB): ") 
	for {  
		 fmt.Scan(&incur)  
		  if incur == "USD" || incur =="EUR" || incur == "RUB" { 
			   return incur  
			 } 
		    fmt.Println("Вы неправельно ввели исходную валюту. Повторите ввод.")
		} 
}

func getOutputCurrency(inputCurrency string) string { 
	var outcur string 
	fmt.Println("Выберите желаемую валюту(USD, EUR, RUB): ") 
	for {  
		fmt.Scan(&outcur)  
		if outcur == "USD" || outcur == "EUR" || outcur == "RUB" {    
		if outcur != inputCurrency {   
			  return outcur    
			  }  
			    fmt.Println("Исходная и желаемая валюты совпадают. Повторите ввод.")  
		} 
				fmt.Println("Вы неправельно ввели желаемую валюту. Повторите ввод.") 
	} 
}

func getInputAmount() float64 { 
	var sum float64
	  	fmt.Println("Введите сумму конвертации: ")
	for {   
		fmt.Scan(&sum) 
		if sum >= 0 {   
			 return sum   
			 } 
		fmt.Println("Ошибка. Введите положительное число") 
		} 
}

func convert(courses courses, inputAmount float64, inputCurrency string, outputCurrency string) float64 {
	 result := inputAmount * courses[inputCurrency][outputCurrency]
	  return result
 }