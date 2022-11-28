package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	red := color.New(color.FgRed, color.Bold).PrintFunc()
	yellow := color.New(color.FgYellow).PrintFunc()
	green := color.New(color.FgGreen).PrintFunc()

	rand.Seed(time.Now().UnixNano())
	randomNum := randInt(0, 100)
	tries := 0

	for {
		green("Enter guess: ")
		var guess string
		fmt.Scanln(&guess)

		if guess == "q" || guess == "" {
			break
		}

		guessInt64, err := strconv.ParseInt(guess, 10, 8)
		if err != nil {
			red("Invalid input\n")
			continue
		}
		var guessInt int = int(guessInt64)

		if !(guessInt >= 0 && guessInt <= 100) {
			red("Input must be between 0 and 100\n")
			continue
		} else {
			tries++
		}

		if guessInt > randomNum {
			yellow("Try a smaller number\n")
		} else if guessInt < randomNum {
			yellow("Try a larger number\n")
		} else {
			green("Correct! You got it in ", tries, " tries! 🎉 \n")
			break
		}
	}
}
