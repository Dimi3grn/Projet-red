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
var obj3 obj = obj{3, "Sword", 1, "Equipement"}

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
	fmt.Printf("╒══════════╡Votre inventaire╞══════════╕\n")
	for cpt, v := range u.inv {
		fmt.Printf(" .%d - %s ⨯ %d\n",
			cpt+1, v.name, v.amout)
	}
	fmt.Printf("\n vous avez %d/%d objets dans votre inventaire\n╘══════════════════════════════════════╛\n", len(u.inv), u.invSize)
	fmt.Println("Tapez le numéro de l'objet à utiliser ou 'exit' pour sortir.")
	var choix int
	fmt.Scan(&choix)
	if choix == 3 {
		for i, item := range u.inv {
			if item.name == "Livre de Sort: Boule de Feu" {
				u.spellBook("Boule de feu")
				u.inv = append(u.inv[:i], u.inv[i+1:]...)
				break
			}
		}
	}
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
