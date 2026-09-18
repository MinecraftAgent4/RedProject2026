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
	helmet []Object
	torso []Object
	boots []Object
}

func (c *Character) AddInventory(object string) {
	c.inventaire = append(c.inventaire, Object{nom: object})
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
			if p.effet != nil {
				p.effet()
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
	fmt.Println(
		"nom : ", c.nom,
		'\n',
		"classe : ", c.classe,
		'\n',
		"PV : ", c.pv_actuelle, "/", c.pv_total,
		'\n',
	)
}

func (c *Character) isDead() {
	if c.pv_actuelle <= 0 {
		c.AddPV(c.pv_total / 2)
	}
}

func (c *Character) inventoryLimit() bool {
	if len(c.inventaire) == maxslots {
		return false
	}
	return true
}
func (c *Character) upgradeInventorySlot() {
	c.maxslots += 10
}