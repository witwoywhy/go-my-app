package app

import "myapp/libs/log"

type HandlerFunc func(ctx WebFrameworkContext) error

type HandleFunc func(ctx WebFrameworkContext, l log.Logger) error
