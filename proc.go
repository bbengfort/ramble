package ramble

import (
	"os"
	"os/signal"
	"syscall"
)

//===========================================================================
// OS Signal Handlers
//===========================================================================

// ShutdownSignal allows the registration of a a function to be called when
// SIGINT (CTRL+C) or SIGTERM is sent to the process, after which os.Exit
// is called with 0 if the function does not return an error or 1 if it does.
func ShutdownSignal(shutdown func() error) {
	// Make signal channel and register notifiers for Interupt and Terminate
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, syscall.SIGTERM)

	// Block until we receive a signal on the channel
	<-sigchan

	defer os.Exit(0)

	// Shutdown now that we've received the signal
	err := shutdown()
	if err != nil {
		os.Exit(1)
	}
}
