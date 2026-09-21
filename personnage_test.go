package main

import "testing"

func TestEquipArmorEquipsPieceAndRemovesItFromInventory(t *testing.T) {
	armor := Armure{nom: "casque", defense: 2, emplacement: EmplacementCasque}
	character := Character{inventaire: []Object{armor}}

	character.equipArmor(armor)

	if character.equipe.helmet == nil || *character.equipe.helmet != armor {
		t.Fatalf("helmet = %#v, want %#v", character.equipe.helmet, armor)
	}
	if len(character.inventaire) != 0 {
		t.Fatalf("inventory length = %d, want 0", len(character.inventaire))
	}
}

func TestEquipArmorReturnsPreviousPieceToInventory(t *testing.T) {
	first := Armure{nom: "casque leger", defense: 1, emplacement: EmplacementCasque}
	second := Armure{nom: "casque lourd", defense: 3, emplacement: EmplacementCasque}
	character := Character{inventaire: []Object{first, second}}

	character.equipArmor(first)
	character.equipArmor(second)

	if character.equipe.helmet == nil || *character.equipe.helmet != second {
		t.Fatalf("helmet = %#v, want %#v", character.equipe.helmet, second)
	}
	if len(character.inventaire) != 1 || character.inventaire[0] != first {
		t.Fatalf("inventory = %#v, want the previous helmet", character.inventaire)
	}
}

func TestEquipArmorIgnoresUnknownPiece(t *testing.T) {
	armor := Armure{nom: "bottes", defense: 1, emplacement: EmplacementBottes}
	character := Character{}

	character.equipArmor(armor)

	if character.equipe.boots != nil || len(character.inventaire) != 0 {
		t.Fatalf("character changed when armor was absent from inventory: %#v", character)
	}
}