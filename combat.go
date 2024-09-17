package main

import (
	"fmt"
	"time"
)

// Définition de la structure Boss
type Boss struct {
	name       string
	maxHP      int
	hp         int
	attack     int
	initiative int
}

// Initialisation des différents boss
func InitBosses() []Boss {
	return []Boss{
		{name: "Sword Warrior", maxHP: 50, hp: 50, attack: 7, initiative: 4},
		{name: "Orc", maxHP: 80, hp: 80, attack: 10, initiative: 5},
		{name: "Dragon", maxHP: 120, hp: 120, attack: 15, initiative: 6},
	}
}

// Définition de la structure Monstre
type Monstre struct {
	name       string
	maxHP      int
	hp         int
	attack     int
	initiative int
}

// Initialisation d'un gobelin d'entraînement
func InitGoblin() Monstre {
	return Monstre{
		name:       "Gobelin d'entraînement",
		maxHP:      40,
		hp:         40,
		attack:     3,
		initiative: 3,
	}
}

// Fonction pour gérer le pattern de combat du gobelin
func (g *Monstre) goblinPattern(turn int, u *character) {
	red := "\033[31m"
	reset := "\033[0m"
	var damage int

	if turn%3 == 0 {
		damage = g.attack * 2
	} else {
		damage = g.attack
	}

	u.hp -= damage
	if u.hp < 0 {
		u.hp = 0 // Empêcher les HP négatifs
	}

	// Détails de l'attaque
	fmt.Printf("%s vous inflige %s%d%s points de dégâts\n", g.name, red, damage, reset)
	time.Sleep(1 * time.Second)
}

func (b *Boss) bossPattern(turn int, u *character) {
	red := "\033[31m"
	reset := "\033[0m"
	var damage int

	if turn%3 == 0 {
		damage = b.attack * 2
	} else {
		damage = b.attack
	}

	u.hp -= damage
	if u.hp < 0 {
		u.hp = 0 // Empêcher les HP négatifs
	}

	// Détails de l'attaque
	fmt.Printf("%s inflige %s%d%s points de dégâts à %s\n", b.name, red, damage, reset, u.name)
	time.Sleep(1 * time.Second)
}

// Fonction pour démarrer le combat contre une série de boss
func (u *character) StartBossSequence() {
	bosses := InitBosses() // Charger la liste des boss
	turn := 1              // Suivi des tours de combat

	// Boucle de progression à travers les boss
	for _, boss := range bosses {
		fmt.Printf("Un %s apparaît avec %d points de vie et attaque à %d !\n", boss.name, boss.hp, boss.attack)
		time.Sleep(1 * time.Second)

		// Détermine qui commence en fonction de l'initiative
		playerTurn := u.initiative >= boss.initiative

		// Boucle de combat
		for u.hp > 0 && boss.hp > 0 {
			fmt.Printf("Vos points de vie : %d | Points de vie du %s : %d\n", u.hp, boss.name, boss.hp)
			time.Sleep(1 * time.Second)

			if playerTurn {
				// Tour du joueur
				fmt.Println("Que voulez-vous faire ?")
				fmt.Println("1. Attaquer")
				fmt.Println("2. Accéder à l'inventaire")
				var choice string
				fmt.Scan(&choice)
				clear()

				switch choice {
				case "1":
					// Attaquer le boss
					damage := 5 // Supposons des dégâts fixes
					boss.hp -= damage
					fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", boss.name, damage)
					time.Sleep(1 * time.Second)
					if boss.hp <= 0 {
						fmt.Printf("Vous avez vaincu le %s !\n", boss.name)
						break
					}
				case "2":
					// Accéder à l'inventaire
					u.accessFightInventory()
					// Continue le combat après la gestion de l'inventaire
					continue
				default:
					fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
					continue
				}
			} else {
				// Tour du boss avec son pattern
				boss.bossPattern(turn, u)
				if u.hp <= 0 {
					fmt.Println("Vous avez été vaincu par le boss...")
					return
				}
			}

			// Changement de tour pour le prochain round
			playerTurn = !playerTurn
			turn++
		}

		// Vérification si le joueur est toujours en vie pour combattre le boss suivant
		if u.hp <= 0 {
			fmt.Println("Vous avez perdu. Recommencez depuis le début.")
			return
		}
	}

	fmt.Println("Félicitations ! Vous avez vaincu tous les boss !")
}

// Fonction pour démarrer le combat contre un gobelin d'entraînement
func (u *character) StartCombat() {
	goblin := InitGoblin()
	turn := 1 // Suivi des tours de combat

	// Détermine qui commence en fonction de l'initiative
	playerTurn := u.initiative >= goblin.initiative

	// Combat loop
	for u.hp > 0 && goblin.hp > 0 {
		fmt.Printf("Vos points de vie : %d | Points de vie du %s : %d\n", u.hp, goblin.name, goblin.hp)
		time.Sleep(1 * time.Second)

		if playerTurn {
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
				damage := 5 // Assuming fixed attack damage
				goblin.hp -= damage
				fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", goblin.name, damage)
				time.Sleep(1 * time.Second)
				if goblin.hp <= 0 {
					fmt.Println("Vous avez vaincu le gobelin !")
					return
				}

			case "2":
				// Access inventory
				u.accessFightInventory()
				// Continue combat after handling inventory
				continue

			default:
				fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
				continue
			}
		} else {
			// Goblin's turn with attack pattern
			goblin.goblinPattern(turn, u)
			if u.hp <= 0 {
				fmt.Println("Vous avez été vaincu par le gobelin...")
				return
			}
		}

		// Switch turn for the next round
		playerTurn = !playerTurn
		turn++ // Increment turn count
	}
}

// Fonction pour accéder à l'inventaire pendant le combat
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
