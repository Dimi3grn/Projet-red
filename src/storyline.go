package main

import (
	"fmt"
	"math/rand"
	"time"
)

// initialisation mobs
func initKnight() Monstre {
	return Monstre{
		name:       "Chevalier apprentit",
		maxHP:      25,
		hp:         25,
		attack:     2,
		initiative: 1,
		exp:        300,
	}
}

func initAncientDemon() Monstre {
	return Monstre{
		name:       "Démon ancient",
		maxHP:      40,
		hp:         40,
		attack:     3,
		initiative: 3,
		exp:        800,
	}
}
func initDragon() Monstre {
	return Monstre{
		name:       "Dragon",
		maxHP:      100,
		hp:         100,
		attack:     5,
		initiative: 5,
		exp:        2000,
	}
}

// variables servant a memoriser si le boss est mort
var KnightDefeated bool = true
var AncientDemonDefeated bool = false
var DragonDeafeated bool = false

// fonctions des patterns d'attaque des monstres
func (g *Monstre) knightPattern(u *character) {
	rand.Seed(time.Now().UnixNano())
	red := "\033[31m"
	reset := "\033[0m"
	var damage int
	if rand.Intn(5) == 0 {
		damage = g.attack * 2
		fmt.Println("The Knight hit a critical hit")
	} else {
		// Regular attack: 100% damage
		damage = g.attack
	}

	u.hp -= damage
	if u.hp < 0 {
		u.hp = 0 // Prevent negative HP
	}

	// Display attack details
	fmt.Printf("%s vous inflige à %s%d%s points de dégâts\n", g.name, red, damage, reset)
	time.Sleep(1 * time.Second)

}

func (g *Monstre) ancientDemonPattern(u *character) {
	var isonfire bool = false
	rand.Seed(time.Now().UnixNano())
	red := "\033[31m"
	reset := "\033[0m"
	var damage int
	if rand.Intn(100) < 60 {
		damage = g.attack
		isonfire = true
		fmt.Println("Le Démon brule votre âme")
	} else {
		// Regular attack: 100% damage
		damage = g.attack
	}
	u.hp -= damage
	if u.hp < 0 {
		u.hp = 0 // Prevent negative HP
	}
	fmt.Printf("%s vous inflige à %s%d%s points de dégâts\n", g.name, red, damage, reset)
	time.Sleep(1 * time.Second)
	if isonfire {
		for cpt := 0; cpt < 3; cpt++ {
			u.hp -= 1
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Le goblin vous a enflamé, vous avez désormais %s%d%s points de vie\n", red, u.hp, reset)
			if u.hp < 0 {
				u.hp = 0 // Prevent negative HP
			}
		}
	}
	isonfire = false
	// Display attack details
	time.Sleep(1 * time.Second)
}

func (g *Monstre) dragonPattern(u *character) { //fonction basic, pas d'attaque spéciale, a compléter
	var hasprinted bool = false
	rand.Seed(time.Now().UnixNano())
	red := "\033[31m"
	reset := "\033[0m"
	var damage int
	damage = g.attack

	if g.hp < 40 {

		if hasprinted == false {
			fmt.Println("Le dragon rentre dans une nouvelle phase, ses dégats sont augmentés")

		}
		damage += 2
	}

	u.hp -= damage
	if u.hp < 0 {
		u.hp = 0 // Prevent negative HP
	}

	// Display attack details
	fmt.Printf("%s vous inflige à %s%d%s points de dégâts\n", g.name, red, damage, reset)
	time.Sleep(1 * time.Second)

}

//fonctions pour chaque boss, dans l'ordre

