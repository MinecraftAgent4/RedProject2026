package main

import "fmt"

type Character struct {
	nom         string
	classe      string
	niveau      int
	pv_total    int
	pv_actuelle int
	inventaire  []string
}

func (perso Character) accessInventory() {
	for _,i := range perso.inventaire {
		fmt.Println("- "+i)
	}
}
