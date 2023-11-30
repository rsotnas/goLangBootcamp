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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
	args := os.Args

	rand.NewSource(time.Now().UnixNano())

	computerChoice := rand.Intn(11)

	fmt.Println(computerChoice)
	userChoice, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println("Wrong choice")
	}

	if userChoice == computerChoice {
		fmt.Println("User wins")
	} else {
		fmt.Println("Computer wins")
	}
}
