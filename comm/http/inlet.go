package http

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"stock/comm/http/mux"

	"stock/etc"
)

var DefaultMux = New()

type Mux struct {
	*mux.Mux
}

func New() *Mux {
	m := mux.New()
	Ctx := func(next mux.Handler) mux.Handler {
		return func(ctx context.Context) {
			c := NewContext(ctx)
			next(c)
		}
	}
	m.Use(Ctx, Log, PanicRecover)
	return &Mux{m}
}

func Route(addr string, middleware ...mux.Middleware) *Mux {
	m := &Mux{New().Mux.Mux(addr)}
	m.Use(middleware...)
	return m
}

func ListenAndServeEtc(key string, h http.Handler) {
	addr := fmt.Sprintf(":%d", etc.Int("applet", key))
	log.Printf("applet %s listening on %s", key, addr)
	ListenAndServe(addr, h)
}

func ListenAndServe(addr string, h http.Handler) {
	http.ListenAndServe(addr, h)
}

type Handler func(*Context) error

func (m *Mux) Group(path string, f func(m *Mux), middleware ...mux.Middleware) {
	nm := &Mux{m.Mux.Mux(path, middleware...)}
	f(nm)
}

func (m *Mux) Get(path string, handler Handler, middleware ...mux.Middleware) {
	m.handle("GET", path, handler, middleware...)
}

func (m *Mux) Post(path string, handler Handler, middleware ...mux.Middleware) {
	m.handle("POST", path, handler, middleware...)
}

func (m *Mux) Put(path string, handler Handler, middleware ...mux.Middleware) {
	m.handle("PUT", path, handler, middleware...)
}

func (m *Mux) Delete(path string, handler Handler, middleware ...mux.Middleware) {
	m.handle("DELETE", path, handler, middleware...)
}

func (m *Mux) handle(method, path string, f func(*Context) error, middleware ...mux.Middleware) {
	mh := func(ctx context.Context) {
		c := ctx.(*Context)
		c.Error(f(c))
		c.FillReply()
	}
	m.Handle(method, path, mh, middleware...)
}
