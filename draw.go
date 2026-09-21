package main

import (
	"fmt"
	"strings"
	"unicode"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawMap(player *rl.Rectangle, blocks []mapBlock, characters []mapCharacter, message string, perso *Character) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(28, 37, 61, 255))

	for _, block := range blocks {
		rl.DrawRectangleRec(block.rect, block.color)
		rl.DrawRectangleLinesEx(block.rect, 2, rl.DarkGray)
	}

	for _, character := range characters {
		rl.DrawRectangleRec(character.rect, character.color)
		rl.DrawText(character.name, int32(character.rect.X-20), int32(character.rect.Y-22), 12, rl.RayWhite)
	}

	rl.DrawRectangleRec(*player, rl.SkyBlue)

	rl.DrawText(message, 24, 20, 18, rl.RayWhite)
	rl.DrawText(perso.nom+" | "+perso.classe+" | PV "+itoa(perso.pv_actuelle)+"/"+itoa(perso.pv_total),
		24, 50, 18, rl.Gold)
	rl.DrawText("Déplacement : A/D ou ←/→ | Saut : Espace | Menu : M",
		24, 510, 16, rl.LightGray)

	rl.EndDrawing()
}

func drawClassScreen(selected string) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(28, 37, 61, 255))
	drawPanel("CRÉATION DU PERSONNAGE")

	rl.DrawText("Choisis une classe :", 90, 130, 28, rl.RayWhite)
	rl.DrawText("1 - netrunner       (80 PV)", 110, 200, 24, rl.RayWhite)
	rl.DrawText("2 - merc            (100 PV)", 110, 245, 24, rl.RayWhite)
	rl.DrawText("3 - cyberpsycho     (120 PV)", 110, 290, 24, rl.RayWhite)

	if selected != "" {
		rl.DrawText("Classe sélectionnée : "+selected, 110, 360, 22, rl.Gold)
	}
	rl.DrawText("Appuie sur 1, 2 ou 3", 110, 420, 18, rl.LightGray)
	rl.EndDrawing()
}

func drawGamePanel(screen string, perso *Character, market *Market, charcudoc *Market, message string) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(28, 37, 61, 255))

	switch screen {
	case "menu":
		drawMenuPanel()
	case "perso":
		drawCharacterPanel(perso)
	case "inventaire":
		drawInventoryPanel(perso)
	case "marchand":
		drawMarketPanel("MARCHAND", market)
	case "charcudoc":
		drawMarketPanel("CHARCUDOC", charcudoc)
	case "gang":
		drawGangPanel()
	}

	if screen != "menu" && screen != "gang" {
		rl.DrawText(message, 50, 475, 18, rl.Gold)
	}
	rl.DrawText("Échap : retour", 760, 500, 16, rl.LightGray)
	rl.EndDrawing()
}

func drawMenuPanel() {
	drawPanel("LE QUARTIER")
	entries := []string{
		"1. Perso",
		"2. Inventaire",
		"3. EXIT",
	}

	for i, entry := range entries {
		rl.DrawText(entry, 110, int32(135+i*48), 25, rl.RayWhite)
	}
	rl.DrawText("Les services sont disponibles auprès des PNJ sur la carte.", 110, 330, 17, rl.LightGray)
	rl.DrawText("Choisis un numéro pour agir", 110, 385, 18, rl.LightGray)
}

func drawCharacterPanel(perso *Character) {
	drawPanel("PERSONNAGE")
	rl.DrawText("Nom : "+perso.nom, 90, 130, 25, rl.RayWhite)
	rl.DrawText("Classe : "+perso.classe, 90, 175, 25, rl.RayWhite)
	rl.DrawText("Niveau : "+itoa(perso.niveau), 90, 220, 25, rl.RayWhite)
	rl.DrawText("PV : "+itoa(perso.pv_actuelle)+" / "+itoa(perso.pv_total), 90, 265, 25, rl.RayWhite)
	rl.DrawText("Argent : "+itoa(perso.money)+" po", 90, 310, 25, rl.Gold)
	rl.DrawText("Emplacements : "+itoa(len(perso.inventaire))+" / "+itoa(perso.maxslots),
		90, 355, 25, rl.RayWhite)
}

func drawInventoryPanel(perso *Character) {
	drawPanel("INVENTAIRE")

	if len(perso.inventaire) == 0 {
		rl.DrawText("Inventaire vide.", 90, 145, 24, rl.RayWhite)
		return
	}

	for i, object := range perso.inventaire {
		if i >= 8 {
			break
		}
		rl.DrawText(fmt.Sprintf("%d. %s", i+1, object.Nom()),
			90, int32(120+i*42), 21, rl.RayWhite)
	}
}

func drawMarketPanel(title string, market *Market) {
	drawPanel(title)

	for i, trade := range market.liste_offres {
		line := fmt.Sprintf("%d. %s", i+1, trade.result.Nom())

		if trade.price > 0 {
			line += " - " + itoa(trade.price) + " po"
		}

		if len(trade.ingredients) > 0 {
			ingredients := make([]string, 0, len(trade.ingredients))
			for _, ingredient := range trade.ingredients {
				ingredients = append(ingredients,
					fmt.Sprintf("%d %s", ingredient.quantité, ingredient.nom))
			}
			line += " [" + strings.Join(ingredients, ", ") + "]"
		}

		rl.DrawText(line, 70, int32(125+i*55), 20, rl.RayWhite)
	}

	rl.DrawText("4. Retour à la carte", 70, 355, 20, rl.RayWhite)
	rl.DrawText("Appuie sur le numéro de l'offre pour acheter.", 70, 405, 17, rl.LightGray)
}

func drawGangPanel() {
	drawPanel("GUERRE DE GANG")
	rl.DrawText("Le gang est prêt à se battre.", 90, 145, 24, rl.RayWhite)
	rl.DrawText("Cette fonctionnalité n'est pas encore disponible.", 90, 195, 20, rl.LightGray)
	rl.DrawText("4 : retour à la carte", 90, 260, 20, rl.Gold)
}

func drawPanel(title string) {
	rl.DrawRectangle(45, 35, 870, 440, rl.NewColor(40, 48, 70, 255))
	rl.DrawRectangleLines(45, 35, 870, 440, rl.Gold)
	rl.DrawText(title, 70, 65, 32, rl.Gold)
}

func readNameInput(current string) string {
	if rl.IsKeyPressed(rl.KeyBackspace) && len(current) > 0 {
		return current[:len(current)-1]
	}

	for {
		char := rl.GetCharPressed()
		if char == 0 {
			break
		}

		r := rune(char)
		if unicode.IsLetter(r) && len(current) < 18 {
			current += string(r)
		}
	}
	return current
}
