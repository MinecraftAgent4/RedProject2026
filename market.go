package main

type Object struct {
	nom    string
	prix   int
	effect func(*Character)
}

type Market struct {
	liste_des_objet []Object
}

func initMarket() Market {
	market := Market{}
	market.add_potion(12, potion{nom: "steampack de basse qualité", effet: func(c *Character) { c.AddPV(20) }})
	market.add_potion(28, potion{nom: "steampack", effet: func(c *Character) { c.AddPV(50) }})
	market.add_potion(55, potion{nom: "steampack de grande qualité", effet: func(c *Character) { c.AddPV(80) }})
	market.add_object(40, "Grenade à fragmentation", nil)
	market.add_object(22, "Grenade fumigène", nil)
	market.add_object(48, "Grenade incendiaire", nil)
	market.add_object(35, "Grenade paralysante", nil)
	return market
}

func (m *Market) add_object(price int, object string, effet func(*Character)) {
	m.liste_des_objet = append(m.liste_des_objet, Object{nom: object, prix: price, effect: effet})
}

func (m *Market) add_potion(price int, pot potion) {
	m.add_object(price, pot.nom, pot.effet)
}
