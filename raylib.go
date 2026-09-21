package main

import rl "github.com/gen2brain/raylib-go/raylib"

type platform struct {
	rect  rl.Rectangle
	color rl.Color
}

type coin struct {
	position  rl.Vector2
	collected bool
}

type enemy struct {
	name     string
	position rl.Vector2
	hp       int
	maxHP    int
	attack   int
	defeated bool
}

type combatState struct {
	enemyIndex  int
	playerHP    int
	maxPlayerHP int
	defending   bool
	finished    bool
	message     string
}

type animation struct {
	texture   rl.Texture2D
	frame     int
	timer     float32
	frameTime float32
}

func main() {
	rl.InitWindow(960, 540, "RedProject - Platformer Prototype")
	rl.SetTargetFPS(60)
	playerAnimation := animation{texture: rl.LoadTexture("asset/Free 3 Cyberpunk Sprites Pixel Art(2)/3 Cyborg/Cyborg_run.png"), frameTime: 0.1}
	enemyAnimation := animation{texture: rl.LoadTexture("asset/Free 3 Cyberpunk Sprites Pixel Art(2)/1 Biker/Biker_run.png"), frameTime: 0.14}

	player := rl.Rectangle{X: 80, Y: 350, Width: 30, Height: 42}
	velocity := rl.Vector2{}
	onGround := false
	coins := []coin{
		{position: rl.Vector2{X: 310, Y: 355}},
		{position: rl.Vector2{X: 620, Y: 265}},
		{position: rl.Vector2{X: 910, Y: 175}},
		{position: rl.Vector2{X: 1260, Y: 315}},
	}
	platforms := []platform{
		{rect: rl.Rectangle{X: 0, Y: 450, Width: 1500, Height: 90}, color: rl.DarkGreen},
		{rect: rl.Rectangle{X: 230, Y: 400, Width: 160, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 500, Y: 330, Width: 180, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 780, Y: 240, Width: 180, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 1080, Y: 380, Width: 190, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 1370, Y: 300, Width: 130, Height: 24}, color: rl.Green},
	}
	goal := rl.Rectangle{X: 1430, Y: 258, Width: 22, Height: 42}
	enemies := []enemy{
		{name: "Rival de ruelle", position: rl.Vector2{X: 430, Y: 408}, hp: 30, maxHP: 30, attack: 7},
		{name: "Drone hostile", position: rl.Vector2{X: 850, Y: 198}, hp: 42, maxHP: 42, attack: 9},
		{name: "Garde du boss", position: rl.Vector2{X: 1160, Y: 348}, hp: 55, maxHP: 55, attack: 12},
	}
	combat := (*combatState)(nil)
	camera := rl.Camera2D{Offset: rl.Vector2{X: 320, Y: 0}, Target: rl.Vector2{X: 320, Y: 0}, Zoom: 1}

	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()
		if delta > 0.05 {
			delta = 0.05
		}

		if combat == nil && (rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD)) {
			player.X += 220 * delta
		}
		if combat == nil && (rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA)) {
			player.X -= 220 * delta
		}
		if combat == nil {
			updateAnimation(&playerAnimation, delta)
		}
		updateAnimation(&enemyAnimation, delta)
		if combat == nil && (rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyW)) && onGround {
			velocity.Y = -430
			onGround = false
		}

		if combat == nil {
			velocity.Y += 1000 * delta
			player.Y += velocity.Y * delta
			onGround = false
			for _, currentPlatform := range platforms {
				if velocity.Y >= 0 && rl.CheckCollisionRecs(player, currentPlatform.rect) {
					previousBottom := player.Y - velocity.Y*delta + player.Height
					if previousBottom <= currentPlatform.rect.Y+4 {
						player.Y = currentPlatform.rect.Y - player.Height
						velocity.Y = 0
						onGround = true
					}
				}
			}

			if player.Y > 650 {
				player.X = 80
				player.Y = 350
				velocity = rl.Vector2{}
			}
			if player.X < 0 {
				player.X = 0
			}
			if player.X > 1470 {
				player.X = 1470
			}
		}

		for index := range coins {
			coinRect := rl.Rectangle{X: coins[index].position.X - 10, Y: coins[index].position.Y - 10, Width: 20, Height: 20}
			if !coins[index].collected && rl.CheckCollisionRecs(player, coinRect) {
				coins[index].collected = true
			}
		}
		if combat == nil {
			for index := range enemies {
				if !enemies[index].defeated && rl.CheckCollisionRecs(player, enemyRectangle(enemies[index])) {
					combat = &combatState{enemyIndex: index, playerHP: 100, maxPlayerHP: 100, message: "Choisis une action."}
					break
				}
			}
		} else {
			updateCombat(combat, &enemies[combat.enemyIndex])
			if combat.playerHP <= 0 {
				player.X = 80
				player.Y = 350
				velocity = rl.Vector2{}
				combat = nil
			} else if combat.finished || enemies[combat.enemyIndex].defeated {
				combat = nil
			}
		}

		camera.Target.X = player.X
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(22, 31, 52, 255))
		rl.BeginMode2D(camera)
		rl.DrawCircle(1350, 90, 42, rl.Gold)
		for _, currentPlatform := range platforms {
			rl.DrawRectangleRec(currentPlatform.rect, currentPlatform.color)
		}
		for _, currentCoin := range coins {
			if !currentCoin.collected {
				rl.DrawCircleV(currentCoin.position, 10, rl.Gold)
				rl.DrawCircleLines(int32(currentCoin.position.X), int32(currentCoin.position.Y), 10, rl.Orange)
			}
		}
		for _, currentEnemy := range enemies {
			if !currentEnemy.defeated {
				drawSprite(enemyAnimation, currentEnemy.position, 1.25)
			}
		}
		rl.DrawRectangleRec(goal, rl.Red)
		drawSprite(playerAnimation, rl.Vector2{X: player.X + player.Width/2, Y: player.Y + player.Height}, 1.35)
		rl.EndMode2D()

		collected := 0
		for _, currentCoin := range coins {
			if currentCoin.collected {
				collected++
			}
		}
		rl.DrawText("REDPROJECT PLATFORMER", 24, 20, 24, rl.RayWhite)
		rl.DrawText("A/D or arrows: move   SPACE: jump", 24, 52, 18, rl.LightGray)
		rl.DrawText("Coins: "+itoa(collected)+"/"+itoa(len(coins)), 24, 82, 18, rl.Gold)
		rl.DrawText("Enemies: fight on contact", 24, 110, 18, rl.Orange)
		if rl.CheckCollisionRecs(player, goal) {
			rl.DrawText("GOAL REACHED!", 360, 110, 28, rl.Lime)
		}
		if combat != nil {
			drawCombat(combat, enemies[combat.enemyIndex])
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
	rl.UnloadTexture(playerAnimation.texture)
	rl.UnloadTexture(enemyAnimation.texture)
}

