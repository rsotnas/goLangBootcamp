// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	var total int
	var showCount string
	var firstNumber int
	var err error
	var secondNumber int
	var err2 error

	firstNumber, err = strconv.Atoi(os.Args[1])
	secondNumber, err2 = strconv.Atoi(os.Args[2])

	if err != nil || err2 != nil {
		return
	}

	for i := firstNumber; i <= secondNumber; i++ {

		if i%2 == 0 {
			showCount += strconv.Itoa(i)
			total += i

			if i != secondNumber {
				showCount += " + "
			} else {
				showCount += " = "
				showCount += strconv.Itoa(total)
			}
		}
	}
	fmt.Println(showCount)
}
