package main

type Object struct {
	nom    string
	prix   int
	effect func()
}

type Market struct {
	liste_des_objet []Object
}

func initMarket() Market {
	market := Market{}
	market.add_object(12, "steampack de basse qualité", nil)
	market.add_object(28, "steampack", nil)
	market.add_object(55, "steampack de grande qualité", nil)
	market.add_object(40, "Grenade à fragmentation", nil)
	market.add_object(22, "Grenade fumigène", nil)
	market.add_object(48, "Grenade incendiaire", nil)
	market.add_object(35, "Grenade paralysante", nil)
	return market
}

func (m *Market) add_object(price int, object string, effet func()) {
	m.liste_des_objet = append(m.liste_des_objet, Object{nom: object, prix: price, effect: effet})
}

func (m *Market) add_potion(price int, pot potion) {
	m.add_object(price, pot.nom, pot.effet)
}