func updateAnimation(current *animation, delta float32) {
	current.timer += delta
	if current.timer >= current.frameTime {
		current.timer -= current.frameTime
		current.frame = (current.frame + 1) % 6
	}
}

func drawSprite(current animation, position rl.Vector2, scale float32) {
	frameWidth := float32(current.texture.Width / 6)
	source := rl.Rectangle{X: float32(current.frame) * frameWidth, Y: 0, Width: frameWidth, Height: float32(current.texture.Height)}
	destination := rl.Rectangle{X: position.X - frameWidth*scale/2, Y: position.Y - float32(current.texture.Height)*scale, Width: frameWidth * scale, Height: float32(current.texture.Height) * scale}
	rl.DrawTexturePro(current.texture, source, destination, rl.Vector2{}, 0, rl.White)
}

func enemyRectangle(currentEnemy enemy) rl.Rectangle {
	return rl.Rectangle{X: currentEnemy.position.X - 16, Y: currentEnemy.position.Y - 32, Width: 32, Height: 42}
}

func updateCombat(combat *combatState, currentEnemy *enemy) {
	damage := 0
	mouse := rl.GetMousePosition()
	attackButton := rl.Rectangle{X: 520, Y: 330, Width: 135, Height: 42}
	defendButton := rl.Rectangle{X: 670, Y: 330, Width: 135, Height: 42}
	fleeButton := rl.Rectangle{X: 670, Y: 382, Width: 135, Height: 42}
	clicked := rl.IsMouseButtonPressed(rl.MouseButtonLeft)
	if rl.IsKeyPressed(rl.KeyOne) || (clicked && rl.CheckCollisionPointRec(mouse, attackButton)) {
		damage = 12
		combat.message = "Tu attaques et infliges " + itoa(damage) + " degats."
	} else if rl.IsKeyPressed(rl.KeyTwo) || (clicked && rl.CheckCollisionPointRec(mouse, defendButton)) {
		combat.defending = true
		combat.message = "Tu te prepares a bloquer la prochaine attaque."
	} else if rl.IsKeyPressed(rl.KeyThree) || (clicked && rl.CheckCollisionPointRec(mouse, fleeButton)) {
		combat.message = "Tu prends la fuite."
		combat.finished = true
		return
	} else {
		return
	}

	currentEnemy.hp -= damage
	if currentEnemy.hp <= 0 {
		currentEnemy.defeated = true
		combat.message = "Victoire !"
		return
	}

	enemyDamage := currentEnemy.attack
	if combat.defending {
		enemyDamage /= 2
		combat.defending = false
	}
	combat.playerHP -= enemyDamage
	combat.message += " L'ennemi riposte : " + itoa(enemyDamage) + " degats."
}

