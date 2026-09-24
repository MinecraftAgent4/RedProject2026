package main

import (
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Animation gère une spritesheet horizontale découpée en frames.
type Animation struct {
	texture     rl.Texture2D
	frameWidth  int32
	frameHeight int32
	frameCount  int
	current     int
	timer       float32
	frameTime   float32
	loop        bool
}

func NewAnimation(path string, frameWidth, frameHeight, frameCount int, frameTime float32) *Animation {
	return &Animation{
		texture:     rl.LoadTexture(path),
		frameWidth:  int32(frameWidth),
		frameHeight: int32(frameHeight),
		frameCount:  frameCount,
		frameTime:   frameTime,
		loop:        true,
	}
}

func (a *Animation) Update(delta float32) {
	if a == nil || a.frameCount <= 1 {
		return
	}

	if a.frameTime <= 0 {
		a.frameTime = 0.1
	}

	a.timer += delta
	for a.timer >= a.frameTime {
		a.timer -= a.frameTime
		a.current++
		if a.current >= a.frameCount {
			if a.loop {
				a.current = 0
			} else {
				a.current = a.frameCount - 1
			}
		}
	}
}

func (a *Animation) Draw(x, y, width, height float32) {
	if a == nil || a.texture.ID == 0 {
		return
	}

	source := rl.Rectangle{
		X:      float32(a.current) * float32(a.frameWidth),
		Y:      0,
		Width:  float32(a.frameWidth),
		Height: float32(a.frameHeight),
	}
	destination := rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	rl.DrawTexturePro(a.texture, source, destination, rl.Vector2{}, 0, rl.White)
}

func (a *Animation) Unload() {
	if a != nil && a.texture.ID != 0 {
		rl.UnloadTexture(a.texture)
		a.texture = rl.Texture2D{}
	}
}

// initPNJAnimations charge les sprites Idle des trois PNJ de la carte.
func initPNJAnimations() map[string]*Animation {
	base := "asset valide"
	return map[string]*Animation{
		"MARCHAND":       NewAnimation(filepath.Join(base, "marchand", "Idle.png"), 48, 48, 4, 0.25),
		"CHARCUDOC":      NewAnimation(filepath.Join(base, "charcudoc", "Idle.png"), 48, 48, 6, 0.25),
		"GUERRE DE GANG": NewAnimation(filepath.Join(base, "guerre de gang", "Idle.png"), 72, 72, 4, 0.25),
	}
}

func unloadPNJAnimations(anims map[string]*Animation) {
	for _, animation := range anims {
		animation.Unload()
	}
}
