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
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {

	args := os.Args
	var year int
	var err error

	if len(args) != 2 {
		fmt.Println("Give me a year number.")
		return
	} else {
		year, err = strconv.Atoi(args[1])
		// fmt.Println(year, err)
		if err != nil || year <= 0 || len(args[1]) != 4 {
			fmt.Printf("Wrong info: %d. Error: %q\n", year, err)
			return
		}
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			fmt.Printf("%d is a leap year!\n", year)
		} else {
			fmt.Printf("%d is not a leap year!\n", year)
		}
	}
}