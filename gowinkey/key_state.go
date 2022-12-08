package gowinkey

const keyDown = 0x8000

// KeyState represents the state of some key.
type KeyState uint

const (
	// KeyUp represents a key being up.
	KeyUp KeyState = iota
	// KeyDown represents a key being down.
	KeyDown
)

func getKeyState(key int) KeyState {
	state := _getKeyState(key)
	if state&keyDown == keyDown {
		return KeyDown
	}
	return KeyUp
}