func (u *character) StartFight1() {
	//ascii art pour le chevalier et le joueur
	fmt.Printf("Vous commencez votre histoire en tant que %s\n", u.classe)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Dès la sortie du chateau, vous croisez un chevalier corrompu vous demandant de l'argent")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Abattez le pour mettre fin a cette injustice")
	time.Sleep(3 * time.Second)
	yellow := "\033[33m"
	reset := "\033[0m"
	Knight := initKnight()
	turn := 1 // Track combat turns

	// Determine who starts based on initiative
	playerTurn := u.initiative >= Knight.initiative
	if !playerTurn {
		fmt.Println("le chevalier vous attaque car il a plus d'initiative que vous")
		time.Sleep(1 * time.Second)
	} else {
		fmt.Println("Votre niveau d'initiative vous permet d'attaquer le chevalier en premier")
		time.Sleep(1 * time.Second)
	}
	// Combat loop
	for u.hp > 0 && Knight.hp > 0 {
		if playerTurn {
			mainchar()
			fmt.Print("\t")
			health_bar(u.hp, u.maxHp)
			fmt.Print("\t\t\t\t\t\t")
			health_bar(Knight.hp, Knight.maxHP)
			fmt.Print("\n")
		}

		// Display the current status
		time.Sleep(1 * time.Second)

		if playerTurn {
			// Player's turn
			fmt.Println("Que voulez-vous faire ?")
			fmt.Printf("%s1.%s Attaquer\n", yellow, reset)
			fmt.Printf("%s2.%s Accéder à l'inventaire\n", yellow, reset)
			fmt.Printf("%s3.%s jeter un sort\n", yellow, reset)
			fmt.Printf("%s4.%s Se replier\n", yellow, reset)
			var choice string
			fmt.Scan(&choice)
			clear()

			switch choice {
			case "1":
				// Attack the goblin
				damage := u.attack // Assuming fixed attack damage, you can adjust this if needed
				Knight.hp -= damage
				fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", Knight.name, damage)
				time.Sleep(1 * time.Second)
				if Knight.hp <= 0 {
					fmt.Println("Vous avez vaincu le chevalier !")
					u.exp += Knight.exp
					u.updateXp()
					KnightDefeated = true
					return
				}

			case "2":
				// Access inventory
				u.accessFightInventory(&Knight) // Assume this handles healing, then continue combat
			case "3":
				clear()
				for cpt, v := range u.skills {
					fmt.Printf(" %s.%d%s - %s\n", yellow, cpt+1, reset, v)
				}

				var choix string
				fmt.Scan(&choix)
				var choixInt int
				_, err := fmt.Sscanf(choix, "%d", &choixInt)
				if err != nil || choixInt <= 0 || choixInt > len(u.skills) {
					clear()
					fmt.Println("Choix invalide, veuillez réessayer.")
					continue
				}
				item := u.skills[choixInt-1]
				switch item {
				case "Coup de poing":
					damage := 8 // Assuming fixed attack damage, you can adjust this if needed
					Knight.hp -= damage
					fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", Knight.name, damage)
					time.Sleep(1 * time.Second)
					if Knight.hp <= 0 {
						fmt.Println("Vous avez vaincu le chevalier !")
						u.exp += Knight.exp
						u.updateXp()
						KnightDefeated = true
						return
					}
				case "Livre de Sort: Boule de Feu":
					damage := 18 // Assuming fixed attack damage, you can adjust this if needed
					Knight.hp -= damage
					fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", Knight.name, damage)
					time.Sleep(1 * time.Second)
					if Knight.hp <= 0 {
						fmt.Println("Vous avez vaincu le chevalier !")
						fmt.Println("Vous pouvez relancer un Combat pour continuer l'histoire")
						u.exp += Knight.exp
						u.updateXp()
						KnightDefeated = true
						return
					}
				}

			case "4":
				clear()
				loop()
			default:
				fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
				continue
			}
		} else {
			// Goblin's turn with attack pattern
			Knight.knightPattern(u)
			var toreturn = false
			if u.dead() {
				fmt.Println("Vous avez été vaincu par le chevalier...")
				toreturn = true
			}
			if toreturn {
				u.dead()
				return
			}
		}

		// Switch turn for the next round
		playerTurn = !playerTurn
		turn++ // Increment turn count
	}
}

func (u *character) StartFight2() {
	//ascii art pour le chevalier et le joueur
	fmt.Printf("Vous vous retrouvez maintenant face à un Démon ancient\n")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("vous sentez une présence obscure vous envahir")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Surmontez vos peurs en l'abattant")
	time.Sleep(3 * time.Second)
	yellow := "\033[33m"
	reset := "\033[0m"
	green := "\033[32m"
	demon := initAncientDemon()
	turn := 1 // Track combat turns

	// Determine who starts based on initiative
	playerTurn := u.initiative >= demon.initiative
	if !playerTurn {
		fmt.Println("Le Démon Ancien vous attaque car il a plus d'initiative que vous")
		time.Sleep(1 * time.Second)
	} else {
		fmt.Println("Votre niveau d'initiative vous permet d'attaquer le chevalier en premier")
		time.Sleep(1 * time.Second)
	}
	// Combat loop
	for u.hp > 0 && demon.hp > 0 {
		if playerTurn {
			printdem()
			fmt.Print("\t")
			health_bar(u.hp, u.maxHp)
			fmt.Print("\t\t\t\t\t\t\t\t")
			health_bar(demon.hp, demon.maxHP)
			fmt.Print("\n")
		}
		// Display the current status
		fmt.Printf("Vos points de vie : %s%d%s |%s Points de vie du %s : %s%d\n%s", green, u.hp, yellow, reset, demon.name, green, demon.hp, reset)
		time.Sleep(1 * time.Second)

		if playerTurn {
			// Player's turn
			fmt.Println("Que voulez-vous faire ?")
			fmt.Printf("%s1.%s Attaquer\n", yellow, reset)
			fmt.Printf("%s2.%s Accéder à l'inventaire\n", yellow, reset)
			fmt.Printf("%s3.%s Se replier\n", yellow, reset)
			var choice string
			fmt.Scan(&choice)
			clear()

			switch choice {
			case "1":
				// Attack the goblin
				damage := u.attack // Assuming fixed attack damage, you can adjust this if needed
				demon.hp -= damage
				fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", demon.name, damage)
				time.Sleep(1 * time.Second)
				if demon.hp <= 0 {
					fmt.Println("Vous avez vaincu le chevalier !")
					u.exp += demon.exp
					u.updateXp()
					AncientDemonDefeated = true
					return
				}

			case "2":
				// Access inventory
				u.accessFightInventory(&demon) // Assume this handles healing, then continue combat
			case "3":
				clear()
				loop()
			default:
				fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
				continue
			}
		} else {
			// Goblin's turn with attack pattern
			demon.ancientDemonPattern(u)
			var toreturn = false
			if u.dead() {
				fmt.Println("Vous avez été vaincu par le chevalier...")
				toreturn = true
			}
			if toreturn {
				u.dead()
				return
			}
		}

		// Switch turn for the next round
		playerTurn = !playerTurn
		turn++ // Increment turn count
	}
}

