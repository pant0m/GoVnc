package gowinkey

import (
	"fmt"
	"syscall"
)

var user32 = syscall.MustLoadDLL("user32")
var keystate = user32.MustFindProc("GetAsyncKeyState")
var mapVirtualKey = user32.MustFindProc("MapVirtualKeyA")

// virtualKeyToString returns the string representation of vk.
//
// See https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapvirtualkeya
// for more details.
func virtualKeyToString(vk VirtualKey) string {
	char, _, _ := syscall.Syscall(mapVirtualKey.Addr(), 2, uintptr(vk), 2, 0)
	return fmt.Sprint(char)
}

// _getKeyState determines whether the given key is up or down
// at the time this function is called, and whether the key
// was pressed after a previous call to _getKeyState.
//
// See https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getasynckeystate
// for more details.
func _getKeyState(key int) int32 {
	state, _, _ := syscall.Syscall(keystate.Addr(), 1, uintptr(key), 0, 0)
	return int32(state)
}
