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
func IsKeyDown(key Key) bool {
	return keys[int(key)]
}

// Returns true or false if key was just pressed this frame
func IsKeyJustDown(key Key) bool {
	return keys[int(key)] && !keysLast[int(key)]
}

// Returns true or false if key was just released this frame
func IsKeyJustReleased(key Key) bool {
	return !keys[int(key)] && keysLast[int(key)]
}

// Returns true or false when mouse button is being held
func IsMouseButtonDown(button MouseButton) bool {
	return mouseButtons[int(button)]
}

// Returns true or false if mouse button was just pressed this frame
func IsMouseButtonJustDown(button MouseButton) bool {
	return mouseButtons[int(button)] && !mouseButtonsLast[int(button)]
}

// Returns true or false if mouse button was just released this frame
func IsMouseButtonJustReleased(button MouseButton) bool {
	return !mouseButtons[int(button)] && mouseButtonsLast[int(button)]
}

// Get mouse position
func GetMousePosition() (x, y int) {
	return window.MousePosition()
}


// I copied these from prototype, so you don't have to import 2 things in your files :)
type Key int
const (
	KeyA Key = 1 + iota
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyNum0
	KeyNum1
	KeyNum2
	KeyNum3
	KeyNum4
	KeyNum5
	KeyNum6
	KeyNum7
	KeyNum8
	KeyNum9
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyEnter
	KeyNumEnter
	KeyLeftControl
	KeyRightControl
	KeyLeftShift
	KeyRightShift
	KeyLeftAlt
	KeyRightAlt
	KeyLeft
	KeyRight
	KeyUp
	KeyDown
	KeyEscape
	KeySpace
	KeyBackspace
	KeyTab
	KeyHome
	KeyEnd
	KeyPageDown
	KeyPageUp
	KeyDelete
	KeyInsert
	KeyNumAdd
	KeyNumSubtract
	KeyNumMultiply
	KeyNumDivide
	KeyCapslock
	KeyPrint
	KeyPause
)

// I copied these from prototype, so you don't have to import 2 things in your files :)
type MouseButton int
const (
	LeftButton MouseButton = iota
	MiddleButton
	RightButton
)