func (u *character) StartFight3() {
	//ascii art pour le chevalier et le joueur
	fmt.Println("Vous tombez nez à nez avec un dragon millénaire")
	time.Sleep(1 * time.Second)
	fmt.Println("Son souffle incandescent fait vibrer l'air, et ses écailles sombres absorbent toute lumière.")
	time.Sleep(1 * time.Second)
	fmt.Println("Ses yeux ardents vous fixent, témoins d'une colère ancestrale.")
	time.Sleep(1 * time.Second)
	fmt.Println("Chaque battement de ses ailes soulève des rafales, tandis qu'un grondement résonne dans sa gorge.")
	time.Sleep(3 * time.Second)
	fmt.Println("Le Dragon de la Désolation vous fixe.")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Il n’y a ni échappatoire, ni pitié à attendre. Le combat pour votre survie commence.")
	time.Sleep(3 * time.Second)
	yellow := "\033[33m"
	reset := "\033[0m"
	green := "\033[32m"
	dragon := initDragon()
	turn := 1 // Track combat turns

	// Determine who starts based on initiative
	playerTurn := u.initiative >= dragon.initiative
	if !playerTurn {
		fmt.Println("le dragon vous attaque car il a plus d'initiative que vous")
		time.Sleep(1 * time.Second)
	} else {
		fmt.Println("Votre niveau d'initiative vous permet d'attaquer le dragon en premier")
		time.Sleep(1 * time.Second)
	}
	// Combat loop
	for u.hp > 0 && dragon.hp > 0 {
		if playerTurn {
			printdrag()
			fmt.Print("\t")
			health_bar(u.hp, u.maxHp)
			fmt.Print("\t\t\t\t\t\t\t\t")
			health_bar(dragon.hp, dragon.maxHP)
			fmt.Print("\n")
		}

		// Display the current status
		fmt.Printf("Vos points de vie : %s%d%s |%s Points de vie du %s : %s%d\n%s", green, u.hp, yellow, reset, dragon.name, green, dragon.hp, reset)
		time.Sleep(1 * time.Second)

		if playerTurn {
			// Player's turn
			fmt.Println("Que voulez-vous faire ?")
			fmt.Printf("%s1.%s Attaquer\n", yellow, reset)
			fmt.Printf("%s2.%s Accéder à l'inventaire\n", yellow, reset)
			fmt.Printf("%s3.%s Se replier\n", yellow, reset)
			var choice string
			fmt.Scan(&choice)
			clear()

			switch choice {
			case "1":
				// Attack the goblin
				damage := 5 // Assuming fixed attack damage, you can adjust this if needed
				dragon.hp -= damage
				fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", dragon.name, damage)
				time.Sleep(1 * time.Second)
				if dragon.hp <= 0 {
					fmt.Println("Vous avez vaincu le chevalier !")
					DragonDeafeated = true
					return
				}

			case "2":
				// Access inventory
				u.accessFightInventory(&dragon) // Assume this handles healing, then continue combat
			case "3":
				clear()
				loop()
			default:
				fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
				continue
			}
		} else {
			// Goblin's turn with attack pattern
			dragon.dragonPattern(u)
			var toreturn = false
			if u.dead() {
				fmt.Println("Vous avez été vaincu par le chevalier...")
				toreturn = true
			}
			if toreturn {
				u.dead()
				return
			}
		}

		// Switch turn for the next round
		playerTurn = !playerTurn
		turn++ // Increment turn count
	}
}

// sends you to the right fight
func (u *character) getCurrentFight() {
	if KnightDefeated == false {
		u.StartFight1()
	} else if AncientDemonDefeated == false {
		u.StartFight2()
	} else if DragonDeafeated == false {
		u.StartFight3()
	}
}
