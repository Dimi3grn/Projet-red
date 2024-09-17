package main

import "fmt"

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
		attack: 5,
	}
}

func (u *character) StartCombat() {
	goblin := InitGoblin()

	// Combat loop
	for u.hp > 0 && goblin.hp > 0 {
		// Display the current status
		fmt.Printf("Vos points de vie : %d | Points de vie du %s : %d\n", u.hp, goblin.name, goblin.hp)

		// Player's turn
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Accéder à l'inventaire")
		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			// Attack the goblin
			damage := 5 // Assuming fixed attack damage, you can adjust this if needed
			goblin.hp -= damage
			fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", goblin.name, damage)
			if goblin.hp <= 0 {
				fmt.Println("Vous avez vaincu le gobelin !")
				return
			}

			// Goblin's turn
			u.hp -= goblin.attack
			fmt.Printf("Le %s vous attaque pour %d points de dégâts !\n", goblin.name, goblin.attack)
			if u.hp <= 0 {
				fmt.Println("Vous avez été vaincu par le gobelin...")
				return
			}

		case "2":
			// Access inventory
			u.accessFightInventory() // Assuming this function handles inventory and healing
			// After using the inventory, continue the loop to let the player act again

		default:
			fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
		}
	}
}

// Fonction pour attaquer le monstre
func (u *character) Attack(monstre *Monstre) {
	damage := u.attack // Suppose que `attack` est un attribut de `character`
	monstre.hp -= damage
	if monstre.hp < 0 {
		monstre.hp = 0
	}
	fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", monstre.name, damage)
}

// Fonction pour que le monstre attaque le joueur
func (m *Monstre) MonstreAttack(joueur *character) {
	damage := m.attack
	joueur.hp -= damage // Suppose que `hp` est un attribut de `character`
	if joueur.hp < 0 {
		joueur.hp = 0
	}
	fmt.Printf("Le %s vous attaque pour %d points de dégâts !\n", m.name, damage)
}

// Fonction de boucle de combat
func (u *character) CombatLoop() {
	goblin := InitGoblin()
	for u.hp > 0 && goblin.hp > 0 {
		// Affichage des statistiques
		fmt.Printf("Vos points de vie : %d | Points de vie du gobelin : %d\n", u.hp, goblin.hp)

		// Le joueur attaque
		u.Attack(&goblin)
		if goblin.hp <= 0 {
			fmt.Println("Vous avez vaincu le gobelin !")
			return
		}

		// Le gobelin attaque
		goblin.MonstreAttack(u)
		if u.hp <= 0 {
			fmt.Println("Vous avez été vaincu par le gobelin...")
			return
		}
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
			clear()
			u.spellBook(item.name)
		default:
			fmt.Println("Cet objet ne peut pas être utilisé ou équipé.")
		}
	}
}
