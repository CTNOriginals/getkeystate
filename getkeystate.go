//go:build windows

package getkeystate

// #include "./main.h"
import "C"

// Get the current state of key.
// key accepts the raw key code interger for VK_BUTTONS listed here: https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes.
//
// Returns a key state value that is a 2 bit int.
// The left most bit represents wether it is currently held down (active).
// The right most bit represents wether the key is currently toggled.
//
// All keys will return a toggled on/off value, but this functionality is mostly useful for keys like CapsLock.
func GetKeyState(key int) int {
	var state = C.C_GetKeyState(C.int(key))

	switch state {
	case 1: // Toggled
		return 0b01
	case -128: // Active >| not toggled
		return 0b10
	case -127: // Active >| toggled
		return 0b11
	default: // idle
		return 0b00
	}
}

// Takes a keyState that is returned by GetKeyState().
//
// Returns wether this key state signals if the key is toggled.
func IsToggled(keyState int) bool {
	return (keyState & 1) != 0
}

// Takes a keyState that is returned by GetKeyState().
//
// Returns wether this key state signals if the key is being held down.
func IsActive(keyState int) bool {
	return (keyState & 2) != 0
}
