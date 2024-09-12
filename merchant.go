package main

import "fmt"

var healthPotAvailable bool = true
var poisonPotAvailable bool = true
var price_2 int = 50
var quant_2 int = 3

func (u *character) accessMerchant() {
	clear()
	if healthPotAvailable || poisonPotAvailable {
		fmt.Printf("╒══════════╡Marchand╞══════════╕\n \tPurse : %d\n", u.purse)
		if healthPotAvailable {
			fmt.Println(" 1. - Health Pot (gratuit) ⨯ 1")
		} else {
			fmt.Println(" ̶1̶.̶ ̶-̶ ̶H̶e̶a̶l̶t̶h̶ ̶P̶o̶t̶ ̶(̶g̶r̶a̶t̶u̶i̶t̶)̶ ̶⨯̶ ̶0")
		}
		if poisonPotAvailable {
			fmt.Printf(" 2. - Poison Pot (%d) ⨯ %d\n", price_2, quant_2)
		} else {
			fmt.Println(" ̶2̶.̶ ̶-̶ ̶P̶o̶i̶s̶o̶n̶ ̶P̶o̶t̶ ̶(̶0̶)̶ ̶⨯̶ ̶0̶")
		}
		fmt.Println("╘══════════════════════════════╛")
		fmt.Println("⎸'exit'\tpour quitter le marchand")
	} else {
		fmt.Printf("╒══════════╡Marchand╞══════════╕\n\tPurse : %d\n", u.purse)
		fmt.Println("Le marchand n'a plus rien a proposer.")
		fmt.Println("╘══════════════════════════════╛")
		fmt.Println("⎸'exit'\tpour quitter le marchand")
	}

	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "1":
		if healthPotAvailable {
			u.addInventory(obj1) // Ajoute la potion de vie
			fmt.Println("Vous avez acheté une Potion de vie !")
			healthPotAvailable = false // Potion n'est plus disponible après l'achat
		} else {
			fmt.Println("Le marchand n'a plus de potions de vie.")
		}
		u.accessMerchant()

	case "2":
		if poisonPotAvailable && u.purse >= price_2 { // Vérifie si le joueur a assez d'argent
			u.addInventory(obj2) // Ajoute une potion de poison
			u.purse -= price_2   // Déduit 50 pièces d'or de la bourse
			fmt.Printf("Vous avez acheté une Potion de poison pour %d pièces d'or ! Il vous reste %d pièces d'or.\n", price_2, u.purse)
			quant_2-- // Réduit la quantité disponible chez le marchand
			if quant_2 == 0 {
				poisonPotAvailable = false // Potion n'est plus disponible après avoir été épuisée
			}
		} else if u.purse < price_2 {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cette potion.")
		} else {
			fmt.Println("Le marchand n'a plus de potions de poison.")
		}
		u.accessMerchant()

	case "exit":
		loop()

	default:
		fmt.Println("Choix non valide")
		u.accessMerchant()
	}
}
