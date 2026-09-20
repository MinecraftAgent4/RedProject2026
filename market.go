package main

type Market struct {
	liste_offres []Trade
}

type Trade struct {
	result      Object
	price       int
	ingredients []Object
}

func (m *Market) add_object(object Object, price int, ingredients []Object) {
	m.liste_offres = append(m.liste_offres, Trade{object, price, ingredients})
}
