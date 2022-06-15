package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("Monster Slayer Game!")
	fmt.Println("Monster Slayer Game!")
	fmt.Println("Good Luck!!")
}

func ShowAvailableActions(isSpecialAttackAvailable bool) {
	fmt.Println("Please choose an action")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal Self")

	if isSpecialAttackAvailable {
		fmt.Println("(3) Special Attack")
	}
}
