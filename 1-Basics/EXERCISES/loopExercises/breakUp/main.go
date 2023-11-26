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
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	var total int
	var showCount string
	var firstNumber int
	var err error
	var secondNumber int
	var err2 error
	var i int

	firstNumber, err = strconv.Atoi(os.Args[1])
	secondNumber, err2 = strconv.Atoi(os.Args[2])

	if err != nil || err2 != nil {
		return
	}

	i = firstNumber

	for i <= secondNumber {

		if i%2 == 0 {
			showCount += strconv.Itoa(i)
			total += i

			if i != secondNumber {
				showCount += " + "
			} else {
				showCount += " = "
				showCount += strconv.Itoa(total)
				break
			}
		}
		i++
	}
	fmt.Println(showCount)
}
