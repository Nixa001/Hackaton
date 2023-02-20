package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		designation                    string
		quantite                       int
		prixUnit, total, payee, change float64
	)

	reader := bufio.NewReader(os.Stdin)

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

		total += float64(quantite) * prixUnit
	}

	fmt.Printf("Le montant total est de %.2f\n", total)

	fmt.Println("Entrez la somme versée :")
	fmt.Scan(&payee)

	change = payee - total
	fmt.Printf("Le montant de la monnaie est de %.2f\n", change)
}
