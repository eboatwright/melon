package main


import (
	"github.com/gonutz/prototype/draw"
)


// Draws image from spritesheet.png
func DrawImage(x, y, width, height, srcX, srcY int) {
	window.DrawImageFilePart(
		"spritesheet.png",
		srcX, srcY, width, height,
		x * pIXEL_SIZE, y * pIXEL_SIZE, width * pIXEL_SIZE, height * pIXEL_SIZE,
		0,
	)
}

// Draws text
func DrawText(text string, x, y int, scale float32, color draw.Color) {
	window.DrawScaledText(text, x * pIXEL_SIZE, y * pIXEL_SIZE, scale * float32(pIXEL_SIZE), color)
}

// Draws rect
func DrawRect(x, y, width, height int, color draw.Color) {
	window.FillRect(x * pIXEL_SIZE, y * pIXEL_SIZE, width * pIXEL_SIZE, height * pIXEL_SIZE, color)
}

// Basically just draws rect with requested color
func Clear(color draw.Color) {
	DrawRect(0, 0, 320, 200, color)
}

// Returns a draw.Color {}
func Color(r, g, b, a float32) draw.Color {
	return draw.Color { r, g, b, a }
}

var (
	WHITE = Color(255, 255, 255, 255)
	BLACK = Color(0, 0, 0, 255)
)