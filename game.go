package main

import (
	"fmt"
	"strings"
	"unicode"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type mapBlock struct {
	rect  rl.Rectangle
	color rl.Color
}

type mapCharacter struct {
	name  string
	rect  rl.Rectangle
	color rl.Color
}

func game() {
	rl.InitWindow(960, 540, "RedProject")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	marketObjets := initMarket()
	charcudocObjets := initCharcudoc()

	perso := initCharacter(
		"Joueur",
		"",
		1,
		100,
		100,
		[]Object{
			Resource{nom: "Ferraille", quantité: 10, quantité_max: 99},
			Resource{nom: "Composants", quantité: 4, quantité_max: 99},
			Resource{nom: "Poudre", quantité: 5, quantité_max: 99},
		},
		100,
	)

	screen := "classe"
	nameInput := ""
	selectedClass := ""

	blocks := []mapBlock{
		{rect: rl.Rectangle{X: 0, Y: 450, Width: 960, Height: 90}, color: rl.DarkGreen},
		{rect: rl.Rectangle{X: 80, Y: 360, Width: 180, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 350, Y: 290, Width: 180, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 630, Y: 360, Width: 180, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 820, Y: 220, Width: 100, Height: 24}, color: rl.Green},
	}

	characters := []mapCharacter{
		{name: "MARCHAND", rect: rl.Rectangle{X: 145, Y: 318, Width: 30, Height: 42}, color: rl.Gold},
		{name: "CHARCUDOC", rect: rl.Rectangle{X: 420, Y: 248, Width: 30, Height: 42}, color: rl.Purple},
		{name: "GUERRE DE GANG", rect: rl.Rectangle{X: 650, Y: 318, Width: 30, Height: 42}, color: rl.Red},
	}

	player := rl.Rectangle{X: 40, Y: 408, Width: 32, Height: 42}
	playerSpeed := float32(220)
	velocityY := float32(0)
	onGround := true
	message := "Explore la map — M pour ouvrir le menu"

	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()
		if screen == "classe" {
			if rl.IsKeyPressed(rl.KeyOne) {
				selectedClass = "netrunner"
				perso.classe = selectedClass
				perso.pv_total = 80
				perso.pv_actuelle = 80
				screen = "nom"
			}
			if rl.IsKeyPressed(rl.KeyTwo) {
				selectedClass = "merc"
				perso.classe = selectedClass
				perso.pv_total = 100
				perso.pv_actuelle = 100
				screen = "nom"
			}
			if rl.IsKeyPressed(rl.KeyThree) {
				selectedClass = "cyberpsycho"
				perso.classe = selectedClass
				perso.pv_total = 120
				perso.pv_actuelle = 120
				screen = "nom"
			}
			drawClassScreen(selectedClass)
			continue
		}

		if screen == "nom" {
			nameInput = readNameInput(nameInput)

			if rl.IsKeyPressed(rl.KeyEnter) {
				if validName(nameInput) {
					perso.nom = formatName(nameInput)
					screen = "map"
					message = "Bienvenue " + perso.nom + " ! — M pour ouvrir le menu"
				}
			}

			rl.BeginDrawing()
			rl.ClearBackground(rl.NewColor(28, 37, 61, 255))
			drawPanel("CRÉATION DU PERSONNAGE")
			rl.DrawText("Choisis ton nom :", 90, 150, 24, rl.RayWhite)
			rl.DrawRectangle(90, 200, 780, 55, rl.DarkGray)
			rl.DrawRectangleLines(90, 200, 780, 55, rl.Gold)
			rl.DrawText(nameInput+"_",
				110, 215, 24, rl.RayWhite)
			rl.DrawText("Entrée : valider    |    Retour arrière : effacer",
				90, 300, 18, rl.LightGray)
			rl.DrawText("Classe : "+selectedClass, 90, 340, 20, rl.Gold)
			rl.EndDrawing()
			continue
		}

		if screen != "map" {
			handlePanelInput(&screen, &perso, &marketObjets, &charcudocObjets, &message)
			drawGamePanel(screen, &perso, &marketObjets, &charcudocObjets, message)
			continue
		}

		if rl.IsKeyPressed(rl.KeyM) {
			screen = "menu"
			continue
		}
		direction := float32(0)
		if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
			direction = 1
		}
		if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
			direction = -1
		}

		if rl.IsKeyPressed(rl.KeySpace) && onGround {
			velocityY = -520
			onGround = false
		}

		velocityY += 1000 * delta
		player.Y += velocityY * delta
		onGround = false

		for _, block := range blocks {
			if velocityY >= 0 && overlapsX(player, block.rect) {
				previousBottom := player.Y - velocityY*delta + player.Height
				if previousBottom <= block.rect.Y && player.Y+player.Height >= block.rect.Y {
					player.Y = block.rect.Y - player.Height
					velocityY = 0
					onGround = true
				}
			}
		}

		moveHorizontal(&player, direction*playerSpeed*delta, blocks)

		if player.X < 0 {
			player.X = 0
		}
		if player.X+player.Width > 960 {
			player.X = 960 - player.Width
		}
		if player.Y > 540 {
			player.X = 40
			player.Y = 408
			velocityY = 0
			onGround = true
		}

		message = "Explore la map — M pour ouvrir le menu"

		for _, character := range characters {
			if rl.CheckCollisionRecs(player, character.rect) {
				message = "Appuie sur E pour parler au " + character.name
				if rl.IsKeyPressed(rl.KeyE) {
					switch character.name {
					case "MARCHAND":
						screen = "marchand"
					case "CHARCUDOC":
						screen = "charcudoc"
					case "GUERRE DE GANG":
						screen = "gang"
					}
				}
			}
		}

		drawMap(&player, blocks, characters, message, &perso)
	}

	fmt.Println("Au revoir !")
}

