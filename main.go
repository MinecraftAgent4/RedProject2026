package main

import "fmt"

func initCharacter(nom string, classe string, niveau int, pv_total int, pv_actuelle int, inventaire []string) Character {
	return Character{nom, classe, niveau, pv_total, pv_actuelle, inventaire}
}

func main() {
	statue := "menu"
	perso := initCharacter("teste", "netrunner", 1, 20, 20, []string{"weapon#1", "helmet", "boots"})
	for statue == "EXIT" {
		fmt.Println(statue)
		if statue == "perso" {
			perso.displayInfo()
		}
		if statue == "inventaire" {
			perso.accessInventory()
		}
	}

}
