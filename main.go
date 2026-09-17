package main

import "fmt"

func initCharacter(nom string, classe string, niveau int, pv_total int, pv_actuelle int, inventaire []string) Character {
	return Character{nom, classe, niveau, pv_total, pv_actuelle, inventaire}
}

func main() {
	statue := "menu"
	perso := initCharacter("teste", "netrunner", 1, 20, 20, []string{"weapon#1", "helmet", "boots"})
	for statue != "EXIT" {
		fmt.Println(menu)
		fmt.Print("Choix : ")
		fmt.Scanln(&statue)

		if statue == "1" {
			perso.displayInfo()
		}
		if statue == "5" {
			perso.accessInventory()
		}
		if statue == "2" || statue == "3" || statue == "4" {
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
