package gowinkey

import (
	"time"
)

// listener listens for global key events.
type listener struct {
	modifiers   Modifiers
	pressedKeys KeySet
	predicates  []Predicate
}

func newListener(predicates ...Predicate) *listener {
	return &listener{
		pressedKeys: make(KeySet),
		predicates:  predicates,
	}
}

// Listen listens for global key events, sending them on the
// events channel.
// Listen halts execution and closes the events channel as
// soon as stopFn is called.
// Listen does not block.
// The given predicates act as a filter. Only events matching
// all predicates will be sent on the events channel.
func Listen(predicates ...Predicate) (events <-chan KeyEvent) {
	return newListener(predicates...).listen()
}

// listen listens for global key events, sending them on the
// returned channel. listen halts execution and closes the
// returned channel as soon as the returned function is called.
// listen does not block.
func (l *listener) listen() <-chan KeyEvent {
	events, stopChan := make(chan KeyEvent), make(chan bool)

	go func() {
		l.swallowQueuedStates()
		l.doListen(events, stopChan)
	}()

	return events
}

// swallowQueuedStates drains the message queue so that the
// listener does not catch any events that were issue before
// listener.listen was called.
func (l listener) swallowQueuedStates() {
	for i := 0; i < 256; i++ {
		_getKeyState(i)
	}
}

// doListen listens for global key events,
// sending them on the events channel.
func (l *listener) doListen(events chan KeyEvent, stopChan <-chan bool) {
Outer:
	for {
		select {
		case <-stopChan:
			break Outer
		default:
			time.Sleep(10 * time.Millisecond)
			l.listenOnce(events)
		}
	}
}

// listenOnce listens for the state of each of the 254 known
// virtual keys and sends according key events on the events
// channel.
func (l *listener) listenOnce(events chan KeyEvent) {
	for i := 0; i < 255; i++ {
		key := VirtualKey(i)
		if l.isDuplicateModifier(key) {
			continue
		}

		state := getKeyState(i)
		event := KeyEvent{
			VirtualKey: key,
			State:      state,
		}

		if state == KeyDown {
			if !l.pressedKeys.Contains(key) {
				l.pressedKeys.Add(key)
				l.processModifier(key, KeyDown)

				l.applyModifiers(&event)
				event.PressedKeys = l.pressedKeys
				if l.satisfiesPredicates(event) {
					events <- event
				}
			}
		} else {
			if l.pressedKeys.Contains(key) {
				l.pressedKeys.Delete(key)
				l.processModifier(key, KeyUp)

				l.applyModifiers(&event)
				event.PressedKeys = l.pressedKeys
				if l.satisfiesPredicates(event) {
					events <- event
				}
			}
		}
	}
}

// isDuplicateModifier reports whether the given
// virtual key represents a duplicate modifier.
//
// This is needed, because the Windows API fires two events
// when a modifier key is pressed - one for the specific key
// (say, VK_LSHIFT) and one for the "raw" modifier
// (say, VK_SHIFT).
func (l listener) isDuplicateModifier(key VirtualKey) bool {
	return key == VK_SHIFT || key == VK_CONTROL || key == VK_MENU
}

// processModifier extracts modifier information from the given
// given virtual key and updates the listener.modifiers accordingly.
func (l *listener) processModifier(key VirtualKey, state KeyState) {
	mod := l.keyToModifier(key)
	if state == KeyDown {
		l.modifiers |= mod
	} else {
		if !l.modifierCounterpartPressed(key) {
			l.modifiers = l.modifiers.RemoveModifiers(mod)
		}
	}
}

// keyToModifier returns the modifier associated with
// the given virtual key. If the key does not represent
// any modifier, keyToModifier returns 0.
func (l listener) keyToModifier(key VirtualKey) Modifiers {
	switch key {
	case VK_SHIFT, VK_LSHIFT, VK_RSHIFT:
		return ModifierShift
	case VK_CONTROL, VK_LCONTROL, VK_RCONTROL:
		return ModifierControl
	case VK_MENU, VK_LMENU, VK_RMENU:
		return ModifierMenu
	}
	return 0
}

// modifierCounterpartPressed reports whether the "counterpart" of
// the given modifier key is pressed, where by counterpart we mean
// the left version for right modifier keys and vice versa.
//
// If key does not represent a modifier key,
// modifierCounterpartPressed returns false.
func (l listener) modifierCounterpartPressed(key VirtualKey) bool {
	switch key {
	case VK_LSHIFT:
		return l.pressedKeys.Contains(VK_RSHIFT)
	case VK_RSHIFT:
		return l.pressedKeys.Contains(VK_LSHIFT)
	case VK_LCONTROL:
		return l.pressedKeys.Contains(VK_RCONTROL)
	case VK_RCONTROL:
		return l.pressedKeys.Contains(VK_LCONTROL)
	case VK_LMENU:
		return l.pressedKeys.Contains(VK_RMENU)
	case VK_RMENU:
		return l.pressedKeys.Contains(VK_LMENU)
	}
	return false
}

// applyModifiers applies modifiers for
// the currently pressed keys to the event.
func (l listener) applyModifiers(event *KeyEvent) {
	eventMod := l.keyToModifier(event.VirtualKey)
	event.Modifiers = l.modifiers.RemoveModifiers(eventMod)
}

// satisfiesPredicates reports whether the given
// key event satisfies all listener.predicates.
func (l listener) satisfiesPredicates(event KeyEvent) bool {
	for _, p := range l.predicates {
		if !p(event) {
			return false
		}
	}
	return true
}
