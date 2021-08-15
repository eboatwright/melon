package main


var (
	playerPos      = VZERO
	playerVel      = VZERO
	playerFriction = 0.85
	playerSpeed    = 0.7
)


func main() {
	Start(&Game {
		Title:     "Melon-GO - eboatwright",
		Update:    func() {
			input := Vector2 {
				float64(GetInputAxis([]Key { KeyA }, []Key { KeyD })),
				float64(GetInputAxis([]Key { KeyW }, []Key { KeyS })),
			}
			input.Normalize()
			playerVel.Add(Vector2Multiply(input, playerSpeed))
			playerVel.Multiply(playerFriction)
			playerPos.Add(Vector2Floor(playerVel))	
		},
		Draw:      func() {
			Clear(WHITE)
			DrawImage(int(playerPos.X), int(playerPos.Y), 9, 14, 0, 0)
		},
	})
}