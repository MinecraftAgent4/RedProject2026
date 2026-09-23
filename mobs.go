package main

type Monster struct {
	name string
	pv_total    int
	pv_actuelle int
	weapon Arme
	armor Armure
}

type Entity interface {
	Name() string
	MaxHp() int
	Hp() int
	Atk() int
	Def() int
	AddPV(n int)
	Dmg(n int)
}

func (mob *Monster) Name() string {return mob.name}
func (mob *Monster) MaxHp() int {return mob.pv_total}
func (mob *Monster) Hp() int {return mob.pv_actuelle}
func (mob *Monster) Atk() int {return mob.weapon.Dmg()}
func (mob *Monster) Def() int {return mob.armor.Defense()}

func (mob *Monster) Dmg(n int) {
	mob.pv_actuelle -= n
	if mob.pv_actuelle < 0 {
		mob.pv_actuelle = 0
	}
}

func (mob *Monster) AddPV(n int) {
	mob.pv_actuelle += n
	if mob.pv_actuelle > mob.pv_total {
		mob.pv_actuelle = mob.pv_total
	}
}


func (guy *Character) Name() string {return guy.nom}
func (guy *Character) MaxHp() int {return guy.pv_total}
func (guy *Character) Hp() int {return guy.pv_actuelle}
func (guy *Character) Atk() int {return guy.pv_actuelle}
func (guy *Character) Def() int {
	e := guy.equipe
	return e.helmet.Defense() + e.torso.Defense() + e.boots.Defense()
}

func (guy *Character) Dmg(n int) {
	guy.pv_actuelle -= n
	if guy.pv_actuelle < 0 {
		guy.pv_actuelle = 0
	}
}

func (guy *Character) AddPV(n int) {
	guy.pv_actuelle += n
	if guy.pv_actuelle > guy.pv_total {
		guy.pv_actuelle = guy.pv_total
	}
}

func initMonster(nom string, hp int, arme Arme, armure Armure) Monster {
	return Monster{nom, hp, hp, arme, armure}
}

func initGoblin() Monster {
	return initMonster(
		"Goblin d'entraînement",
		40,
		Melee{"Dague", 5},
		nil,
	)
}