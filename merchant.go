package main

import "fmt"

var quant_2 int = 10 // Par exemple, initialisez avec une quantité de 10 potions

var wolfFur obj = obj{5, "Fourrure de Loup", 1, "Matériel"}
var trollSkin obj = obj{6, "Peau de Troll", 1, "Matériel"}
var boarLeather obj = obj{7, "Cuir de Sanglier", 1, "Matériel"}
var ravenFeather obj = obj{8, "Plume de Corbeau", 1, "Matériel"}

// Coûts des articles
var healthPotPrice int = 3
var poisonPotPrice int = 6
var fireSpellBookPrice int = 25
var wolfFurPrice int = 4
var trollSkinPrice int = 7
var boarLeatherPrice int = 3
var ravenFeatherPrice int = 1

// Disponibilité des articles
var healthPotAvailable bool = true
var poisonPotAvailable bool = true
var fireSpellBookBought bool = false
var wolfFurAvailable bool = true
var trollSkinAvailable bool = true
var boarLeatherAvailable bool = true
var ravenFeatherAvailable bool = true

func (u *character) accessMerchant() {
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"
	if healthPotAvailable || poisonPotAvailable || !fireSpellBookBought || wolfFurAvailable || trollSkinAvailable || boarLeatherAvailable || ravenFeatherAvailable {
		fmt.Printf("╒══════════╡%sMarchand%s╞══════════╕\n \tPurse : %d\n", yellow, reset, u.purse)
		if healthPotAvailable {
			fmt.Printf(" %s1.%s - Potion de vie (%d) ⨯ 1\n", yellow, reset, healthPotPrice)
		} else {
			fmt.Println(" ̶1̶.̶ ̶-̶ ̶P̶o̶t̶i̶o̶n̶ ̶d̶e̶ ̶v̶i̶e̶ ̶(̶g̶r̶a̶t̶u̶i̶t̶)̶ ̶⨯̶ ̶0̶")
		}
		if poisonPotAvailable {
			fmt.Printf(" %s2.%s - Potion de poison (%d) ⨯ 1\n", yellow, reset, poisonPotPrice)
		} else {
			fmt.Println(" ̶2̶.̶ ̶-̶ ̶P̶o̶i̶s̶o̶n̶ ̶P̶o̶t̶ ̶(̶0̶)̶ ̶⨯̶ ̶0̶")
		}
		if !fireSpellBookBought {
			fmt.Printf(" %s3.%s - Livre de Sort: Boule de Feu (%d pièces d'or)\n", yellow, reset, fireSpellBookPrice)
		} else {
			fmt.Println(" ̶3̶.̶ ̶-̶ ̶L̶i̶v̶r̶e̶ ̶d̶e̶ ̶S̶o̶r̶t̶:̶ ̶B̶o̶u̶l̶e̶ ̶d̶e̶ ̶F̶e̶u̶")
		}
		if wolfFurAvailable {
			fmt.Printf(" %s4.%s - Fourrure de Loup (%d pièces d'or)\n", yellow, reset, wolfFurPrice)
		} else {
			fmt.Println(" ̶4̶.̶ ̶-̶ ̶F̶o̶u̶r̶r̶u̶r̶e̶ ̶d̶e̶ ̶L̶o̶u̶p̶")
		}
		if trollSkinAvailable {
			fmt.Printf(" %s5.%s - Peau de Troll (%d pièces d'or)\n", yellow, reset, trollSkinPrice)
		} else {
			fmt.Println(" ̶5̶.̶ ̶-̶ ̶P̶e̶a̶u̶ ̶d̶e̶ ̶T̶r̶o̶l̶l̶")
		}
		if boarLeatherAvailable {
			fmt.Printf(" %s6.%s - Cuir de Sanglier (%d pièces d'or)\n", yellow, reset, boarLeatherPrice)
		} else {
			fmt.Println(" ̶6̶.̶ ̶-̶ ̶C̶u̶i̶r̶ ̶d̶e̶ ̶S̶a̶n̶g̶l̶i̶e̶r̶")
		}
		if ravenFeatherAvailable {
			fmt.Printf(" %s7.%s - Plume de Corbeau (%d pièces d'or)\n", yellow, reset, ravenFeatherPrice)
		} else {
			fmt.Println(" ̶7̶.̶ ̶-̶ ̶P̶l̶u̶m̶e̶ ̶d̶e̶ ̶C̶o̶r̶b̶e̶a̶u̶")
		}
		fmt.Println("╘══════════════════════════════╛")
		fmt.Printf("%s⎸%s'exit'%s\tpour quitter le marchand\n", yellow, red, reset)
	} else {
		fmt.Printf("╒══════════╡%sMarchand%s╞══════════╕\n \tPurse : %d\n", yellow, reset, u.purse)
		fmt.Println("Le marchand n'a plus rien à proposer.")
		fmt.Println("╘══════════════════════════════╛")
		fmt.Printf("%s⎸%s'exit'%s\tpour quitter le marchand\n", yellow, red, reset)
	}

	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "1":
		if healthPotAvailable && u.purse >= healthPotPrice {
			u.addInventory(obj1)      // Ajoute la potion de vie
			u.purse -= healthPotPrice // Déduit le coût de la potion de vie
			clear()
			fmt.Println("Vous avez acheté une Potion de vie !")
			healthPotAvailable = false // Potion n'est plus disponible après l'achat
		} else if u.purse < healthPotPrice {
			clear()
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cette potion.")
		} else {
			clear()
			fmt.Println("Le marchand n'a plus de potions de vie.")
		}
		u.accessMerchant()

	case "2":
		if poisonPotAvailable && u.purse >= poisonPotPrice {
			u.addInventory(obj2)      // Ajoute une potion de poison
			u.purse -= poisonPotPrice // Déduit le coût de la potion de poison
			fmt.Printf("Vous avez acheté une Potion de poison pour %d pièces d'or ! Il vous reste %d pièces d'or.\n", poisonPotPrice, u.purse)
			quant_2-- // Réduit la quantité disponible chez le marchand
			if quant_2 == 0 {
				poisonPotAvailable = false // Potion n'est plus disponible après avoir été épuisée
			}
		} else if u.purse < poisonPotPrice {
			clear()
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cette potion.")
		} else {
			clear()
			fmt.Println("Le marchand n'a plus de potions de poison.")
		}
		u.accessMerchant()

	case "3":
		if u.purse >= fireSpellBookPrice && !fireSpellBookBought {
			u.addInventory(fireSpellBook) // Ajoute le Livre de Sort dans l'inventaire
			u.purse -= fireSpellBookPrice // Déduit le coût du livre de la bourse
			fireSpellBookBought = true    // Le livre est marqué comme acheté
			clear()
			fmt.Printf("Vous avez acheté le Livre de Sort: Boule de Feu pour %d pièces d'or.\n", fireSpellBookPrice)
		} else if fireSpellBookBought {
			clear()
			fmt.Println("Vous avez déjà acheté ce livre.")
		} else {
			clear()
			fmt.Println("Vous n'avez pas assez d'argent.")
		}
		u.accessMerchant()
	case "4":
		if wolfFurAvailable && u.purse >= wolfFurPrice {
			u.addInventory(wolfFur) // Ajoute la Fourrure de Loup
			u.purse -= wolfFurPrice // Déduit le coût de la fourrure
			clear()
			fmt.Printf("Vous avez acheté une Fourrure de Loup pour %d pièces d'or ! Il vous reste %d pièces d'or.\n", wolfFurPrice, u.purse)
			wolfFurAvailable = false // Fourrure n'est plus disponible après l'achat
		} else if u.purse < wolfFurPrice {
			clear()
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cette fourrure.")
		} else {
			clear()
			fmt.Println("Le marchand n'a plus de fourrures de loup.")
		}
		u.accessMerchant()

	case "5":
		if trollSkinAvailable && u.purse >= trollSkinPrice {
			u.addInventory(trollSkin) // Ajoute la Peau de Troll
			u.purse -= trollSkinPrice // Déduit le coût de la peau de troll
			clear()
			fmt.Printf("Vous avez acheté une Peau de Troll pour %d pièces d'or ! Il vous reste %d pièces d'or.\n", trollSkinPrice, u.purse)
			trollSkinAvailable = false // Peau n'est plus disponible après l'achat
		} else if u.purse < trollSkinPrice {
			clear()
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cette peau.")
		} else {
			clear()
			fmt.Println("Le marchand n'a plus de peaux de troll.")
		}
		u.accessMerchant()

	case "6":
		if boarLeatherAvailable && u.purse >= boarLeatherPrice {
			u.addInventory(boarLeather) // Ajoute le Cuir de Sanglier
			u.purse -= boarLeatherPrice // Déduit le coût du cuir
			clear()
			fmt.Printf("Vous avez acheté un Cuir de Sanglier pour %d pièces d'or ! Il vous reste %d pièces d'or.\n", boarLeatherPrice, u.purse)
			boarLeatherAvailable = false // Cuir n'est plus disponible après l'achat
		} else if u.purse < boarLeatherPrice {
			clear()
			fmt.Println("Vous n'avez pas assez d'argent pour acheter ce cuir.")
		} else {
			clear()
			fmt.Println("Le marchand n'a plus de cuirs de sanglier.")
		}
		u.accessMerchant()

	case "7":
		if ravenFeatherAvailable && u.purse >= ravenFeatherPrice {
			u.addInventory(ravenFeather) // Ajoute la Plume de Corbeau
			u.purse -= ravenFeatherPrice // Déduit le coût de la plume
			clear()
			fmt.Printf("Vous avez acheté une Plume de Corbeau pour %d pièces d'or ! Il vous reste %d pièces d'or.\n", ravenFeatherPrice, u.purse)
			ravenFeatherAvailable = false // Plume n'est plus disponible après l'achat
		} else if u.purse < ravenFeatherPrice {
			clear()
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cette plume.")
		} else {
			clear()
			fmt.Println("Le marchand n'a plus de plumes de corbeau.")
		}
		u.accessMerchant()

	case "exit":
		clear()
		fmt.Println("Vous quittez le marchand.")
		return

	default:
		clear()
		fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
		u.accessMerchant()
	}
}
