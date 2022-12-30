package main

import "fmt"

var strategyOp func(int, int) int

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func main() {
	strategyOp = add
	fmt.Println(strategyOp(6, 3))
	strategyOp = sub
	fmt.Println(strategyOp(6, 3))
	strategyOp = mul
	fmt.Println(strategyOp(6, 3))
	strategyOp = div
	fmt.Println(strategyOp(6, 3))
}
