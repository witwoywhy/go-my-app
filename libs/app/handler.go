package app

import "myapp/libs/log"

type HandleFunc func(ctx WebFrameworkContext, l log.Logger) error
