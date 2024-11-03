package main

import "fmt"

func main() {
	var transactions []float64
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println("Сумма всех транзакций:", calculateBalance(transactions))
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию (n для выхода): ")
	_, _ = fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	var transactionsSum float64
	for _, value := range transactions {
		transactionsSum += value
	}
	return transactionsSum
}
