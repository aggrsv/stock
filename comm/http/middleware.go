package http

import (
	"log"
	"net/http/httputil"
	"time"

	"github.com/pborman/uuid"
	"golang.org/x/net/context"

	"stock/comm/http/mux"

	"shendu.com/etc"
)

var (
	logLevel = etc.Int("log", "level")
)

var (
	PanicRecover = func(next mux.Handler) mux.Handler {
		return func(ctx context.Context) {
			if !etc.Dev() {
				defer func() {
					if err := recover(); err != nil {
						log.Println("============recover============")
						log.Println(err)
						log.Println("============recover============")
						mux.NotAvailableHandler(ctx)
						return
					}
				}()
			}
			next(ctx)
		}
	}

	Log = func(next mux.Handler) mux.Handler {
		return func(ctx context.Context) {
			var (
				c = ctx.(*Context)
				r = c.rs.Request()
				b = time.Now()
			)
			requestId := uuid.New()
			r.Header.Set("X-Request-Id", requestId)
			switch logLevel {
			case 1:
				log.Printf("%s REQ - %s %s", requestId, r.Method, r.URL.Path)
				next(ctx)
				log.Printf("%s RES - %d - %v", requestId, c.reply.Code, time.Since(b))
			case 2:
				d, _ := httputil.DumpRequest(r, true)
				log.Printf("%s REQ\n%s", requestId, string(d))
				next(ctx)
				log.Printf("%s RES\n%s - %v", requestId, string(c.reply.Json()), time.Since(b))
			default:
				next(ctx)
			}
		}
	}
)
