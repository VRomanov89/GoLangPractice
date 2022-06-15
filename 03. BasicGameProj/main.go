package main

import (
	"fmt"

	"github.com/vladromanov/monster/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := "" // "Player" || "Monster"

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	//User actions
	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	fmt.Println(userChoice)

	return ""
}

func endGame() {

}
