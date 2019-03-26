package appcontext

import (
	"fmt"

	"github.com/mahakamcloud/netd/config"
	newrelic "github.com/newrelic/go-agent"
)

type appContext struct {
	newrelicApp newrelic.Application
}

var context *appContext

type appContextError struct {
	Error error
}

func panicIfError(err error, werr error) {
	if err != nil {
		panic(appContextError{werr})
	}
}

func Init() {
	newrelicApp, err := newrelic.NewApplication(config.NewRelic())
	panicIfError(err, fmt.Errorf("Unable initiate NewRelic: %v", err))
	context = &appContext{
		newrelicApp: newrelicApp,
	}
}

func NewrelicApp() newrelic.Application {
	return context.newrelicApp
}
