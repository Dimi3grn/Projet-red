package main

import "fmt"

func clear() {
	fmt.Printf("\033[H\033[2J")
}

func readTer() string {
	var inputvalue string
	fmt.Scan(&inputvalue)
	return inputvalue
}

func loop() {
	fmt.Println("╭'menu' pour acceder au menu\n╰'exit' pour quitter le jeu")
	valeur := readTer()
	if valeur == "menu" {
		accessMenu()
	} else if valeur == "exit" {
		fmt.Println("exit the game successfully")
	} else {
		loop()
	}
}

func accessMenu() {
	clear()
	var option string
	fmt.Println("\tvous êtes dans le menu\n╭'stats' pour afficher les satistiques\n⎸'mark' pour afficher le marchand\n⎸'inv' pour afficher l'inventaire\n╰'exit' pour quitter")
	fmt.Scan(&option)
	switch option {
	case "dmg":
		MyChar.takeDamage(1)
	case "mark":
		MyChar.accessMerchant()
		loop()
	case "stats":
		MyChar.displayinfo()
		loop()
	case "inv":
		MyChar.accessInventory()
		fmt.Println("⎸'hp' pour récuperer hp à partir des Heatlh Pot")
		rep := readTer()
		if rep == "hp" {
			MyChar.takePot()
		}
		loop()
	case "exit":
		fmt.Println("exited")
		loop()
	default:
		fmt.Println("commande invalide")
		accessMenu()
	}
}
