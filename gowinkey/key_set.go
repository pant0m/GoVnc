package gowinkey

// KeySet represents a set of virtual keys.
type KeySet map[VirtualKey]struct{}

// NewKeySet constructs a new KeySet from the given virtual keys.
func NewKeySet(keys ...VirtualKey) KeySet {
	set := make(KeySet)
	set.Add(keys...)
	return set
}

// Slice returns a slice containing the virtual keys in s.
func (s KeySet) Slice() []VirtualKey {
	slice, i := make([]VirtualKey, len(s)), 0
	for key := range s {
		slice[i] = key
		i++
	}
	return slice
}

// Add adds the given virtual keys to s.
func (s KeySet) Add(keys ...VirtualKey) {
	for _, key := range keys {
		s[key] = struct{}{}
	}
}

// Delete deletes the given virtual keys from s.
func (s KeySet) Delete(keys ...VirtualKey) {
	for _, key := range keys {
		delete(s, key)
	}
}

// Contains reports whether s contains the given virtual key.
func (s KeySet) Contains(key VirtualKey) bool {
	_, ok := s[key]
	return ok
}

// ContainsAll reports whether s contains all of the given virtual keys.
func (s KeySet) ContainsAll(keys ...VirtualKey) bool {
	for _, key := range keys {
		if !s.Contains(key) {
			return false
		}
	}
	return true
}

// ContainsAny reports whether s contains any of the given virtual keys.
func (s KeySet) ContainsAny(keys ...VirtualKey) bool {
	for _, key := range keys {
		if s.Contains(key) {
			return true
		}
	}
	return false
}
