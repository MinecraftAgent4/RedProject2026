package main

import "fmt"

func initCharacter(nom string, classe string, niveau int, pv_total int, pv_actuelle int, inventaire []Object) Character {
	return Character{
		nom:         nom,
		classe:      classe,
		niveau:      niveau,
		pv_total:    pv_total,
		pv_actuelle: pv_actuelle,
		inventaire:  inventaire,
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

	return initCharacter(nom, classe_choisie, 1, pv, pv, []Object{})
}

func main() {
	statue := "menu"
	perso := création_perso()
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
			fmt.Println(market)
			var choixMarket int
			fmt.Print("Choix : ")
			if _, err := fmt.Scanln(&choixMarket); err == nil {
				if choixMarket == 1 {
					perso.AddInventory("steampack de basse qualité")
				}
				if choixMarket == 2 {
					perso.AddInventory("steampack")
				}
				if choixMarket == 3 {
					perso.AddInventory("steampack de grande qualité")
				}
				if choixMarket == 4 {
					perso.AddInventory("Grenade à fragmentation")
				}
				if choixMarket == 5 {
					perso.AddInventory("Grenade fumigène")
				}
				if choixMarket == 6 {
					perso.AddInventory("Grenade incendiaire")
				}
				if choixMarket == 7 {
					perso.AddInventory("Grenade paralysante")
				}
				if choixMarket == 8 {
					continue
				}
				if choixMarket < 1 || choixMarket > 8 {
					fmt.Println("Choix invalide : entre un nombre entre 1 et 8.")
				}
			}
		}
		if statue == "5" {
			perso.accessInventory()
		}
		if statue == "3" || statue == "4" {
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
