package main

import "fmt"

func initCharacter(nom string, classe string, niveau int, pv_total int, pv_actuelle int, inventaire []Object, money int) Character {
	return Character{
		nom:         nom,
		classe:      classe,
		niveau:      niveau,
		pv_total:    pv_total,
		pv_actuelle: pv_actuelle,
		inventaire:  inventaire,
		maxslots:    10,
		money:       money,
	}
}

func création_perso() Character {
	classe := []string{"netrunner", "merc", "cyberpsycho"}
	classe_choisie := ""
	classeValide := false

	for !classeValide {
		fmt.Println("Choisis une classe :")
		fmt.Println("1 -", classe[0])
		fmt.Println("2 -", classe[1])
		fmt.Println("3 -", classe[2])
		fmt.Scanln(&classe_choisie)

		if classe_choisie == "1" {
			classe_choisie = classe[0]
			classeValide = true
		}
		if classe_choisie == "2" {
			classe_choisie = classe[1]
			classeValide = true
		}
		if classe_choisie == "3" {
			classe_choisie = classe[2]
			classeValide = true
		}
		if !classeValide {
			fmt.Println("Choix invalide.")
		}
	}

	pv := 120
	if classe_choisie == "netrunner" {
		pv = 80
	}
	if classe_choisie == "merc" {
		pv = 100
	}

	nom := ""
	fmt.Print("Entre ton nom : ")
	fmt.Scanln(&nom)

	return initCharacter(nom, classe_choisie, 1, pv/2, pv, []Object{
		Resource{nom: "Ferraille", quantité: 10, quantité_max: 99},
		Resource{nom: "Composants", quantité: 4, quantité_max: 99},
		Resource{nom: "Poudre", quantité: 5, quantité_max: 99},
	}, 100)
}

func main() {
	statue := "menu"
	perso := création_perso()
	marketObjets := initMarket()
	charcudocObjets := initCharcudoc()

	for statue != "EXIT" {
		fmt.Println(menu)
		fmt.Print("Choix : ")
		if _, err := fmt.Scanln(&statue); err != nil {
			break
		}

		if statue == "1" {
			perso.displayInfo()
		}
		if statue == "2" {
			fmt.Println(marketMenu(marketObjets, perso))
			var choixMarket int
			fmt.Print("Choix : ")
			if _, err := fmt.Scanln(&choixMarket); err == nil {
				if choixMarket == len(marketObjets.liste_offres)+1 {
					continue
				}
				if choixMarket >= 1 && choixMarket <= len(marketObjets.liste_offres) {
					_, message := marketObjets.buy(&perso, choixMarket)
					fmt.Println(message)
				} else {
					fmt.Println("Choix invalide.")
				}
			}
		}
		if statue == "5" {
			perso.accessInventory()
		}
		if statue == "3" {
			fmt.Println(charcudocMenu(charcudocObjets, perso))
			var choixCharcudoc int
			fmt.Print("Choix : ")
			if _, err := fmt.Scanln(&choixCharcudoc); err == nil {
				if choixCharcudoc == len(charcudocObjets.liste_offres)+1 {
					continue
				}
				if choixCharcudoc >= 1 && choixCharcudoc <= len(charcudocObjets.liste_offres) {
					_, message := charcudocObjets.buy(&perso, choixCharcudoc)
					fmt.Println(message)
				} else {
					fmt.Println("Choix invalide.")
				}
			}
		}
		if statue == "4" {
			fmt.Println("Cette fonctionnalite n'est pas encore disponible.")
		}
		if statue != "1" && statue != "2" && statue != "3" && statue != "4" && statue != "5" && statue != "6" {
			fmt.Println("Choix invalide : entre un nombre entre 1 et 6.")
		}
		if statue == "6" {
			statue = "EXIT"
		}
	}
	fmt.Println("Au revoir !")

}
