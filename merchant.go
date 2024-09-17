package main

import "fmt"

var quant_healthPot int = 10
var quant_poisonPot int = 10
var quant_wolfFur int = 5
var quant_trollSkin int = 5
var quant_boarLeather int = 5
var quant_ravenFeather int = 5
var quant_spellBook int = 1 // Le livre de sort est unique
var wolfFur obj = obj{5, "Fourrure de Loup", 1, "Matériel", 0}
var trollSkin obj = obj{6, "Peau de Troll", 1, "Matériel", 0}
var boarLeather obj = obj{7, "Cuir de Sanglier", 1, "Matériel", 0}
var ravenFeather obj = obj{8, "Plume de Corbeau", 1, "Matériel", 0}
var quant_InvSpace int = 3

// Coûts des articles
var healthPotPrice int = 0
var poisonPotPrice int = 20
var fireSpellBookPrice int = 50
var wolfFurPrice int = 20
var trollSkinPrice int = 15
var boarLeatherPrice int = 7
var ravenFeatherPrice int = 5
var inventorySpacePrice int = 30

// Disponibilité des articles
var healthPotAvailable bool = true
var poisonPotAvailable bool = true
var fireSpellBookBought bool = false
var wolfFurAvailable bool = true
var trollSkinAvailable bool = true
var boarLeatherAvailable bool = true
var ravenFeatherAvailable bool = true
var inventorySpaceAvaible bool = true

