package gowinkey

// Predicate represents a filter for key events.
type Predicate func(event KeyEvent) bool

// Selective returns a Predicate allowing only
// events for the given virtual keys to pass.
func Selective(keys ...VirtualKey) Predicate {
	set := NewKeySet(keys...)
	return func(event KeyEvent) bool {
		return set.Contains(event.VirtualKey)
	}
}

// KeyboardOnly is a Predicate allowing only
// events for keys on the keyboard to pass.
func KeyboardOnly(event KeyEvent) bool {
	switch event.VirtualKey {
	case VK_LBUTTON, VK_RBUTTON, VK_MBUTTON, VK_XBUTTON1, VK_XBUTTON2:
		return false
	}
	return true
}

// PressedOnly is a Predicate allowing
// only events for key presses to pass.
func PressedOnly(event KeyEvent) bool {
	return event.State == KeyDown
}

// ReleasedOnly is a Predicate allowing
// only events for key releases to pass.
func ReleasedOnly(event KeyEvent) bool {
	return event.State == KeyUp
}

