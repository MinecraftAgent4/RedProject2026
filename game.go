package main

import rl "github.com/gen2brain/raylib-go/raylib"

type mapBlock struct {
	rect  rl.Rectangle
	color rl.Color
}

func game() {
	rl.InitWindow(960, 540, "RedProject")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	blocks := []mapBlock{
		{rect: rl.Rectangle{X: 0, Y: 450, Width: 960, Height: 90}, color: rl.DarkGreen},
		{rect: rl.Rectangle{X: 80, Y: 360, Width: 180, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 350, Y: 290, Width: 180, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 630, Y: 360, Width: 180, Height: 24}, color: rl.Green},
		{rect: rl.Rectangle{X: 820, Y: 220, Width: 100, Height: 24}, color: rl.Green},
	}
	player := rl.Rectangle{X: 40, Y: 408, Width: 32, Height: 42}
	playerSpeed := float32(220)
	velocityY := float32(0)
	onGround := true

	for !rl.WindowShouldClose() {
		delta := rl.GetFrameTime()
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

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(28, 37, 61, 255))
		for _, block := range blocks {
			rl.DrawRectangleRec(block.rect, block.color)
			rl.DrawRectangleLinesEx(block.rect, 2, rl.DarkGray)
		}
		rl.DrawRectangleRec(player, rl.SkyBlue)
		rl.EndDrawing()
	}
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
