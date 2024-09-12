package main

import (
	"fmt"
	"math/rand"
)

type character struct {
	name    string
	classe  string
	lvl     int
	exp     int
	maxHp   int
	hp      int
	inv     []obj
	invSize int
	alive   bool
	purse   int
	skills  []string
}

var MyChar character = character{"jack", "elfe", 1, 0, 0, 0, []obj{}, 0, true, 200, nil}

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
	MyChar.hp = 1
	MyChar.skills = append(MyChar.skills, "Coup de poing")
}

func (u character) displayinfo() {
	clear()
	red := "\033[31m"
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Printf("╒════╡%sVos stats%s╞════╕\n name : %s\n classe : %s\n level : %d\n experience : %d/%d\n hp : %d/%d\n Purse : %d\n",
		yellow, reset, u.name, u.classe, u.lvl, u.exp, 100*u.lvl, u.hp, u.maxHp, u.purse)
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
	fmt.Printf("\n╘═══════════════════╛\n%s⎸%s'exit'%s\tpour quitter le marchand\n", yellow, red, reset)
	read := readTer()
	if read == "exit" {
		clear()
	} else {
		clear()
		u.displayinfo()
	}
}

func (u *character) takeDamage(dmg int) {
	u.hp -= dmg
	fmt.Printf("Vous avez pris %d points de dégâts. HP restants : %d/%d\n", dmg, u.hp, u.maxHp)
	u.dead()
}

func (u *character) dead() {
	if u.hp <= 0 {
		fmt.Println("Vous êtes mort... WASTED!")
		u.alive = false
		u.hp = u.maxHp / 2
		u.alive = true
		fmt.Printf("Vous avez été ressuscité avec %d HP.\nVous avez perdu la moitié de votre argent\n", u.hp)
		u.purse /= 2
		loop()
	}
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
	fmt.Printf("Vous avez appris le sort: %s\n", spell)
}
