package gowinkey

import (
	"math/bits"
	"strconv"
	"strings"
)

// Modifiers represents modifiers that are pressed
// alongside some virtual key. See the flags below
// for more info.
type Modifiers uint

// These flags define which modifiers are being pressed alongside
// some virtual key. Bits are or'ed to build a modifier field. See
// KeyEvent.Modifiers for more info.
//
// Note that these flags do not differentiate between the
// 'left' and 'right' versions of keys. So, for instance,
// pressing either 'left shift' or 'right shift' will simply
// result in ModifierShift being used to build the modifier
// field.
const (
	// ModifierShift identifies any 'shift' modifier.
	ModifierShift Modifiers = 1 << iota
	// ModifierMenu identifies any 'alt' modifier.
	ModifierMenu
	// ModifierControl identifies any 'ctrl' modifier.
	ModifierControl
)

// modifiersToStr maps each of the known
// modifiers to its string representation.
var modifiersToStr = map[Modifiers]string{
	ModifierShift:   "Shift",
	ModifierMenu:    "Alt",
	ModifierControl: "Ctrl",
}

// HasModifiers reports whether m contains the given modifiers.
func (m Modifiers) HasModifiers(modifiers Modifiers) bool {
	return m&modifiers != 0
}

// RemoveModifiers returns a copy of m with the given modifiers removed.
func (m Modifiers) RemoveModifiers(modifiers Modifiers) Modifiers {
	return m & ^modifiers
}

// String returns the string representation of m.
func (m Modifiers) String() string {
	var mods []string
	for mod, s := range modifiersToStr {
		if m.HasModifiers(mod) {
			mods = append(mods, s)
		}
	}
	if len(mods) != bits.OnesCount(uint(m)) {
		panic("unknown modifiers: " + m.toBinary())
	}
	return strings.Join(mods, " + ")
}

// toBinary returns the string representation of m in base 2.
func (m Modifiers) toBinary() string {
	return strconv.FormatInt(int64(m), 2)
}
