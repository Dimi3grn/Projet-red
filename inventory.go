package main

import (
	"fmt"
	"time"
)

type obj struct {
	id     int
	name   string
	amount int
	cath   string
	buff   int
}

var obj1 obj = obj{1, "Health Pot", 1, "Consumable", 0}
var obj2 obj = obj{2, "Poison Pot", 1, "Consumable", 0}

// var obj3 obj = obj{3, "Sword", 1, "Equipement"}
var fireSpellBook obj = obj{4, "Livre de Sort: Boule de Feu", 1, "Book", 0}

func (u *character) addInventory(item obj) {
	if len(u.inv) < u.invSize {
		for i, invItem := range u.inv {
			if invItem.id == item.id {
				u.inv[i].amount += item.amount
				return
			}
		}
		item.amount = 1
		u.inv = append(u.inv, item)
	} else {
		fmt.Println("Votre inventaire est plein !")
	}
}

func (u *character) removeInventory(item obj) {
	for i, invItem := range u.inv {
		if invItem.id == item.id {
			if invItem.amount > 1 {
				u.inv[i].amount -= 1
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
	for {
		fmt.Printf("╒══════════╡%sVotre inventaire%s╞══════════╕\n", yellow, reset)
		for cpt, v := range u.inv {
			fmt.Printf(" %s.%d%s - %s ⨯ %d\n", yellow, cpt+1, reset, v.name, v.amount)
		}
		fmt.Printf("\n vous avez %s%d/%d%s objets dans votre inventaire\n╘══════════════════════════════════════╛\n", yellow, len(u.inv), u.invSize, reset)
		fmt.Printf("Tapez le numéro de l'objet à utiliser.\n")
		fmt.Printf("%s╭%s'[slot number]'%s pour utiliser/equipper l'objet\n%s╰%s'exit'%s pour quitter l'Inventaire\n", yellow, red, reset, yellow, red, reset)

		var choix string
		fmt.Scan(&choix)

		if choix == "exit" {
			clear()
			loop() // Retourne au menu principal
			return
		}

		var choixInt int
		_, err := fmt.Sscanf(choix, "%d", &choixInt)
		if err != nil || choixInt <= 0 || choixInt > len(u.inv) {
			clear()
			fmt.Println("Choix invalide, veuillez réessayer.")
			continue
		}

		item := u.inv[choixInt-1]

		// Check if item is equipable or usable
		switch item.cath {
		case "EquipHead", "EquipChest", "EquipBoots":
			clear()
			u.equipItem(item)
		case "Consumable":
			clear()
			u.useConsumable(item)
		case "Book":
			u.spellBook(item.name)
		default:
			fmt.Println("Cet objet ne peut pas être utilisé ou équipé.")
		}
	}
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
		fmt.Println("Vous avez bu une potion de poison !")

		// Inflige 3 dégâts par seconde pendant 3 secondes (pour un total de 9 dégâts)
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second) // Délai de 1 seconde entre chaque tick
			u.hp -= 3                   // Inflige 3 points de dégâts

			// Vérifier si le joueur est mort
			if u.hp <= 0 {
				u.hp = 0
				fmt.Println("Vous êtes mort suite aux effets du poison.")
				u.dead() // Appelle la fonction "dead" si le joueur est mort
				return
			}

			fmt.Printf("Vous avez %d points de vie après %d secondes de poison.\n", u.hp, i+1)
		}

		fmt.Printf("Le poison s'est dissipé. Vous avez %d points de vie.\n", u.hp)

		// Retirer la potion de poison de l'inventaire après utilisation
		u.removeInventory(u.inv[cpt])
	} else {
		fmt.Println("Pas de potions de poison dans l'inventaire.")
	}
}
func (u *character) equipItem(item obj) {

	// Check which body part the item belongs to
	switch item.cath {
	case "EquipHead":
		u.equipHead(item)
	case "EquipChest":
		u.equipBody(item)
	case "EquipBoots":
		u.equipLegs(item)
	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
	}
}

func (u *character) equipHead(item obj) {
	if u.stuff.head.name != "" {
		u.maxHp -= u.stuff.head.buff
		u.maxHp += item.buff
		u.addInventory(u.stuff.head)
		fmt.Printf("Vous avez retiré %s et équipé %s.\n", u.stuff.head.name, item.name)
		u.removeInventory(item)
	} else {
		u.maxHp += item.buff
		fmt.Printf("Vous avez équipé %s.\n", item.name)
		u.removeInventory(item)
	}
	u.stuff.head = item

}

func (u *character) equipBody(item obj) {
	if u.stuff.body.name != "" {
		fmt.Printf("Vous avez retiré %s et équipé %s.\n", u.stuff.body.name, item.name)
	} else {
		fmt.Printf("Vous avez équipé %s.\n", item.name)
	}
	u.stuff.body = item
	u.maxHp += 25
}

func (u *character) equipLegs(item obj) {
	if u.stuff.legs.name != "" {
		fmt.Printf("Vous avez retiré %s et équipé %s.\n", u.stuff.legs.name, item.name)
	} else {
		fmt.Printf("Vous avez équipé %s.\n", item.name)
	}
	u.stuff.legs = item
	u.maxHp += 20
}

func (u *character) useConsumable(item obj) {
	switch item.name {
	case "Health Pot":
		u.takePot()
	case "Poison Pot":
		u.takePoisonPot()
	default:
		fmt.Println("Cet objet ne peut pas être utilisé.")
	}
}
