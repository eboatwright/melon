package main


var x = 0
func main() {
	Start(&Game {
		Title:     "Melon-GO - eboatwright",
		Update:    func() {
			if IsKeyDown(KeySpace) {
				x++
			}
		},
		Draw:      func() {
			Clear(WHITE)
			DrawImage(x, 1, 9, 14, 0, 0)
		},
	})
}