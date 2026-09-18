package main

type Spell struct {
	name   string
	effect func(*Character)
}

func createSpellBook(spell Spell, price int) Object {
	return Object{
		nom:    "Livre de Sort : " + spell.name,
		prix:   price,
		effect: spell.effect,
	}
}
