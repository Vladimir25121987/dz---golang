package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("КАЛЬКУЛЯТОР")
	results := []float64{}
	for {
	number := scanNumber()
	results = append(results, number)
	fmt.Println("Продолжить ввод чисел? (да/нет): ")
	var cont string
	fmt.Scan(&cont)
	if cont != "да" {
		break
	}
}
	operation := scanOperation()
	switch operation {
	case "AVG":
		fmt. Printf("Результат вычислений: %.f\n", calculateAvg(results))
	case "SUM":
		fmt. Printf("Результат вычислений: %.f\n", calculateSum(results))
	case "MED":
		fmt. Printf("Результат вычислений: %.f\n", calculateMed(results))
	default:
		fmt. Println("Нет операции. Попробуйте заново")
	 }	
   }

func scanOperation() string {
	var operation string
	fmt.Println("Введите нужную операцию(AVG, SUM, MED): ")
	fmt.Scan(&operation)
	return operation
}

func scanNumber() float64 {
	var number float64
	fmt.Println("Введите числа через запятую: ")
	fmt.Scan(&number)
	return number
}

func calculateAvg(results []float64) float64 {
	var sum float64
	for _, number := range results {
		sum += number
	}
	return sum/ float64(len(results))
}

func calculateSum(results []float64) float64 {
	var sum float64
	for _, namber := range results {
		sum += namber
	}
	return sum
}

func calculateMed(results []float64) float64 {
	resultsCopy := make([]float64, len(results))
	copy(resultsCopy, results)
	sort.Float64s(resultsCopy)
	n:= len(resultsCopy)
	if  n % 2 != 0 {
		return resultsCopy[n / 2]
	}
        return (resultsCopy[n / 2 - 1] + resultsCopy[n / 2])/ 2 
	}
