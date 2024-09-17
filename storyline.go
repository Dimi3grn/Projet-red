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
	}
}

func initAncientDemon() Monstre {
	return Monstre{
		name:       "Démon ancient",
		maxHP:      40,
		hp:         40,
		attack:     3,
		initiative: 3,
	}
}
func initDragon() Monstre {
	return Monstre{
		name:       "Dragon",
		maxHP:      100,
		hp:         100,
		attack:     5,
		initiative: 5,
	}
}

// variables servant a memoriser si le boss est mort
var KnightDefeated bool = true
var AncientDemonDefeated bool = true
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
			fmt.Printf("Le goblin vous a enflamé, vous avez désormais %d points de vie\n", u.hp)
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
	rand.Seed(time.Now().UnixNano())
	red := "\033[31m"
	reset := "\033[0m"
	var damage int
	if rand.Intn(5) == 0 {
		damage = g.attack * 2
		fmt.Println("The dragon hit a critical hit")
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

//fonctions pour chaque boss, dans l'ordre

func (u *character) StartFight1() {
	//ascii art pour le chevalier et le joueur
	fmt.Printf("Vous commencez votre histoire en tant que %s\n", u.classe)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Dès la sortie du chateau, vous croisez un chevalier corrompu vous demandant ...")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Abbatez le pour mettre fin a cette injustice")
	time.Sleep(3 * time.Second)
	yellow := "\033[33m"
	reset := "\033[0m"
	green := "\033[32m"
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
		// Display the current status
		fmt.Printf("Vos points de vie : %s%d%s |%s Points de vie du %s : %s%d\n%s", green, u.hp, yellow, reset, Knight.name, green, Knight.hp, reset)
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
				Knight.hp -= damage
				fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", Knight.name, damage)
				time.Sleep(1 * time.Second)
				if Knight.hp <= 0 {
					fmt.Println("Vous avez vaincu le chevalier !")
					KnightDefeated = true
					return
				}

			case "2":
				// Access inventory
				u.accessFightInventory() // Assume this handles healing, then continue combat
			case "3":
				clear()
				loop()
			default:
				fmt.Println("Choix non valide. Veuillez essayer de nouveau.")
				continue
			}
		} else {
			// Goblin's turn with attack pattern
			Knight.knightPattern(u)
			if u.hp <= 0 {
				fmt.Println("Vous avez été vaincu par le chevalier...")
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
	fmt.Printf("Vous vous retrouvez maintenant face à un Démon ancient%s\n", u.classe)
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
				damage := 5 // Assuming fixed attack damage, you can adjust this if needed
				demon.hp -= damage
				fmt.Printf("Vous attaquez le %s pour %d points de dégâts !\n", demon.name, damage)
				time.Sleep(1 * time.Second)
				if demon.hp <= 0 {
					fmt.Println("Vous avez vaincu le chevalier !")
					AncientDemonDefeated = true
					return
				}

			case "2":
				// Access inventory
				u.accessFightInventory() // Assume this handles healing, then continue combat
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
			if u.hp <= 0 {
				fmt.Println("Vous avez été vaincu par le Démon ancient...")
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
				u.accessFightInventory() // Assume this handles healing, then continue combat
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
			if u.hp <= 0 {
				fmt.Println("Vous avez été vaincu par le chevalier...")
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
