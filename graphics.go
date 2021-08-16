package melon


import (
	"github.com/gonutz/prototype/draw"
)


// Draws image from spritesheet.png
func DrawImage(path string, destX, destY, destWidth, destHeight, srcX, srcY, srcWidth, srcHeight int) {
	window.DrawImageFilePart(
		path,
		srcX, srcY, srcWidth, srcHeight,
		destX * pIXEL_SIZE, destY * pIXEL_SIZE, destWidth * pIXEL_SIZE, destHeight * pIXEL_SIZE,
		0,
	)
}

// Draws text
func DrawText(text string, x, y int, scale float32, color draw.Color) {
	window.DrawScaledText(text, x * pIXEL_SIZE, y * pIXEL_SIZE, scale * float32(pIXEL_SIZE), color)
}

// Draws rect
func DrawRect(x, y, width, height int, color draw.Color, filled bool) {
	if filled {
		window.FillRect(x * pIXEL_SIZE, y * pIXEL_SIZE, width * pIXEL_SIZE, height * pIXEL_SIZE, color)
	} else {
		window.DrawRect(x * pIXEL_SIZE, y * pIXEL_SIZE, width * pIXEL_SIZE, height * pIXEL_SIZE, color)
	}
}

// Draws Ellipse either filled or outline
func DrawEllipse(x, y, width, height int, color draw.Color, filled bool) {
	if filled {
		window.FillEllipse(x, y, width, height, color)
	} else {
		window.DrawEllipse(x, y, width, height, color)
	}
}

// Draws line
func DrawLine(x1, y1, x2, y2 int, color draw.Color) {
	window.DrawLine(x1, y1, x2, y2, color)
}

// Basically just draws rect with requested color
func Clear(color draw.Color) {
	DrawRect(0, 0, 320, 200, color, true)
}