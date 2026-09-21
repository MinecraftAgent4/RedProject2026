package main

type Monster struct {
	name        string
	pv_total    int
	pv_actuelle int
	weapon      Arme
}

type Entity interface {
	Name() string
	MaxHp() int
	Hp() int
}

func (mob Monster) Name() string { return mob.name }
func (mob Monster) MaxHp() int   { return mob.pv_total }
func (mob Monster) Hp() int      { return mob.pv_actuelle }

func (guy Character) Name() string { return guy.nom }
func (guy Character) MaxHp() int   { return guy.pv_total }
func (guy Character) Hp() int      { return guy.pv_actuelle }

func initGoblin() Monster {
	name := "Gobelin d'entraînement"
	pv_total := 40
	pv_actuelle := 40
	weapon := Arme{"dague", 5}
	return Monster{name: name, pv_total: pv_total, pv_actuelle: pv_actuelle, weapon: weapon}
}
