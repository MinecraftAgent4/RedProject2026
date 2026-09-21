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

func main() {
	rl.InitWindow(960, 540, "RedProject - Platformer Prototype")
	rl.SetTargetFPS(60)

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
	camera := rl.Camera2D{Offset: rl.Vector2{X: 320, Y: 0}, Target: rl.Vector2{X: 320, Y: 0}, Zoom: 1}

	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()
		if delta > 0.05 {
			delta = 0.05
		}

		if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
			player.X += 220 * delta
		}
		if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
			player.X -= 220 * delta
		}
		if (rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyW)) && onGround {
			velocity.Y = -430
			onGround = false
		}

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

		for index := range coins {
			coinRect := rl.Rectangle{X: coins[index].position.X - 10, Y: coins[index].position.Y - 10, Width: 20, Height: 20}
			if !coins[index].collected && rl.CheckCollisionRecs(player, coinRect) {
				coins[index].collected = true
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
		rl.DrawRectangleRec(goal, rl.Red)
		rl.DrawRectangleRec(player, rl.SkyBlue)
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
		if rl.CheckCollisionRecs(player, goal) {
			rl.DrawText("GOAL REACHED!", 360, 110, 28, rl.Lime)
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
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
