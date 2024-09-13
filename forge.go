package main

import "fmt"

// Ajout des nouveaux objets à fabriquer
var adventureHat obj = obj{9, "Chapeau de l'aventurier", 1, "Equipement"}
var adventureTunic obj = obj{10, "Tunique de l'aventurier", 1, "Equipement"}
var adventureBoots obj = obj{11, "Bottes de l'aventurier", 1, "Equipement"}

// Menu principal avec l'ajout du Forgeron
func (u *character) mainMenu() {
	fmt.Println("Bienvenue dans le menu principal")
	fmt.Println("1. Accéder au marchand")
	fmt.Println("2. Accéder au forgeron")
	fmt.Println("3. Quitter")

	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "1":
		clear()
		u.accessMerchant()
	case "2":
		clear()
		u.accessBlacksmith() // Ajout du forgeron
	case "3":
		clear()
		fmt.Println("Vous quittez le jeu.")
		return
	default:
		clear()
		fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
		u.mainMenu()
	}
}

// Accès au forgeron
func (u *character) accessBlacksmith() {
	//red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"

	fmt.Printf("╒══════════╡%sForgeron%s╞══════════╕\n \tPurse : %d\n", yellow, reset, u.purse)
	fmt.Println("1. Fabriquer un Chapeau de l'aventurier (5 pièces d'or)")
	fmt.Println("2. Fabriquer une Tunique de l'aventurier (5 pièces d'or)")
	fmt.Println("3. Fabriquer des Bottes de l'aventurier (5 pièces d'or)")
	fmt.Println("4. Retour au menu principal")

	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "1":
		u.craftAdventureHat()
	case "2":
		u.craftAdventureTunic()
	case "3":
		u.craftAdventureBoots()
	case "4":
		clear()
		u.mainMenu()
	default:
		clear()
		fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
		u.accessBlacksmith()
	}
}

// Fonction pour fabriquer le Chapeau de l'aventurier
func (u *character) craftAdventureHat() {
	if u.purse < 5 {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour fabriquer cet équipement.")
	} else if u.checkInventory("Plume de Corbeau") < 1 || u.checkInventory("Cuir de Sanglier") < 1 {
		fmt.Println("Vous n'avez pas assez de matériaux pour fabriquer le Chapeau de l'aventurier.")
	} else {
		u.purse -= 5
		u.removeItem("Plume de Corbeau", 1)
		u.removeItem("Cuir de Sanglier", 1)
		u.addInventory(adventureHat)
		fmt.Println("Vous avez fabriqué un Chapeau de l'aventurier !")
	}
	u.accessBlacksmith()
}

// Fonction pour fabriquer la Tunique de l'aventurier
func (u *character) craftAdventureTunic() {
	if u.purse < 5 {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour fabriquer cet équipement.")
	} else if u.checkInventory("Fourrure de Loup") < 2 || u.checkInventory("Peau de Troll") < 1 {
		fmt.Println("Vous n'avez pas assez de matériaux pour fabriquer la Tunique de l'aventurier.")
	} else {
		u.purse -= 5
		u.removeItem("Fourrure de Loup", 2)
		u.removeItem("Peau de Troll", 1)
		u.addInventory(adventureTunic)
		fmt.Println("Vous avez fabriqué une Tunique de l'aventurier !")
	}
	u.accessBlacksmith()
}

// Fonction pour fabriquer les Bottes de l'aventurier
func (u *character) craftAdventureBoots() {
	if u.purse < 5 {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour fabriquer cet équipement.")
	} else if u.checkInventory("Fourrure de Loup") < 1 || u.checkInventory("Cuir de Sanglier") < 1 {
		fmt.Println("Vous n'avez pas assez de matériaux pour fabriquer les Bottes de l'aventurier.")
	} else {
		u.purse -= 5
		u.removeItem("Fourrure de Loup", 1)
		u.removeItem("Cuir de Sanglier", 1)
		u.addInventory(adventureBoots)
		fmt.Println("Vous avez fabriqué des Bottes de l'aventurier !")
	}
	u.accessBlacksmith()
}

// Fonction pour vérifier le nombre d'un item dans l'inventaire
func (u *character) checkInventory(itemName string) int {
	count := 0
	for _, item := range u.inv {
		if item.name == itemName {
			count++
		}
	}
	return count
}

// Fonction pour retirer un certain nombre d'items de l'inventaire
func (u *character) removeItem(itemName string, qty int) {
	newInventory := []obj{}
	for _, item := range u.inv {
		if item.name == itemName && qty > 0 {
			qty--
			continue
		}
		newInventory = append(newInventory, item)
	}
	u.inv = newInventory
}
