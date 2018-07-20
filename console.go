package ramble

import "fmt"

// Colorize the specified text in 256 terminal colors
func Colorize(color int, text string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m ", color, text)
}
