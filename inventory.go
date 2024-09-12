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
	clear()
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Printf("╒══════════╡%sVotre inventaire%s╞══════════╕\n", yellow, reset)
	for cpt, v := range u.inv {
		fmt.Printf(" %s.%d%s - %s ⨯ %d\n",
			yellow, cpt+1, reset, v.name, v.amout)
	}
	fmt.Printf("\n vous avez %s%d/%d%s objets dans votre inventaire\n╘══════════════════════════════════════╛\n", yellow, len(u.inv), u.invSize, reset)
	fmt.Printf("Tapez le numéro de l'objet à utiliser ou %s'exit'%s pour sortir.\n", red, reset)
	var choix int
	fmt.Scan(&choix)
	if choix == 0 {
		loop()
		return
	}

	// Vérifier si le choix est valide
	if choix > 0 && choix <= len(u.inv) {
		item := u.inv[choix-1] // Récupérer l'objet choisi
		if item.name == "Livre de Sort: Boule de Feu" {
			u.spellBook("Boule de feu")
			// Retirer l'objet après utilisation
			u.inv = append(u.inv[:choix-1], u.inv[choix:]...)
			fmt.Println("Vous avez appris le sort 'Boule de feu' !")
		} else {
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
		if k.id == 1 {
			isInside = true
			cpt = j
		}
	}
	if isInside {
		if u.hp < u.maxHp {
			u.hp += 5
			if u.hp > u.maxHp {
				u.hp = u.maxHp
			}
			fmt.Printf("Vous avez maintenant %d points de vie\n", u.hp)
			if u.inv[cpt].amout == 1 {
				u.inv = append(u.inv[:cpt], u.inv[cpt+1:]...)
			} else {
				u.inv[cpt].amout -= 1
			}
		} else {
			fmt.Println("vous êtes déjà à votre limite de vie")
		}
	} else {
		fmt.Println("Pas de potions de soin dans l'inventaire")
	}
}
