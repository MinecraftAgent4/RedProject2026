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
	equipe equipement
}

type equipement struct {
	helmet Armure
	torso Armure
	boots Armure
}

func (c *Character) GiveItem(object Object) {
	if c.inventoryLimit() {
		c.inventaire = append(c.inventaire, object)
	}
}

func (c *Character) AddPV(x int) {
	c.pv_actuelle += x

	if c.pv_actuelle > c.pv_total {
		c.pv_actuelle = c.pv_total
	}
}

func (c *Character) TakePot(p Potion) {
	for i, objet := range c.inventaire {
		if objet.Nom() == p.nom {
			if p.effect != nil {
				p.effect()
			}

			c.inventaire = append(c.inventaire[:i], c.inventaire[i+1:]...)
			return
		}
	}
}

func (perso Character) accessInventory() {
	for _, i := range perso.inventaire {
		fmt.Println("- " + i.Nom())
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
