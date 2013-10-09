// Package cursor provides ANSI Escape sequences
// introduced in VT-100 terminals for cursor and screen
// manipulation operations.
package cursor

import "fmt"

var Esc = "\x1b"

// MoveTo returns ANSI escape sequence to move cursor
// to specified position on screen.
func MoveTo(line, col int) string {
	return fmt.Sprintf("%s[%d;%dH", Esc, line, col)
}

// MoveUp returns ANSI escape sequence to move cursor
// up n lines.
func MoveUp(n int) string {
	return fmt.Sprintf("%s[%dA", Esc, n)
}

// MoveDown returns ANSI escape sequence to move cursor
// down n lines.
func MoveDown(n int) string {
	return fmt.Sprintf("%s[%dB", Esc, n)
}

// MoveRight returns ANSI escape sequence to move cursor
// right n columns.
func MoveRight(n int) string {
	return fmt.Sprintf("%s[%dC", Esc, n)
}

// MoveLeft returns ANSI escape sequence to move cursor
// left n columns.
func MoveLeft(n int) string {
	return fmt.Sprintf("%s[%dD", Esc, n)
}

// MoveUpperLeft returns ANSI escape sequence to move cursor
// to upper left corner of screen.
func MoveUpperLeft(n int) string {
	return fmt.Sprintf("%s[%dH", Esc, n)
}

// MoveNextLine returns ANSI escape sequence to move cursor
// to next line.
func MoveNextLine() string {
	return fmt.Sprintf("%sE", Esc)
}

// ClearLineRight returns ANSI escape sequence to clear line
// from right of the cursor.
func ClearLineRight() string {
	return fmt.Sprintf("%s[0K", Esc)
}

// ClearLineLeft returns ANSI escape sequence to clear line
// from left of the cursor.
func ClearLineLeft() string {
	return fmt.Sprintf("%s[1K", Esc)
}

// ClearEntireLine returns ANSI escape sequence to clear current line.
func ClearEntireLine() string {
	return fmt.Sprintf("%s[2K", Esc)
}

// ClearScreenDown returns ANSI escape sequence to clear screen below.
// the cursor.
func ClearScreenDown() string {
	return fmt.Sprintf("%s[0J", Esc)
}

// ClearScreenUp returns ANSI escape sequence to clear screen above.
// the cursor.
func ClearScreenUp() string {
	return fmt.Sprintf("%s[1J", Esc)
}

// ClearEntireScreen returns ANSI escape sequence to clear the screen.
func ClearEntireScreen() string {
	return fmt.Sprintf("%s[2J", Esc)
}

// SaveAttributes returns ANSI escape sequence to save current position
// and attributes of the cursor.
func SaveAttributes() string {
	return fmt.Sprintf("%s7", Esc)
}

// RestoreAttributes returns ANSI escape sequence to restore saved position
// and attributes of the cursor.
func RestoreAttributes() string {
	return fmt.Sprintf("%s8", Esc)
}
