package main

import "fmt"

type obj struct {
	id    int
	name  string
	amout int
	cath  string
}

var obj1 obj = obj{1, "Health Pot", 1, "Consumable"}
var obj2 obj = obj{2, "Poison Pot", 1, "Consumable"}

//var obj3 obj = obj{3, "Sword", 1, "Equipement"}
var fireSpellBook obj = obj{4, "Livre de Sort: Boule de Feu", 1, "Book"}

func (u *character) addInventory(item obj) {
	if len(u.inv) < u.invSize {
		for i, invItem := range u.inv {
			if invItem.id == item.id {
				u.inv[i].amout += item.amout
				return
			}
		}
		item.amout = 1
		u.inv = append(u.inv, item)
	} else {
		fmt.Println("Votre inventaire est plein !")
	}
}

func (u *character) removeInventory(item obj) {
	for i, invItem := range u.inv {
		if invItem.id == item.id {
			if invItem.amout > 1 {
				u.inv[i].amout -= 1
			} else {
				u.inv = append(u.inv[:i], u.inv[i+1:]...)
			}
			return
		}
	}
}

func (u *character) accessInventory() {
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Printf("╒══════════╡%sVotre inventaire%s╞══════════╕\n", yellow, reset)
	for cpt, v := range u.inv {
		fmt.Printf(" %s.%d%s - %s ⨯ %d\n",
			yellow, cpt+1, reset, v.name, v.amout)
	}
	fmt.Printf("\n vous avez %s%d/%d%s objets dans votre inventaire\n╘══════════════════════════════════════╛\n", yellow, len(u.inv), u.invSize, reset)
	fmt.Printf("Tapez le numéro de l'objet à utiliser.\n")
	fmt.Printf("%s╭%s'hp'%s pour récuperer hp à partir des Heatlh Pot\n%s╰%s'exit'%s pour quitter l'Inventaire\n", yellow, red, reset, yellow, red, reset)
	var choix int
	fmt.Scan(&choix)
	if choix == 0 {
		clear()
		loop()
		return
	}
	// Si choix == 0, quitter l'inventaire
	if choix == 0 {
		clear()
		loop()
		return
	}

	// Vérifier si le choix est valide
	if choix > 0 && choix <= len(u.inv) {
		item := u.inv[choix-1] // Récupérer l'objet choisi
		switch item.name {
		case "Livre de Sort: Boule de Feu":
			u.spellBook("Boule de feu")
			// Retirer l'objet après utilisation
			u.removeInventory(item)
			fmt.Println("Vous avez appris le sort 'Boule de feu' !")

		case "Health Pot":
			u.takePot() // Utilise une potion de soin

		case "Poison Pot":
			u.takePoisonPot() // Utilise une potion de poison

		default:
			fmt.Println("Cet objet ne peut pas être utilisé.")
		}
	} else {
		fmt.Println("Choix invalide.")
	}

	u.accessInventory() // Réafficher l'inventaire après interaction
}

func (u *character) takePot() {
	isInside := false
	cpt := 0
	for j, k := range u.inv {
		if k.id == 1 { // ID pour Health Pot
			isInside = true
			cpt = j
		}
	}
	if isInside {
		if u.hp < u.maxHp {
			u.hp += 5 // Guérit 5 points de vie
			if u.hp > u.maxHp {
				u.hp = u.maxHp
			}
			fmt.Printf("Vous avez maintenant %d points de vie\n", u.hp)
			u.removeInventory(u.inv[cpt]) // Retirer la potion de l'inventaire
		} else {
			fmt.Println("Vous êtes déjà à votre limite de vie.")
		}
	} else {
		fmt.Println("Pas de potions de soin dans l'inventaire.")
	}
}

func (u *character) takePoisonPot() {
	isInside := false
	cpt := 0
	for j, k := range u.inv {
		if k.id == 2 { // ID pour Poison Pot
			isInside = true
			cpt = j
		}
	}
	if isInside {
		u.hp -= 5 // Inflige 5 points de dégât
		if u.hp < 0 {
			u.hp = 0
		}
		fmt.Printf("Vous avez maintenant %d points de vie (après avoir utilisé une potion de poison)\n", u.hp)
		u.removeInventory(u.inv[cpt]) // Retirer la potion de poison de l'inventaire
	} else {
		fmt.Println("Pas de potions de poison dans l'inventaire.")
	}
}
