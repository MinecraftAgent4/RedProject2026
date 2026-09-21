package main

import rl "github.com/gen2brain/raylib-go/raylib"

func game() {
	rl.InitWindow(960, 540, "RedProject")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkBlue)
		rl.DrawText("RedProject", 380, 240, 40, rl.RayWhite)
		rl.EndDrawing()
	}
}
