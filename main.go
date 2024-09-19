package main

import "fmt"

func main() {
	Selection()
	setclasse() // Sets the player's class and initial inventory
	clear()     // Clears the screen
	fmt.Printf("La race vous étant aléatoirement attribuée est : %s\n", MyChar.classe)
	fmt.Printf("Vous avez donc : %d pv max\n", MyChar.maxHp)
	loop() // Starts the game loop

}
