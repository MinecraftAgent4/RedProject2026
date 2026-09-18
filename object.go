package main

type Object interface {
	Nom() string
}

type Potion struct {
	nom string
	effect func()
}

type Arme struct {
	nom string
	dmg int
}

type Armure struct {
	nom string
	defense int
}

type Item struct {
	nom string
}

type Spell struct {
	name string
	effect func()
}

type SpellBook struct {
	spell Spell
}

func (i Potion) Nom() string {return i.nom}
func (i Arme) Nom() string {return i.nom}
func (i Armure) Nom() string {return i.nom}
func (i Item) Nom() string {return i.nom}
func (i SpellBook) Nom() string {return "Livre de Sort : "+ i.spell.name}
