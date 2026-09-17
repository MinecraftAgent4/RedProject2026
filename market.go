package main

type Object struct {
	nom    string
	prix   int
	effect func()
}

type Market struct {
	liste_des_objet []Object
}

func (m *Market) add_object(price int, object string, effet func()) {
	m.liste_des_objet = append(m.liste_des_objet, Object{nom: object, prix: price, effect: effet})
}

func (m *Market) add_potion(price int, pot potion) {
	m.add_object(price, pot.nom, pot.effet)
}
