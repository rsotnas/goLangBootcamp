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
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {

	winMessages := [2]string{"YOU WON", "YOU'RE AWESOME"}
	lostMessages := [2]string{"YOU LOST. TRY AGAIN?", "LOSER!"}

	args := os.Args

	rand.NewSource(time.Now().UnixNano())

	userChoice1, err1 := strconv.Atoi(args[1])
	userChoice2, err2 := strconv.Atoi(args[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Wrong choice")
	}

	computerChoice := rand.Intn(int(math.Max(float64(userChoice1), float64(userChoice2))))
	fmt.Println(computerChoice)

	if userChoice1 == computerChoice || userChoice2 == computerChoice {
		fmt.Println(winMessages[rand.Intn(len(winMessages))])
	} else {
		fmt.Println(lostMessages[rand.Intn(len(lostMessages))])
	}

}
