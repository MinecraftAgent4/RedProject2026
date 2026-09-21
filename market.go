package main

type Market struct {
	liste_offres []Trade
}

type Trade struct {
	result      Object
	price       int
	ingredients []Resource
}

func (m *Market) addObject(object Object, ingredients ...Resource) {
	m.liste_offres = append(m.liste_offres, Trade{result: object, ingredients: ingredients})
}

func (m *Market) addGoldObject(object Object, price int) {
	m.liste_offres = append(m.liste_offres, Trade{result: object, price: price})
}

func initMarket() Market {
	var market Market
	market.addGoldObject(Item{nom: "Steampack de basse qualité"}, 12)
	market.addGoldObject(Item{nom: "Steampack"}, 28)
	market.addGoldObject(Item{nom: "Grenade à fragmentation"}, 40)
	return market
}

func initCharcudoc() Market {
	var charcudoc Market
	charcudoc.addObject(Item{nom: "Interface neurale"}, Resource{nom: "Composants", quantité: 3}, Resource{nom: "Poudre", quantité: 1})
	charcudoc.addObject(Item{nom: "Optiques cybernétiques"}, Resource{nom: "Composants", quantité: 4}, Resource{nom: "Ferraille", quantité: 2})
	charcudoc.addObject(Item{nom: "Réflexes augmentés"}, Resource{nom: "Composants", quantité: 5}, Resource{nom: "Poudre", quantité: 2})
	return charcudoc
}

func (m Market) buy(character *Character, choice int) (bool, string) {
	if choice < 1 || choice > len(m.liste_offres) {
		return false, "Choix invalide."
	}

	trade := m.liste_offres[choice-1]
	if !character.inventoryLimit() {
		return false, "Inventaire plein."
	}
	if character.money < trade.price {
		return false, "Pas assez de pièces d'or."
	}
	for _, ingredient := range trade.ingredients {
		if character.resourceQuantity(ingredient.nom) < ingredient.quantité {
			return false, "Ressources insuffisantes."
		}
	}
	for _, ingredient := range trade.ingredients {
		character.removeResource(ingredient.nom, ingredient.quantité)
	}
	character.money -= trade.price
	character.GiveItem(trade.result)
	return true, trade.result.Nom() + " acheté."
}
