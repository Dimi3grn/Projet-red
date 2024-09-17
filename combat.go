package main

import (
	"fmt"
	"time"
)

// Définition de la structure Monstre
type Monstre struct {
	name   string
	maxHP  int
	hp     int
	attack int
}

// Initialisation d'un gobelin d'entraînement
func InitGoblin() Monstre {
	return Monstre{
		name:   "Gobelin d'entraînement",
		maxHP:  40,
		hp:     40,
		attack: 3,
	}
}

func (g *Monstre) goblinPattern(turn int, u *character) {
	var damage int
	if turn%3 == 0 {
		// On every 3rd turn, goblin deals 200% damage
		damage = g.attack * 2
	} else {
		// Regular attack: 100% damage
		damage = g.attack
	}

	u.hp -= damage
	if u.hp < 0 {
		u.hp = 0 // Prevent negative HP
	}

	// Display attack details
	fmt.Printf("%s inflige à %s %d de dégâts\n", g.name, u.name, damage)
	time.Sleep(1 * time.Second)

}

// StartCombat updated to use the goblinPattern
func (u *character) StartCombat() {
	goblin := InitGoblin()
	turn := 1 // Track combat turns

	// Combat loop
	for u.hp > 0 && goblin.hp > 0 {
		// Display the current status
		fmt.Printf("Vos points de vie : %d | Points de vie du %s : %d\n", u.hp, goblin.name, goblin.hp)
		time.Sleep(1 * time.Second)
		// Player's turn
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Accéder à l'inventaire")
		var choice string
		fmt.Scan(&choice)
		clear()

		switch choice {
		case "1":
			// Attack the goblin
			damage := 5 // Assuming fixed attack damage, you can adjust this if needed
			goblin.hp -= damage
			fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", goblin.name, damage)
			time.Sleep(1 * time.Second)
			if goblin.hp <= 0 {
				fmt.Println("Vous avez vaincu le gobelin !")
				return
			}

		case "2":
			// Access inventory
			u.accessFightInventory() // Assume this handles healing, then continue combat

		default:
			fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
			continue
		}

		// Goblin's turn with attack pattern
		goblin.goblinPattern(turn, u)
		if u.hp <= 0 {
			fmt.Println("Vous avez été vaincu par le gobelin...")
			return
		}

		turn++ // Increment turn count
	}
}

func (u *character) accessFightInventory() {
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

		if choix == "exit" || choix == "e" {
			clear()
			return // Return to combat without resetting goblin
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
			return
		case "Consumable":
			clear()
			u.useConsumable(item)
			return
		case "Book":
			clear()
			u.spellBook(item.name)
			return
		default:
			fmt.Println("Cet objet ne peut pas être utilisé ou équipé.")
			return
		}
	}
}
