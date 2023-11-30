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
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

func main() {

	winMessages := [2]string{"YOU WON", "YOU'RE AWESOME"}
	lostMessages := [2]string{"YOU LOST. TRY AGAIN?", "LOSER!"}

	args := os.Args

	rand.NewSource(time.Now().UnixNano())

	computerChoice := rand.Intn(11)

	fmt.Println(computerChoice)
	userChoice, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println("Wrong choice")
	}

	if userChoice == computerChoice {
		fmt.Println(winMessages[rand.Intn(len(winMessages))])
	} else {
		fmt.Println(lostMessages[rand.Intn(len(lostMessages))])
	}

}
