package main


import (
	"github.com/gonutz/prototype/draw"
	"github.com/eboatwright/melon"
)


var (
	playerPos      = melon.VZERO
	playerVel      = melon.VZERO
	playerFriction = 0.85
	playerSpeed    = 0.7
	playerDir      = 1
)


func main() {
	melon.Start(&melon.Game {
		Title:     "Top Down Example - Melon",
		Update:    func() {
			input := melon.Vector2 {
				float64(melon.GetInputAxis([]draw.Key { draw.KeyA }, []draw.Key { draw.KeyD })),
				float64(melon.GetInputAxis([]draw.Key { draw.KeyW }, []draw.Key { draw.KeyS })),
			}
			if input.X != 0 {
				playerDir = int(input.X)
			}
			input.Normalize()
			playerVel.Add(melon.Vector2Multiply(input, playerSpeed))
			playerVel.Multiply(playerFriction)
			playerPos.Add(melon.Vector2Floor(playerVel))
		},
		Draw:      func() {
			melon.Clear(draw.White)
			playerOffset := 0
			if playerDir == -1 {
				playerOffset = 9
			}
			melon.DrawImage("data/img/player.png", int(playerPos.X), int(playerPos.Y), 9, 14, playerOffset, 0, 9 * playerDir, 14)
		},
	})
}