func handlePanelInput(screen *string, perso *Character, market *Market, charcudoc *Market, message *string) {
	if rl.IsKeyPressed(rl.KeyEscape) {
		
		*screen = "map"
		return
	}

	if *screen == "menu" {
		switch {
		case rl.IsKeyPressed(rl.KeyOne):
			*screen = "perso"
		case rl.IsKeyPressed(rl.KeyTwo):
			*screen = "inventaire"
		case rl.IsKeyPressed(rl.KeyThree):
			rl.CloseWindow()
		}
		return
	}

	if *screen == "marchand" {
		if rl.IsKeyPressed(rl.KeyOne) {
			_, result := market.buy(perso, 1)
			*message = result
			return
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			_, result := market.buy(perso, 2)
			*message = result
			return
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			_, result := market.buy(perso, 3)
			*message = result
			return
		}
		if rl.IsKeyPressed(rl.KeyFour) {
			*screen = "map"
		}
		return
	}

	if *screen == "charcudoc" {
		if rl.IsKeyPressed(rl.KeyOne) {
			_, result := charcudoc.buy(perso, 1)
			*message = result
			return
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			_, result := charcudoc.buy(perso, 2)
			*message = result
			return
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			_, result := charcudoc.buy(perso, 3)
			*message = result
			return
		}
		if rl.IsKeyPressed(rl.KeyFour) {
			*screen = "map"
		}
		return
	}

	if *screen == "gang" {
		if rl.IsKeyPressed(rl.KeyFour) {
			*screen = "map"
		}
		return
	}
}

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

func validName(name string) bool {
	if name == "" {
		return false
	}

	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formatName(name string) string {
	if name == "" {
		return ""
	}
	runes := []rune(strings.ToLower(name))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func itoa(value int) string {
	return fmt.Sprintf("%d", value)
}

func moveHorizontal(player *rl.Rectangle, movement float32, blocks []mapBlock) {
	player.X += movement
	for _, block := range blocks {
		if overlapsX(*player, block.rect) && touchesBlockSide(*player, block.rect) {
			if movement > 0 {
				player.X = block.rect.X - player.Width
			} else if movement < 0 {
				player.X = block.rect.X + block.rect.Width
			}
		}
	}
}

func touchesBlockSide(player rl.Rectangle, block rl.Rectangle) bool {
	return player.Y < block.Y && player.Y+player.Height > block.Y+4
}

func overlapsX(first rl.Rectangle, second rl.Rectangle) bool {
	return first.X < second.X+second.Width && first.X+first.Width > second.X
}