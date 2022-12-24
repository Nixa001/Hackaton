package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		designation string
		quantite    int
		prixUnit, total, payee, change float64
	)

	// Créer un lecteur de saisie utilisateur
	reader := bufio.NewReader(os.Stdin)

	// Saisir et traiter les données pour chaque article
	for {
		fmt.Println("Entrez la désignation de l'article (ou tapez 'q' pour quitter) :")
		designation, _ = reader.ReadString('\n')
		designation = strings.TrimSpace(designation)

		if designation == "q" {
			break
		}

		fmt.Println("Entrez la quantité :")
		fmt.Scan(&quantite)

		fmt.Println("Entrez le prix unitaire :")
		fmt.Scan(&prixUnit)

		// Mettre à jour le total
		total += float64(quantite) * prixUnit
	}

	// Afficher le montant total
	fmt.Printf("Le montant total est de %.2f\n", total)

	// Saisir et traiter la somme versée
	fmt.Println("Entrez la somme versée :")
	fmt.Scan(&payee)

	// Calculer et afficher la monnaie
	change = payee - total
	fmt.Printf("Le montant de la monnaie est de %.2f\n", change)
}

