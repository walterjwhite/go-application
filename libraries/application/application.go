package application

import (
	"context"
	"flag"

	"github.com/walterjwhite/go-application/libraries/identifier"
	"github.com/walterjwhite/go-application/libraries/logging"
	"github.com/walterjwhite/go-application/libraries/notification"
	"github.com/walterjwhite/go-application/libraries/shutdown"
)

var logFile = flag.String("Log", "", "The log file to write to")

func Configure() context.Context {
	identifier.Log()
	ctx := shutdown.Default()

	flag.Parse()

	logging.Set(*logFile)

	return ctx
}

func OnCompletion() {
	notification.OnCompletion()
}

func Wait(ctx context.Context) {
	// wait for CTRL+C (or context to expire)
	<-ctx.Done()
}
