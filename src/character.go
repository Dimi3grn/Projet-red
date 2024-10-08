package main

import (
	"fmt"
	"math/rand"
)

type equipement struct {
	head obj
	body obj
	legs obj
}

type character struct {
	name       string
	classe     string
	lvl        int
	exp        int
	maxHp      int
	hp         int
	inv        []obj
	invSize    int
	alive      bool
	purse      int
	skills     []string
	stuff      equipement
	attack     int
	initiative int
	maxExp     int
}

var MyChar character = character{"jack", "humain", 1, 0, 0, 0, []obj{}, 0, true, 200, nil, equipement{}, 7, 1, 100}

func setclasse() {
	rng := rand.Intn(3)
	if rng == 1 {
		MyChar.classe = "humain"
	} else if rng == 2 {
		MyChar.classe = "elfe"
	} else {
		MyChar.classe = "nain"
	}
	if MyChar.classe == "humain" {
		MyChar.maxHp = MyChar.lvl * 20
		MyChar.invSize = 10
	} else if MyChar.classe == "elfe" {
		MyChar.maxHp = MyChar.lvl * 15
		MyChar.invSize = 7
		for i := 0; i < 5; i++ {
			MyChar.addInventory(obj1)
		}
	} else if MyChar.classe == "nain" {
		MyChar.maxHp = MyChar.lvl * 30
		MyChar.invSize = 5
	}
	MyChar.hp = MyChar.maxHp
	MyChar.skills = append(MyChar.skills, "Coup de poing")
}

func (u character) displayinfo() {
	clear()
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Printf("╒════╡%sVos stats%s╞════╕\n name : %s\n classe : %s\n level : %d\n experience : %d/%d\n hp : %d/%d\n Purse : %d\n",
		yellow, reset, u.name, u.classe, u.lvl, u.exp, u.maxExp, u.hp, u.maxHp, u.purse)
	fmt.Print(" Skills :\n    ")
	if len(u.skills) > 0 {
		for i, skill := range u.skills {
			if i > 0 {
				fmt.Print("\n    ")
			}
			fmt.Print(skill)
		}
	} else {
		fmt.Print("Aucun")
	}
	fmt.Print("\néquipement :\n    ")
	fmt.Print(u.stuff.head.name)
	fmt.Print("\n    ")
	fmt.Print(u.stuff.body.name)
	fmt.Print("\n    ")
	fmt.Print(u.stuff.legs.name)
	fmt.Printf("\n╘═══════════════════╛\n%s(%s'exit'%s\tpour quitter l'Inventaire\n", yellow, red, reset)
	read := readTer()
	if read == "exit" || read == "e" {
		clear()
	} else {
		clear()
		u.displayinfo()
	}
}

func (u *character) dead() bool {
	if u.hp <= 0 {
		fmt.Println("Vous êtes mort... WASTED!")
		u.alive = false
		u.hp = u.maxHp / 2
		u.alive = true
		fmt.Printf("Vous avez été ressuscité avec %d HP.\nVous avez perdu la moitié de votre argent\n", u.hp)
		u.purse /= 2
		return true
	}
	return false
}

func (u *character) spellBook(spell string) {
	// Vérifie si le sort est déjà appris
	for _, s := range u.skills {
		if s == spell {
			fmt.Println("Vous connaissez déjà ce sort.")
			return
		}
	}
	// Si le sort n'est pas déjà appris, on l'ajoute
	u.skills = append(u.skills, spell)
	u.removeInventory(fireSpellBook)
	fmt.Printf("Vous avez appris le sort: %s\n", spell)
}

func (u *character) HpActualise() {
	if u.classe == "humain" {
		u.maxHp = (u.lvl * 20) + u.stuff.head.buff + u.stuff.body.buff + u.stuff.legs.buff
	} else if u.classe == "elfe" {
		u.maxHp = (u.lvl * 15) + u.stuff.head.buff + u.stuff.body.buff + u.stuff.legs.buff
	} else if u.classe == "nain" {
		u.maxHp = (u.lvl * 30) + u.stuff.head.buff + u.stuff.body.buff + u.stuff.legs.buff
	}
}

func (u *character) updateXp() {
	yellow := "\033[33m"
	reset := "\033[0m"
	for u.exp >= u.maxExp {

		if u.exp >= u.maxExp {
			u.lvl += 1
			u.exp -= u.maxExp
			u.maxExp = 100 * u.lvl
		}
		u.HpActualise()
		fmt.Printf("%sFélicitation, vous avez monté de niveau, vous êtes actuellement au niveau %d!%s\n", yellow, u.lvl, reset)
	}
	u.hp = u.maxHp
}
