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
	equipe Equipement
}

type Equipement struct {
	helmet Casque
	torso Plastron
	boots Bottes
}

func (c *Character) GiveItem(object Object) {
	if c.inventoryLimit() {
		c.inventaire = append(c.inventaire, object)
	}
}

func (c *Character) AddInventory(object Object) {
	c.GiveItem(object)
}

func (c Character) resourceQuantity(name string) int {
	quantity := 0
	for _, object := range c.inventaire {
		resource, ok := object.(Resource)
		if ok && resource.nom == name {
			quantity += resource.quantité
		}
	}
	return quantity
}

func (c *Character) removeResource(name string, quantity int) {
	for index := 0; index < len(c.inventaire) && quantity > 0; index++ {
		resource, ok := c.inventaire[index].(Resource)
		if !ok || resource.nom != name {
			continue
		}
		if resource.quantité <= quantity {
			quantity -= resource.quantité
			c.inventaire = append(c.inventaire[:index], c.inventaire[index+1:]...)
			index--
			continue
		}
		resource.quantité -= quantity
		c.inventaire[index] = resource
		quantity = 0
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
				p.effect(*c)
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
