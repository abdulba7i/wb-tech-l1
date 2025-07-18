package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите первое большое число: ")
	aStr, _ := reader.ReadString('\n')
	aStr = strings.TrimSpace(aStr)

	fmt.Print("Введите второе большое число: ")
	bStr, _ := reader.ReadString('\n')
	bStr = strings.TrimSpace(bStr)

	a := new(big.Int)
	b := new(big.Int)

	if _, ok := a.SetString(aStr, 10); !ok {
		return
	}

	if _, ok := b.SetString(bStr, 10); !ok {
		return
	}

	ArifmeticOperations(a, b)
}

func ArifmeticOperations(a, b *big.Int) {
	sum := new(big.Int).Add(a, b)
	fmt.Println("Сумма: ", sum.String())

	diff := new(big.Int).Sub(a, b)
	fmt.Println("Разность: ", diff.String())

	product := new(big.Int).Mul(a, b)
	fmt.Println("Произведение: ", product.String())

	if b.Sign() != 0 {
		quotient := new(big.Int).Div(a, b)
		fmt.Println("Деление: ", quotient.String())
	} else {
		fmt.Println("Ошибка: делить на ноль нельзя")
	}
}
