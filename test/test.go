package main

import (
	gks "github.com/CTNOriginals/getkeystate"
)

func main() {
	var capslock = gks.GetKeyState(0x14)

	println("Value: ", capslock)
	println("Active: ", capslock.IsActive())
	println("Toggled: ", capslock.IsToggled())
}
