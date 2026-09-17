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

func (c *Character) AddPV(x int) {
	c.pv_actuelle += x

	if c.pv_actuelle > c.pv_total {
		c.pv_actuelle = c.pv_total
	}
}

func (c *Character) TakePot(p potion) {
	for i, objet := range c.inventaire {
		if objet == p.nom {
			if p.effet != nil {
				p.effet()
			}

			c.inventaire = append(c.inventaire[:i], c.inventaire[i+1:]...)
			return
	}
}
}

func (perso Character) accessInventory() {
	for _,i := range perso.inventaire {
		fmt.Println("- "+i)
	}
}
