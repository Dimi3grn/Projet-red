package main

import "fmt"

var healthPotAvailable bool = true
var poisonPotAvailable bool = true
var price_2 int = 50
var quant_2 int = 3
var fireSpellBook obj = obj{4, "Livre de Sort: Boule de Feu", 1, "Book"}
var fireSpellBookPrice int = 100     // Prix du livre
var fireSpellBookBought bool = false // Statut d'achat

func (u *character) accessMerchant() {
	clear()
	if healthPotAvailable || poisonPotAvailable || !fireSpellBookBought {
		fmt.Printf("╒══════════╡Marchand╞══════════╕\n \tPurse : %d\n", u.purse)
		if healthPotAvailable {
			fmt.Println(" 1. - Health Pot (gratuit) ⨯ 1")
		} else {
			fmt.Println(" ̶1̶.̶ ̶-̶ ̶H̶e̶a̶l̶t̶h̶ ̶P̶o̶t̶ ̶(̶g̶r̶a̶t̶u̶i̶t̶)̶ ̶⨯̶ ̶0̶")
		}
		if poisonPotAvailable {
			fmt.Printf(" 2. - Poison Pot (%d) ⨯ %d\n", price_2, quant_2)
		} else {
			fmt.Println(" ̶2̶.̶ ̶-̶ ̶P̶o̶i̶s̶o̶n̶ ̶P̶o̶t̶ ̶(̶0̶)̶ ̶⨯̶ ̶0̶")
		}
		if !fireSpellBookBought {
			fmt.Printf(" 3. - Livre de Sort: Boule de Feu (%d pièces d'or)\n", fireSpellBookPrice)
		} else {
			fmt.Println(" ̶3̶.̶ ̶-̶ ̶L̶i̶v̶r̶e̶ ̶d̶e̶ ̶S̶o̶r̶t̶:̶ ̶B̶o̶u̶l̶e̶ ̶d̶e̶ ̶F̶e̶u̶")
		}
		fmt.Println("╘══════════════════════════════╛")
		fmt.Println("╭'exit'\tpour quitter le marchand")
	} else {
		fmt.Printf("╒══════════╡Marchand╞══════════╕\n\tPurse : %d\n", u.purse)
		fmt.Println("Le marchand n'a plus rien a proposer.")
		fmt.Println("╘══════════════════════════════╛")
		fmt.Println("╭'exit'\tpour quitter le marchand")
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

	case "3":
		if u.purse >= fireSpellBookPrice && !fireSpellBookBought {
			u.addInventory(fireSpellBook) // Ajoute le Livre de Sort dans l'inventaire
			u.purse -= fireSpellBookPrice // Déduit le coût du livre de la bourse
			fireSpellBookBought = true    // Le livre est marqué comme acheté
			fmt.Printf("Vous avez acheté le Livre de Sort: Boule de Feu pour %d pièces d'or.\n", fireSpellBookPrice)
		} else if fireSpellBookBought {
			fmt.Println("Vous avez déjà acheté ce livre.")
		} else {
			fmt.Println("Vous n'avez pas assez d'argent.")
		}
		u.accessMerchant()

	case "exit":
		loop()

	default:
		fmt.Println("Choix non valide")
		u.accessMerchant()
	}
}
