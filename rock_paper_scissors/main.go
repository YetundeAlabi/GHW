package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var choices = map[string]string{
	"R": "Rock",
	"P": "Paper",
	"S": "Scissors",
}

func main() {
	// Seed the random number generator
	playerChoice := getPlayerChoice()
	computerChoice := getComputerChoice()

	fmt.Printf("You chose: %s\n", choices[playerChoice])
	fmt.Printf("Computer chose: %s\n", choices[computerChoice])

	result := determineWinner(playerChoice, computerChoice)
	fmt.Printf("Result: %s\n", result)
}

func getPlayerChoice() string {
	var choice string
	for {
	fmt.Print("Enter your choice(R for Rock, P for Paper, S for Scissors): ")
	
	fmt.Scanln(&choice)
	choice = strings.ToUpper(choice) //convert to uppercase to handle case sensitivity
	if choice == "R" || choice == "P" || choice == "S" {
		break
	} else {
		fmt.Println("Invalid choice. Please enter R, P, or S.")
	}
}
	return choice
}

func getComputerChoice() string {
	choices := []string{"R", "P", "S"}
	return choices[rand.Intn(len(choices))]
}

func determineWinner(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return "It's a tie!"
	}
	switch playerChoice {
	case "R":
		if computerChoice == "S" {
			return "You win!" 
		} else {
			return "Computer wins!"
		}
	case "P":
		if computerChoice == "R" {
			return "You win!" 
		} else {
			return "Computer wins!"
		}
	case "S":
		if computerChoice == "P" {
			return "You win!" 
		} else {
			return "Computer wins!"
		}
	default:
		return "Invalid choice"
	}
}