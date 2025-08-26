# GetKeyState

<a href="https://pkg.go.dev/github.com/CTNOriginals/getkeystate#section-sourcefiles">
    <img src="https://pkg.go.dev/badge/github.com/CTNOriginals/getkeystate#section-sourcefiles.svg" alt="Go Reference">
</a>

Get the state of any keyboard key.

## Requirements

- To be able to compile C code to go: [MinGW-w64](https://sourceforge.net/projects/mingw-w64/files/mingw-w64/mingw-w64-release/)
- For go to allow C code compilation, you may need to include the following settings in your project:
```json
"go.testEnvVars": { "CGO_ENABLED": 1 },
"go.toolsEnvVars": { "CGO_ENABLED": 1 },
```
> These json settings are for the vscode IDE go extension, they may be set differently for other IDE's or may not have to be set at all.

## Install package

```
go get github.com/CTNOriginals/getkeystate
```

## Usage

Import this package and call the GetKeyState function.

You may pass in any key as it's respective [Vertual Key Code](https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes).<br>
For example, the CapsLock key (VK_CAPITAL) would be passed in as `0x14` or `20`.

The return value of this function is a 2 bit int that represents the state of the key:<br>
The __right most__ bit represents wether it is currently held down (active).<br>
The __left most__ bit represents wether the key is currently toggled (works for any key, but mostly useful for keys like CapsLock).<br>

## Example

CapsLock state check
```go
import gks "github.com/CTNOriginals/getkeystate"

func main() {
	var capslock = gks.GetKeyState(0x14)

	println("Value: ", capslock)
	println("Active: ", capslock.IsActive())
	println("Toggled: ", capslock.IsToggled())
}
```

### Here are the different state conditions that this bit of code may produce:

> The following test case assumes that capslock starts as __not__ toggled

|Condition|Active|Toggled|Value|
|---------|:----:|:-----:|:---:|
|CapsLock is __not__ held down|`false`|`false`|`0`|
|CapsLock __is__ held down|`true`|`true`|`3`|
|CapsLock __is__ toggled but is __not__ held down|`false`|`true`|`1`|
|CapsLock __is__ toggled and __is__ held down|`true`|`false`|`2`|
