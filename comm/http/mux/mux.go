package mux

import (
	"net/http"

	"golang.org/x/net/context"
)

type Middleware func(Handler) Handler

type Mux struct {
	entry *muxEntry

	prefix     string
	middleware []Middleware
}

func New() *Mux {
	return &Mux{
		entry:      NewMuxEntry(),
		middleware: make([]Middleware, 0),
	}
}

func (m *Mux) Use(middleware ...Middleware) {
	m.middleware = append(m.middleware, middleware...)
}

func (m *Mux) Group(path string, f func(m *Mux), middleware ...Middleware) {
	nm := m.Mux(path, middleware...)
	f(nm)
}

func (m *Mux) Mux(path string, middleware ...Middleware) *Mux {
	nm := m.dup()
	nm.prefix = m.prefix + path
	nm.middleware = append(nm.middleware, m.middleware...)
	nm.middleware = append(nm.middleware, middleware...)
	return nm
}

func (m *Mux) dup() *Mux {
	nm := new(Mux)
	nm.entry = m.entry
	return nm
}

func (m *Mux) Get(path string, handler Handler, middleware ...Middleware) {
	m.Handle("GET", path, handler, middleware...)
}

func (m *Mux) Post(path string, handler Handler, middleware ...Middleware) {
	m.Handle("POST", path, handler, middleware...)
}

func (m *Mux) Put(path string, handler Handler, middleware ...Middleware) {
	m.Handle("PUT", path, handler, middleware...)
}

func (m *Mux) Delete(path string, handler Handler, middleware ...Middleware) {
	m.Handle("DELETE", path, handler, middleware...)
}

func (m *Mux) Handle(method, path string, handler Handler, middleware ...Middleware) {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	for i := len(m.middleware) - 1; i >= 0; i-- {
		handler = m.middleware[i](handler)
	}
	path = m.prefix + path
	m.entry.Add([]byte(method), []byte(path), handler)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rs := &RequestScope{r: r, w: w}
	ctx := context.WithValue(context.Background(), RequestScopeKey, rs)
	m.entry.Lookup([]byte(r.Method), []byte(r.URL.Path), rs)(ctx)
	rs.done()
}

func (m *Mux) All() {
	lookUp("", m.entry)
}
