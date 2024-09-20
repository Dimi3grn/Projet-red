package main

import (
	"fmt"
	"os"
)

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
	} else if valeur == "m" {
		clear()
		accessMenu()
	} else if valeur == "exit" {
		clear()
		fmt.Println("exit the game successfully")
		os.Exit(0)
	} else if valeur == "e" {
		clear()
		fmt.Println("exit the game successfully")
		os.Exit(0)
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
	fmt.Printf("\t║%sVous êtes dans le Menu%s║\n%s╭%s'stats'%s pour afficher les statistiques\n%s│%s'mark'%s pour afficher le marchand\n%s│%s'inv'%s pour afficher l'inventaire\n%s│%s'forge'%s pour accéder au forgeron\n%s│%s'train'%s pour accédé à l'entraînement\n%s│%s'combat'%s pour acceder à l'histoire principale\n%s╰%s'exit'%s pour quitter\n", yellow, reset, yellow, red, reset, yellow, red, reset, yellow, red, reset, yellow, red, reset, yellow, red, reset, yellow, red, reset, yellow, red, reset)
	fmt.Scan(&option)
	switch option {
	case "mark", "m":
		clear()
		MyChar.accessMerchant()
		loop()
	case "stats", "s":
		clear()
		MyChar.displayinfo()
		loop()
	case "inv", "i":
		clear()
		MyChar.accessInventory()
		rep := readTer()
		if rep == "hp" || rep == "h" {
			MyChar.takePot()
		}
		loop()
	case "forge", "f":
		clear()
		MyChar.accessBlacksmith() // Nouvelle fonction pour accéder au forgeron
		loop()
	case "train", "t":
		clear()
		MyChar.StartCombat() // Nouvelle fonction pour accéder au combat
		loop()
	case "combat", "c":
		clear()
		MyChar.getCurrentFight()
		loop()
	case "exit", "e":
		clear()
		loop()
	default:
		fmt.Println("commande invalide")
		accessMenu()
	}
}

func health_bar(hb_pv, hb_maxhp int) {

	yellow := "\033[33m"
	red := "\033[31m"
	cyan := "\033[36m"
	reset := "\033[0m"

	pixel := float64(hb_pv) / float64(hb_maxhp) * 10

	fmt.Printf("%s[%s", yellow, reset)
	for i := 0; i < 10; i++ {
		if float64(i) < pixel {
			fmt.Printf("%s▆", red)
		} else {
			fmt.Printf("%s▆", cyan)
		}
	}
	fmt.Printf("%s]%s - %s%d%s/%s%d%s hp", yellow, reset, red, hb_pv, reset, red, hb_maxhp, reset)
}
