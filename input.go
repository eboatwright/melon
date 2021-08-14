package main


import (
	"github.com/gonutz/prototype/draw"
)


var (
	keys             [110]bool
	keysLast         [110]bool
	mouseButtons     [3]bool
	mouseButtonsLast [3]bool
)


func updateInput() {
	keysLast = keys
	for i := 0; i < len(keys); i++ {
		keys[i] = window.IsKeyDown(draw.Key(i))
	}

	mouseButtonsLast = mouseButtons
	for i := 0; i < len(mouseButtons); i++ {
		mouseButtons[i] = window.IsMouseDown(draw.MouseButton(i))
	}
}

// Returns true or false when key is being held
func IsKeyDown(key draw.Key) bool {
	return keys[int(key)]
}

// Returns true or false if key was just pressed this frame
func IsKeyJustDown(key draw.Key) bool {
	return keys[int(key)] && !keysLast[int(key)]
}

// Returns true or false if key was just released this frame
func IsKeyJustReleased(key draw.Key) bool {
	return !keys[int(key)] && keysLast[int(key)]
}

// Returns true or false when mouse button is being held
func IsMouseButtonDown(button draw.MouseButton) bool {
	return mouseButtons[int(button)]
}

// Returns true or false if mouse button was just pressed this frame
func IsMouseButtonJustDown(button draw.MouseButton) bool {
	return mouseButtons[int(button)] && !mouseButtonsLast[int(button)]
}

// Returns true or false if mouse button was just released this frame
func IsMouseButtonJustReleased(button draw.MouseButton) bool {
	return !mouseButtons[int(button)] && mouseButtonsLast[int(button)]
}

// Get mouse position
func GetMousePosition() (x, y int) {
	return window.MousePosition()
}

// Returns draw.Key
func Key(key int) draw.Key {
	return draw.Key(key)
}