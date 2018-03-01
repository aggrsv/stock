package context

import (
	"log"
	"strconv"

	"golang.org/x/net/context"

	"stock/comm/http/mux"

	"stock/comm/http"
)

var (
	AppIdKey = "X-Sd-App-Id"
)

var (
	Header = func(next mux.Handler) mux.Handler {
		return func(ctx context.Context) {
			c := ctx.(*http.Context)
			appId, err := strconv.ParseInt(c.Scope().Request().Header.Get(AppIdKey), 10, 64)
			if err != nil {
				log.Println(err)
				mux.NotAvailableHandler(ctx)
				return
			}
			c.Context = WithAppId(c.Context, appId)
			next(c)
		}
	}
)
