package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	userHeight := 1.86
	userKg := 110.0
	IMT := userKg / math.Pow(userHeight, IMTPower)

	fmt.Print(IMT)
}
