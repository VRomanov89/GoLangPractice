package main

import "github.com/vladromanov/monster/interaction"

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

}

func endGame() {

}
