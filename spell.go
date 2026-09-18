package main

type Spell struct {
	name string
	effect func()
}

func createSpellBook(spell Spell, price int) Object {
	return 	Object{"Livre de Sort : " + spell.name, price, spell.effect}
}