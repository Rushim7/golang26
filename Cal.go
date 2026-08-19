// Task: Creating a package on calculations
package main

import (
	"fmt"
	"time"
)

func add(a, b float64) {
	fmt.Println("Addition:", a+b)
}

func sub(a, b float64) {
	fmt.Println("Subtraction:", a-b)
}

func mul(a, b float64) {
	fmt.Println("Multiplication:", a*b)
}

func div(a, b float64) {
	fmt.Println("Division:", a/b)
}

func Calculations() {

	var n, m float64

	fmt.Print("Enter first number: ")
	fmt.Scan(&n)

	fmt.Print("Enter second number: ")
	fmt.Scan(&m)

	go add(n, m)
	go sub(n, m)
	go mul(n, m)
	go div(n, m)

	time.Sleep(time.Second)

}
