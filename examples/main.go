package main

import gks "github.com/CTNOriginals/getkeystate"

func main() {
	var capslock = gks.GetKeyState(gks.VK_CAPITAL)

	println("Value: ", capslock)
	println("Active: ", gks.IsActive(capslock))
	println("Toggled: ", gks.IsToggled(capslock))
}
