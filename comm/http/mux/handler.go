package mux

import (
	"golang.org/x/net/context"
)

type Handler func(context.Context)

var (
	NotFoundHandler     = func(context.Context) {}
	NotAvailableHandler = func(context.Context) {}
	UnauthorizedHandler = func(context.Context) {}
)

func SetNotFoundHandler(f Handler) {
	NotFoundHandler = f
}

func SetNotAvailableHandler(f Handler) {
	NotAvailableHandler = f
}

func SetUnauthorizedHandler(f Handler) {
	UnauthorizedHandler = f
}
