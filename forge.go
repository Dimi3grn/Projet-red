package main

import "fmt"

// Ajout des nouveaux objets à fabriquer
var adventureHat obj = obj{9, "Chapeau de l'aventurier", 1, "EquipHead", 10}
var adventureTunic obj = obj{10, "Tunique de l'aventurier", 1, "EquipChest", 15}
var adventureBoots obj = obj{11, "Bottes de l'aventurier", 1, "EquipBoots", 5}

var defaultHat obj = obj{12, "chapeau par defaut", 1, "EquipHead", 0}
var defaultBody obj = obj{13, "tunique par defaut", 1, "EquipHead", 0}
var defaultLegs obj = obj{14, "pantalon par defaut", 1, "EquipHead", 0}

// Accès au forgeron
func (u *character) accessBlacksmith() {
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"
	green := "\033[32m"

	fmt.Printf("╒══════════╡%sForgeron%s╞══════════╕\n \tPurse : %d\n", yellow, reset, u.purse)

	// Afficher les options de craft avec les ingrédients et le nombre d'objets dans l'inventaire
	fmt.Printf("%s╭%s1.%s Fabriquer un Chapeau de l'aventurier (%s5 $%s)\n", yellow, red, reset, green, reset)
	if u.checkInventory("Plume de Corbeau") >= 1 {
		fmt.Printf("%s│%s  - 1 Plume de Corbeau (%s%d%s)\n", yellow, reset, green, u.checkInventory("Plume de Corbeau"), reset)
	} else {
		fmt.Printf("%s│%s  - 1 Plume de Corbeau (%s%d%s)\n", yellow, reset, red, u.checkInventory("Plume de Corbeau"), reset)
	}
	if u.checkInventory("Cuir de Sanglier") >= 1 {
		fmt.Printf("%s│%s  - 1 Cuir de Sanglier (%s%d%s)\n", yellow, reset, green, u.checkInventory("Cuir de Sanglier"), reset)
	} else {
		fmt.Printf("%s│%s  - 1 Cuir de Sanglier (%s%d%s)\n", yellow, reset, red, u.checkInventory("Cuir de Sanglier"), reset)
	}

	fmt.Printf("%s│%s2.%s Fabriquer une Tunique de l'aventurier (%s5 $%s)\n", yellow, red, reset, green, reset)
	if u.checkInventory("Fourrure de Loup") >= 2 {
		fmt.Printf("%s│%s  - 2 Fourrure de Loup (%s%d%s)\n", yellow, reset, green, u.checkInventory("Fourrure de Loup"), reset)
	} else {
		fmt.Printf("%s│%s  - 2 Fourrure de Loup (%s%d%s)\n", yellow, reset, red, u.checkInventory("Fourrure de Loup"), reset)
	}
	if u.checkInventory("Peau de Troll") >= 1 {
		fmt.Printf("%s│%s  - 1 Peau de Troll (%s%d%s)\n", yellow, reset, green, u.checkInventory("Peau de Troll"), reset)
	} else {
		fmt.Printf("%s│%s  - 1 Peau de Troll (%s%d%s)\n", yellow, reset, red, u.checkInventory("Peau de Troll"), reset)
	}

	fmt.Printf("%s│%s3.%s Fabriquer des Bottes de l'aventurier (%s5 $%s)\n", yellow, red, reset, green, reset)
	if u.checkInventory("Fourrure de Loup") >= 1 {
		fmt.Printf("%s│%s  - 1 Fourrure de Loup (%s%d%s)\n", yellow, reset, green, u.checkInventory("Fourrure de Loup"), reset)
	} else {
		fmt.Printf("%s│%s  - 1 Fourrure de Loup (%s%d%s)\n", yellow, reset, red, u.checkInventory("Fourrure de Loup"), reset)
	}
	if u.checkInventory("Cuir de Sanglier") >= 1 {
		fmt.Printf("%s│%s  - 1 Cuir de Sanglier (%s%d%s)\n", yellow, reset, green, u.checkInventory("Cuir de Sanglier"), reset)
	} else {
		fmt.Printf("%s│%s  - 1 Cuir de Sanglier (%s%d%s)\n", yellow, reset, red, u.checkInventory("Cuir de Sanglier"), reset)
	}

	fmt.Printf("%s╰%s'exit'%s. Retour au menu principal\n", yellow, red, reset)

	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "1":
		u.craftAdventureHat()
	case "2":
		u.craftAdventureTunic()
	case "3":
		u.craftAdventureBoots()
	case "exit", "e":
		clear()
		loop()
	default:
		clear()
		fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
		u.accessBlacksmith()
	}
}

// Fonction pour fabriquer le Chapeau de l'aventurier
func (u *character) craftAdventureHat() {
	clear()
	if u.purse < 5 {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour fabriquer cet équipement.")
	} else if u.checkInventory("Plume de Corbeau") < 1 || u.checkInventory("Cuir de Sanglier") < 1 {
		fmt.Println("Vous n'avez pas assez de matériaux pour fabriquer le Chapeau de l'aventurier.")
	} else {
		u.purse -= 5
		u.removeInventory(ravenFeather)
		u.removeInventory(boarLeather)
		u.addInventory(adventureHat)
		fmt.Println("Vous avez fabriqué un Chapeau de l'aventurier !")
	}
	u.accessBlacksmith()
}

// Fonction pour fabriquer la Tunique de l'aventurier
func (u *character) craftAdventureTunic() {
	clear()
	if u.purse < 5 {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour fabriquer cet équipement.")
	} else if u.checkInventory("Fourrure de Loup") < 2 || u.checkInventory("Peau de Troll") < 1 {
		fmt.Println("Vous n'avez pas assez de matériaux pour fabriquer la Tunique de l'aventurier.")
	} else {
		u.purse -= 5
		u.removeInventory(wolfFur)
		u.removeInventory(wolfFur)
		u.removeInventory(trollSkin)
		u.addInventory(adventureTunic)
		fmt.Println("Vous avez fabriqué une Tunique de l'aventurier !")
	}
	u.accessBlacksmith()
}

// Fonction pour fabriquer les Bottes de l'aventurier
func (u *character) craftAdventureBoots() {
	clear()
	if u.purse < 5 {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour fabriquer cet équipement.")
	} else if u.checkInventory("Fourrure de Loup") < 1 || u.checkInventory("Cuir de Sanglier") < 1 {
		fmt.Println("Vous n'avez pas assez de matériaux pour fabriquer les Bottes de l'aventurier.")
	} else {
		u.purse -= 5
		u.removeInventory(wolfFur)
		u.removeInventory(boarLeather)
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
			count += item.amount
		}
	}
	return count
}

// Fonction pour retirer un certain nombre d'items de l'inventaire
func (u *character) removeItem(itemName string, qty int) {
	fmt.Println(itemName)
	fmt.Println(qty)
	newInventory := []obj{}
	for _, item := range u.inv {
		if item.name == itemName && qty > 0 {
			fmt.Println("entered")
			qty--
			continue
		}
		newInventory = append(newInventory, item)
	}
	u.inv = newInventory
}
