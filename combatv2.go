package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Définir la structure des personnages
type Personnage struct {
	nom   string
	hp    int
	attk  int
	def   int
	purse int
}

// Fonction d'attaque
func (p *Personnage) Attaq(cible *Personnage) {
	// Calculer les dégâts infligés
	deg := p.attk - cible.def
	if deg < 1 {
		deg = 1 // Les dégâts minimum sont de 1
	}

	// Réduire les PV de la cible
	cible.hp -= deg
	fmt.Printf("%s attaque %s et inflige %d dégâts!\n", p.nom, cible.nom, deg)

	if cible.hp <= 0 {
		fmt.Printf("%s est mort!\n", cible.nom)
	}
}

// Vérifier si un personnage est encore vivant
func (p *Personnage) Alive() bool {
	return p.hp > 0
}

// Fonction pour lancer un sort
func (u *character) useSpell(spell string, cible *Personnage) {
	// Vérifier si le joueur connaît le sort
	for _, s := range u.skills {
		if s == spell {
			fmt.Printf("%s utilise le sort %s sur %s!\n", u.name, spell, cible.nom)
			deg := 10 // Dégâts fixes pour le sort "Boule de feu"
			cible.hp -= deg
			fmt.Printf("Le sort %s inflige %d dégâts à %s!\n", spell, deg, cible.nom)
			if cible.hp <= 0 {
				fmt.Printf("%s est mort!\n", cible.nom)
			}
			return
		}
	}
	fmt.Println("Vous ne connaissez pas ce sort.")
}

// Fonction de combat tour par tour avec sorts
func Combat(joueur *character, ennemi *Personnage) {
	fmt.Printf("Le combat commence entre %s et %s!\n", joueur.name, ennemi.nom)

	// Boucle du combat
	for joueur.Alive() && ennemi.Alive() {
		fmt.Println("\n--- Nouveau Tour ---")
		fmt.Printf("%s: %d HP | %s: %d HP\n", joueur.name, joueur.hp, ennemi.hp)

		// Le joueur choisit entre une attaque normale ou un sort
		fmt.Println("Voulez-vous attaquer (a) ou utiliser un sort (s) ?")
		var action string
		fmt.Scan(&action)

		switch action {
		case "a":
			joueur.Attaq(ennemi) // Attaque normale
		case "s":
			fmt.Println("Quel sort voulez-vous utiliser ?")
			for i, s := range joueur.skills {
				fmt.Printf("%d. %s\n", i+1, s)
			}
			var choixSort int
			fmt.Scan(&choixSort)
			if choixSort > 0 && choixSort <= len(joueur.skills) {
				joueur.useSpell(joueur.skills[choixSort-1], ennemi) // Utiliser le sort choisi
			} else {
				fmt.Println("Sort invalide.")
			}
		default:
			fmt.Println("Action non reconnue.")
		}

		if !ennemi.Alive() {
			fmt.Printf("%s a gagné le combat!\n", joueur.name)
			// Le joueur reçoit une récompense en pièces d'or
			gain := rand.Intn(50) + 10 // Gain aléatoire entre 10 et 50 pièces
			joueur.purse += gain
			fmt.Printf("Vous avez gagné %d pièces d'or ! Votre bourse contient maintenant %d pièces d'or.\n", gain, joueur.purse)
			break
		}

		// L'ennemi attaque
		ennemi.Attaq(joueur)
		if !joueur.Alive() {
			fmt.Printf("%s a gagné le combat!\n", ennemi.nom)
			break
		}
	}
}

// Modifier cette fonction pour appeler le combat avec les sorts
func Combatstr() {
	rand.Seed(time.Now().UnixNano())

	// Création du joueur et de l'ennemi
	joueur := &MyChar // Utilisation de votre personnage

	ennemi := Personnage{
		nom:   "Monstre",
		hp:    10,
		attk:  8,
		def:   2,
		purse: 0, // Les ennemis n'ont pas de bourse ici
	}

	// Lancer le combat
	Combat(joueur, &ennemi)
}