func (u *character) accessMerchant() {
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"
	green := "\033[32m"
	if healthPotAvailable || poisonPotAvailable || !fireSpellBookBought || wolfFurAvailable || trollSkinAvailable || boarLeatherAvailable || ravenFeatherAvailable {
		fmt.Printf("╒══════════╡%sMarchand%s╞══════════╕\n \tPurse : %d\n", yellow, reset, u.purse)

		if healthPotAvailable {
			fmt.Printf(" %s1.%s - Potion de vie (%s%d $%s) ⨯ %d\n", yellow, reset, green, healthPotPrice, reset, quant_healthPot)
		} else {
			fmt.Println(" ̶1̶.̶ ̶-̶ ̶P̶o̶t̶i̶o̶n̶ ̶d̶e̶ ̶v̶i̶e̶ ̶(̶0̶ ̶$̶) ̶⨯̶ ̶0̶")
		}

		if poisonPotAvailable {
			fmt.Printf(" %s2.%s - Potion de poison (%s%d $%s) ⨯ %d\n", yellow, reset, green, poisonPotPrice, reset, quant_poisonPot)
		} else {
			fmt.Println(" ̶2̶.̶ ̶-̶ ̶P̶o̶t̶i̶o̶n̶ ̶d̶e̶ ̶p̶o̶i̶s̶o̶n̶ ̶(̶0̶ ̶$̶) ̶⨯̶ ̶0̶")
		}

		if !fireSpellBookBought {
			fmt.Printf(" %s3.%s - Livre de Sort: Boule de Feu (%s%d $%s)\n", yellow, reset, green, fireSpellBookPrice, reset)
		} else {
			fmt.Println(" ̶3̶.̶ ̶-̶ ̶L̶i̶v̶r̶e̶ ̶d̶e̶ ̶S̶o̶r̶t̶:̶ ̶B̶o̶u̶l̶e̶ ̶d̶e̶ ̶F̶e̶u̶ ̶(̶0̶ ̶$̶)")
		}

		if wolfFurAvailable {
			fmt.Printf(" %s4.%s - Fourrure de Loup (%s%d $%s) ⨯ %d\n", yellow, reset, green, wolfFurPrice, reset, quant_wolfFur)
		} else {
			fmt.Println(" ̶4̶.̶ ̶-̶ ̶F̶o̶u̶r̶r̶u̶r̶e̶ ̶d̶e̶ ̶L̶o̶u̶p̶ ̶(̶0̶ ̶$̶) ̶⨯̶ ̶0̶")
		}

		if trollSkinAvailable {
			fmt.Printf(" %s5.%s - Peau de Troll (%s%d $%s) ⨯ %d\n", yellow, reset, green, trollSkinPrice, reset, quant_trollSkin)
		} else {
			fmt.Println(" ̶5̶.̶ ̶-̶ ̶P̶e̶a̶u̶ ̶d̶e̶ ̶T̶r̶o̶l̶l̶ ̶(̶0̶ ̶$̶) ̶⨯̶ ̶0̶")
		}

		if boarLeatherAvailable {
			fmt.Printf(" %s6.%s - Cuir de Sanglier (%s%d $%s) ⨯ %d\n", yellow, reset, green, boarLeatherPrice, reset, quant_boarLeather)
		} else {
			fmt.Println(" ̶6̶.̶ ̶-̶ ̶C̶u̶i̶r̶ ̶d̶e̶ ̶S̶a̶n̶g̶l̶i̶e̶r̶ ̶(̶0̶ ̶$̶) ̶⨯̶ ̶0̶")
		}

		if ravenFeatherAvailable {
			fmt.Printf(" %s7.%s - Plume de Corbeau (%s%d $%s) ⨯ %d\n", yellow, reset, green, ravenFeatherPrice, reset, quant_ravenFeather)
		} else {
			fmt.Println(" ̶7̶.̶ ̶-̶ ̶P̶l̶u̶m̶e̶ ̶d̶e̶ ̶C̶o̶r̶b̶e̶a̶u̶ ̶(̶0̶ ̶$̶) ̶⨯̶ ̶0̶")
		}
		if inventorySpaceAvaible {
			fmt.Printf("\n %s0.%s - 5 espaces d'inventaire (%s%d $%s) ⨯ %d\n", yellow, reset, green, inventorySpacePrice, reset, quant_InvSpace)
		} else {
			fmt.Printf("\nCapacité d'inventaire maximale atteinte\n")
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
	case "0":
		if inventorySpaceAvaible && u.purse > inventorySpacePrice {
			u.invSize += 5
			u.purse -= inventorySpacePrice
			clear()
			fmt.Printf("vous avez maintenant %d espace d'inventaire\n", u.invSize)
			quant_InvSpace -= 1
			if quant_InvSpace == 0 {
				inventorySpaceAvaible = false
			}
		} else if u.purse < inventorySpacePrice {
			clear()
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cette option.")
		} else {
			clear()
			fmt.Println("Le marchand n'a plus d'espace d'inventaire")
		}
		u.accessMerchant()
	case "1":
		if healthPotAvailable && u.purse >= healthPotPrice {
			u.addInventory(obj1)      // Ajoute la potion de vie
			u.purse -= healthPotPrice // Déduit le coût de la potion de vie
			clear()
			fmt.Println("Vous avez acheté une Potion de vie !")
			quant_healthPot-- // Réduit la quantité disponible chez le marchand
			if quant_healthPot == 0 {
				healthPotAvailable = false // Potion n'est plus disponible après avoir été épuisée
			}
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
			quant_poisonPot-- // Réduit la quantité disponible chez le marchand
			if quant_poisonPot == 0 {
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
			quant_wolfFur-- // Réduit la quantité disponible chez le marchand
			if quant_wolfFur == 0 {
				wolfFurAvailable = false // Potion n'est plus disponible après avoir été épuisée
			}
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
			quant_trollSkin-- // Réduit la quantité disponible chez le marchand
			if quant_trollSkin == 0 {
				trollSkinAvailable = false // Potion n'est plus disponible après avoir été épuisée
			}
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
			quant_boarLeather-- // Réduit la quantité disponible chez le marchand
			if quant_boarLeather == 0 {
				boarLeatherAvailable = false // Potion n'est plus disponible après avoir été épuisée
			}
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
			quant_ravenFeather-- // Réduit la quantité disponible chez le marchand
			if quant_ravenFeather == 0 {
				ravenFeatherAvailable = false // Potion n'est plus disponible après avoir été épuisée
			}
		} else if u.purse < ravenFeatherPrice {
			clear()
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cette plume.")
		} else {
			clear()
			fmt.Println("Le marchand n'a plus de plumes de corbeau.")
		}
		u.accessMerchant()

	case "exit", "e":
		clear()
		fmt.Println("Vous quittez le marchand.")
		return

	default:
		clear()
		fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
		u.accessMerchant()
	}
}
