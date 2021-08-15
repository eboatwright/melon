package melon


import (
	"github.com/gonutz/prototype/draw"
)


type function func()

// This is the template for all your games in Melon!
// You send a pointer to one of these, will all the data filled in
// The "Update" and "Draw" variables are just functions :)
type Game struct {
	// The title that will display on the window
	Title    string
	// Update is called 60fps and is before Draw
	Update   function
	// Draw is called 60fps and is after Update
	Draw     function
}


const (
	pIXEL_SIZE = 3
)

var (
	game = &Game {
		Title:    "empty melon project",
		Update:   func() {
		},
		Draw:     func() {
		},
	}
	window draw.Window
)


func render(w draw.Window) {
	window = w
	updateInput()
	game.Update()
	game.Draw()
}

// Call this to start Melon! :D
func Start(g *Game) {
	if g != nil { game = g }
	draw.RunWindow(game.Title, 320 * pIXEL_SIZE, 200 * pIXEL_SIZE, render)
}