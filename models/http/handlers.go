package proto

import (
	"log"

	"golang.org/x/net/context"

	"stock/comm/http/mux"

	"stock/comm/http"
)

var (
	NotFoundHandler = func(ctx context.Context) {
		r := mux.Scope(ctx).Request()
		log.Printf("%s %s NotFound", r.Method, r.URL.Path)
		c := http.NewContext(ctx)
		c.ReplyCode(ActionNotFound)
		c.FillReply()
	}
	ServiceNotAvailableHandler = handler(ServiceNotAvailable)
	UnauthorizedHandler        = handler(TokenExpired)
)

func handler(code Code) mux.Handler {
	return func(ctx context.Context) {
		c := ctx.(*http.Context)
		c.ReplyCode(code)
		c.FillReply()
	}
}

func init() {
	mux.SetNotFoundHandler(NotFoundHandler)
	mux.SetNotAvailableHandler(ServiceNotAvailableHandler)
	mux.SetUnauthorizedHandler(UnauthorizedHandler)
}
