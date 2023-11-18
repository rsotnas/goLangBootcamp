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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://www.merriam-webster.com/words-at-play/why-y-is-sometimes-a-vowel-usage
//  Check out: https://www.dictionary.com/e/w-vowel/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func main() {
	args := os.Args
	vowels := "aeiouAEIOU"
	consonants := "BCDFGHJKLMNPQRSTVXZbcdfghjklmnpqrstvxz"
	semivowels := "yYWw"

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	if strings.IndexAny(args[1], vowels) == 0 {
		fmt.Printf("\"%s\" is a vowel.\n", args[1])
	} else if strings.IndexAny(args[1], consonants) == 0 {
		fmt.Printf("\"%s\" is a consonant.\n", args[1])
	} else if strings.IndexAny(args[1], semivowels) == 0 {
		fmt.Printf("\"%s\" sometimes is a vowel, sometimes not.\n", args[1])
	}

}
