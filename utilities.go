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
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Printf("%s╭%s'menu'%s pour acceder au menu\n%s╰%s'exit'%s pour quitter le jeu\n", yellow, red, reset, yellow, red, reset)
	valeur := readTer()
	if valeur == "menu" {
		clear()
		accessMenu()
	} else if valeur == "exit" {
		clear()
		fmt.Println("exit the game successfully")
	} else {
		clear()
		loop()
	}
}

func accessMenu() {
	clear()
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"
	var option string
	fmt.Printf("\t║%sVous êtes dans le Menu%s║\n%s╭%s'stats'%s pour afficher les satistiques\n%s│%s'mark'%s pour afficher le marchand\n%s│%s'inv'%s pour afficher l'inventaire\n%s╰%s'exit'%s pour quitter\n", yellow, reset, yellow, red, reset, yellow, red, reset, yellow, red, reset, yellow, red, reset)
	fmt.Scan(&option)
	switch option {
	case "dmg":
		MyChar.takeDamage(1)
	case "mark":
		clear()
		MyChar.accessMerchant()
		loop()
	case "stats":
		MyChar.displayinfo()
		loop()
	case "inv":
		clear()
		MyChar.accessInventory()
		rep := readTer()
		if rep == "hp" {
			MyChar.takePot()
		}
		loop()
	case "exit":
		clear()
		loop()
	default:
		fmt.Println("commande invalide")
		accessMenu()
	}
}
