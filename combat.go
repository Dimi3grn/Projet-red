package main

type mob struct {
	name  string
	maxhp int
	hp    int
	dmg   int
}

var goblin mob = mob{"Goblin", 50, 50, 2}
