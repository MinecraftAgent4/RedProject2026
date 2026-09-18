package main

import "fmt"

type Character struct {
	nom         string
	classe      string
	niveau      int
	pv_total    int
	pv_actuelle int
	inventaire  []Object
	maxslots    int
	money       int
	equipe      equipement
}

type equipement struct {
	helmet []Object
	torso  []Object
	boots  []Object
}

func (c *Character) AddInventory(object Object) {
	c.inventaire = append(c.inventaire, object)
}

func (c *Character) AddPV(x int) {
	c.pv_actuelle += x

	if c.pv_actuelle > c.pv_total {
		c.pv_actuelle = c.pv_total
	}
}

func (c *Character) TakePot(p potion) {
	for i, objet := range c.inventaire {
		if objet.nom == p.nom {
			if objet.effect != nil {
				objet.effect(c)
			}

			c.inventaire = append(c.inventaire[:i], c.inventaire[i+1:]...)
			return
		}
	}
}

func (perso Character) accessInventory() {
	for _, i := range perso.inventaire {
		fmt.Println("- " + i.nom)
	}
}

func (c *Character) displayInfo() {
	fmt.Printf("nom : %s\nclasse : %s\nPV : %d / %d\n", c.nom, c.classe, c.pv_actuelle, c.pv_total)
}

func (c *Character) isDead() {
	if c.pv_actuelle <= 0 {
		c.AddPV(c.pv_total / 2)
	}
}
func (c *Character) inventoryLimit() bool {
	if len(c.inventaire) == c.maxslots {
		return false
	}
	return true
}
