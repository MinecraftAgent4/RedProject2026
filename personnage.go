package main

import (
	"fmt"
)

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
	helmet Casque
	torso  Plastron
	boots  Bottes
}

func (c *Character) equipArmor(piece Armure) {
	emplacement := -1
	for index, object := range c.inventaire {
		if object.Nom() == piece.Nom() {
			emplacement = index
			break
		}
	}
	if emplacement == -1 {
		return
	}

	remplace := func(ancien Armure) {
		if ancien != nil {
			c.inventaire[emplacement] = ancien
		} else {
			c.inventaire = append(c.inventaire[:emplacement], c.inventaire[emplacement+1:]...)
		}
	}
	switch armor := piece.(type) {
	case Casque:
		ancien := c.equipe.helmet
		remplace(ancien)
		c.equipe.helmet = armor
	case Plastron:
		ancien := c.equipe.torso
		remplace(ancien)
		c.equipe.torso = armor
	case Bottes:
		ancien := c.equipe.boots
		remplace(ancien)
		c.equipe.boots = armor
	}
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
