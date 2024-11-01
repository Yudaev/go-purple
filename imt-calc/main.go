package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	for {
		fmt.Println("Калькулятор ИМТ")
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			//fmt.Println(err)
			//continue
			panic("ERROR")
		}
		outputResult(IMT)
		if !checkRepeatCalculation() {
			break
		}
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш ИМТ: %.2f\n", IMT)
	fmt.Print(result)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userKg, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделать еще расчет (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
