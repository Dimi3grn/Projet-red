package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Selection() {
	reader := bufio.NewReader(os.Stdin)
	red := "\033[31m"
	Bcyan := "\033[1;36m"
	yellow := "\033[33m"
	reset := "\033[0m"

	for {
		fmt.Printf("%sEntrez votre nom%s (%suniquement en lettres%s)\n : ", yellow, reset, Bcyan, reset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Vérification : le nom contient uniquement des lettres
		if isOnlyLetters(input) {
			// Formater le nom : première lettre en majuscule, le reste en minuscules
			formattedName := formatName(input)
			fmt.Printf("Nom choisi : %s\n", formattedName)
			MyChar.name = formattedName
			break
		} else {
			clear()
			fmt.Printf("%sErreur%s : Le nom doit contenir %suniquement des lettres%s.\n", red, reset, Bcyan, reset)
		}
	}
}

// Vérifie si une chaîne contient uniquement des lettres
func isOnlyLetters(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

// Formate le nom : première lettre en majuscule, reste en minuscules
func formatName(name string) string {
	if len(name) == 0 {
		return ""
	}
	name = strings.ToLower(name)                       // Tout mettre en minuscules
	return strings.ToUpper(string(name[0])) + name[1:] // Première lettre en majuscule, reste en minuscules
}
