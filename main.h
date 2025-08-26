#include <stdio.h>
#include <Windows.h>

// @returns 
// An integer representing the state of the key.
// @note
// Return: 0 and 1 indicate the toggled state, if the key pressed before, 
// it now returns a 1, if its pressed again it will return a 0.
// @note
// Return: -127 indicates that the key is currently held down and after releasing the state will change to a 1.
// -128 also indicates that the key is held down, but instead it will change to a 0 on release
int C_GetKeyState(int key) {
    int state = GetKeyState(key);
    return state;
}