func drawCombat(combat *combatState, currentEnemy enemy) {
	panel := rl.Rectangle{X: 80, Y: 70, Width: 800, Height: 430}
	rl.DrawRectangleRec(panel, rl.NewColor(238, 232, 201, 255))
	rl.DrawRectangleLinesEx(panel, 5, rl.NewColor(51, 66, 67, 255))
	rl.DrawRectangle(80, 300, 800, 200, rl.NewColor(67, 105, 98, 255))
	rl.DrawLine(80, 300, 880, 300, rl.NewColor(51, 66, 67, 255))

	rl.DrawText("UN ADVERSAIRE APPARAIT !", 115, 92, 22, rl.NewColor(51, 66, 67, 255))
	rl.DrawText(currentEnemy.name, 575, 112, 20, rl.NewColor(51, 66, 67, 255))
	drawHPBar(575, 142, currentEnemy.hp, currentEnemy.maxHP)
	drawBattleEnemy(currentEnemy)
	drawBattlePlayer()

	rl.DrawRectangle(105, 330, 365, 135, rl.NewColor(247, 244, 222, 255))
	rl.DrawRectangleLinesEx(rl.Rectangle{X: 105, Y: 330, Width: 365, Height: 135}, 3, rl.NewColor(51, 66, 67, 255))
	rl.DrawText("JOUEUR", 130, 350, 18, rl.NewColor(51, 66, 67, 255))
	rl.DrawText(combat.message, 130, 380, 16, rl.NewColor(51, 66, 67, 255))
	rl.DrawText("PV "+itoa(combat.playerHP)+"/"+itoa(combat.maxPlayerHP), 130, 425, 18, rl.NewColor(51, 66, 67, 255))
	drawHPBar(245, 428, combat.playerHP, combat.maxPlayerHP)

	drawCombatButton(rl.Rectangle{X: 520, Y: 330, Width: 135, Height: 42}, "ATTAQUER", rl.NewColor(199, 76, 70, 255))
	drawCombatButton(rl.Rectangle{X: 670, Y: 330, Width: 135, Height: 42}, "DEFENDRE", rl.NewColor(66, 126, 157, 255))
	drawCombatButton(rl.Rectangle{X: 520, Y: 382, Width: 135, Height: 42}, "OBJET", rl.NewColor(193, 145, 64, 255))
	drawCombatButton(rl.Rectangle{X: 670, Y: 382, Width: 135, Height: 42}, "FUIR", rl.NewColor(92, 95, 100, 255))
	rl.DrawText("Clique une action", 555, 445, 16, rl.RayWhite)
}

func drawHPBar(x int32, y int32, hp int, maxHP int) {
	rl.DrawRectangle(x, y, 210, 15, rl.NewColor(51, 66, 67, 255))
	if hp > 0 {
		width := int32(210 * hp / maxHP)
		color := rl.Green
		if hp*3 < maxHP {
			color = rl.Red
		} else if hp*2 < maxHP {
			color = rl.Gold
		}
		rl.DrawRectangle(x, y, width, 15, color)
	}
}

func drawBattleEnemy(currentEnemy enemy) {
	rl.DrawEllipse(690, 230, 85, 20, rl.NewColor(51, 66, 67, 90))
	rl.DrawCircle(690, 190, 42, rl.Maroon)
	rl.DrawCircle(675, 180, 8, rl.RayWhite)
	rl.DrawCircle(705, 180, 8, rl.RayWhite)
	rl.DrawCircle(675, 180, 3, rl.Black)
	rl.DrawCircle(705, 180, 3, rl.Black)
	rl.DrawText("ENNEMI", 650, 245, 16, rl.NewColor(51, 66, 67, 255))
}

func drawBattlePlayer() {
	rl.DrawEllipse(280, 245, 95, 20, rl.NewColor(51, 66, 67, 90))
	rl.DrawCircle(280, 200, 40, rl.SkyBlue)
	rl.DrawRectangle(252, 200, 56, 52, rl.Blue)
	rl.DrawCircle(268, 190, 7, rl.RayWhite)
	rl.DrawCircle(292, 190, 7, rl.RayWhite)
	rl.DrawCircle(268, 190, 3, rl.Black)
	rl.DrawCircle(292, 190, 3, rl.Black)
	rl.DrawText("TOI", 268, 260, 16, rl.NewColor(51, 66, 67, 255))
}

func drawCombatButton(rect rl.Rectangle, label string, color rl.Color) {
	buttonColor := color
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) {
		buttonColor = rl.ColorBrightness(color, 0.25)
	}
	rl.DrawRectangleRec(rect, buttonColor)
	rl.DrawRectangleLinesEx(rect, 2, rl.RayWhite)
	rl.DrawText(label, int32(rect.X+22), int32(rect.Y+14), 18, rl.RayWhite)
}

func itoa(value int) string {
	if value == 0 {
		return "0"
	}
	result := ""
	for value > 0 {
		result = string(rune('0'+value%10)) + result
		value /= 10
	}
	return result
}
