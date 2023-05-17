package trigger

import (
	"os"
	"os/signal"
)

func Signal(triggerCh chan struct{}, sigs ...os.Signal) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, sigs...)
	for range sigCh {
		triggerCh <- struct{}{}
	}
}
