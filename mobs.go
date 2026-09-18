package main

type Monster struct {
	name string
	pv_total    int
	pv_actuelle int
	weapon Arme
}

func initGoblin() Monster {
	name := "Gobelin d'entraînement"
	pv_total := 40
	pv_actuelle := 40 
	weapon := Arme{"dague", 5}
}