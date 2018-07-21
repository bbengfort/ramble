package ramble

import "time"

// ChatTime returns the current timestamp formatted for the chat window.
func ChatTime() string {
	return time.Now().Format("15:04:05")
}
