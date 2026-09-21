package main

type Market struct {
	liste_offres []Trade
}

type Trade struct {
	result      Object
	ingredients []Resource
}

func (m *Market) addObject(object Object, ingredients ...Resource) {
	m.liste_offres = append(m.liste_offres, Trade{result: object, ingredients: ingredients})
}

func initMarket() Market {
	var market Market
	market.addObject(Item{nom: "Steampack de basse qualité"}, Resource{nom: "Ferraille", quantité: 3})
	market.addObject(Item{nom: "Steampack"}, Resource{nom: "Ferraille", quantité: 5}, Resource{nom: "Composants", quantité: 2})
	market.addObject(Item{nom: "Grenade à fragmentation"}, Resource{nom: "Ferraille", quantité: 2}, Resource{nom: "Poudre", quantité: 3})
	return market
}

func (m Market) buy(character *Character, choice int) (bool, string) {
	if choice < 1 || choice > len(m.liste_offres) {
		return false, "Choix invalide."
	}

	trade := m.liste_offres[choice-1]
	if !character.inventoryLimit() {
		return false, "Inventaire plein."
	}
	for _, ingredient := range trade.ingredients {
		if character.resourceQuantity(ingredient.nom) < ingredient.quantité {
			return false, "Ressources insuffisantes."
		}
	}
	for _, ingredient := range trade.ingredients {
		character.removeResource(ingredient.nom, ingredient.quantité)
	}
	character.GiveItem(trade.result)
	return true, trade.result.Nom() + " fabriqué."
}
