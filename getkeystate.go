//go:build windows

package getkeystate

// #include <stdio.h>
import "C"

// #region KeyState

// A 2 bit int representing a key state.
//
// The right most bit represents wether it is currently held down (active).
// The left most bit represents wether the key is currently toggled.
type KeyState int

// Returns wether this key state signals if the key is being held down.
func (this KeyState) IsActive() bool {
	return (this & 2) != 0
}

// Returns wether this key state signals if the key is toggled.
func (this KeyState) IsToggled() bool {
	return (this & 1) != 0
}

//#endregion

// Get the current state of key.
// key accepts the raw key code interger for VK_BUTTONS listed here: https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes.
//
// Returns a key state value that is a 2 bit int.
// The right most bit represents wether it is currently held down (active).
// The left most bit represents wether the key is currently toggled.
//
// All keys will return a toggled on/off value, but this functionality is mostly useful for keys like CapsLock.
func GetKeyState(key int) KeyState {
	var state = C.C_GetKeyState(C.int(key))

	switch state {
	case 1: // Toggled
		return KeyState(1)
	case -128: // Active >| not toggled
		return KeyState(2)
	case -127: // Active >| toggled
		return KeyState(3)
	default: // idle
		return KeyState(0)
	}
}
