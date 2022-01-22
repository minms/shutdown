package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitTerminationSignal(stopFunc func())  {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer func() {
		signal.Stop(ch)
		close(ch)
	}()
	<-ch
	stopFunc()
}