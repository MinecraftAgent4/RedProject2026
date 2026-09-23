package main

func Attack(a Entity, t Entity, mod int) {
	t.Dmg((a.Atk() * mod) - t.Def())
}

func (m *Monster) goblinPattern(turn int, t *Character) {
	if turn%3 == 2 {
		Attack(m,t, 2)
	} else {
		Attack(m,t, 1)
	}
}