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

func (m *Monster) exp_reward(c *Character) {
	if m.pv_actuelle <= 0 {
		c.exp_joueur += (m.pv_actuelle+m.weapon.Dmg())/2
	}
	if c.exp_joueur >= c.exp_required {
		c.lvlUp()
	}
}