package mux

import (
	"bytes"
)

var (
	slash            = []byte{'/'}
	aliasHolder      = []byte("_:_")
	aliasPrefix byte = ':'
)

type muxEntry struct {
	// the front part of path
	part []byte
	// :alias
	alias []byte
	// methods and assoicated handlers
	entries []*entry
	// trie
	nodes []*muxEntry
}

type entry struct {
	method  []byte
	handler Handler
}

func NewMuxEntry() *muxEntry {
	return &muxEntry{
		entries: make([]*entry, 0),
		nodes:   make([]*muxEntry, 0),
	}
}

func (e *muxEntry) setAlias(alias []byte) {
	if len(e.alias) == 0 {
		e.alias = alias
	}
	if len(e.alias) > 0 && !bytes.Equal(e.alias, alias) {
		panic("the muxEntry part alias set")
	}
}

func (e *muxEntry) trimSlash(path []byte) []byte {
	path = bytes.TrimPrefix(path, slash)
	path = bytes.TrimSuffix(path, slash)
	return path
}

func (e *muxEntry) Lookup(method, path []byte, rs *RequestScope) Handler {
	path = e.trimSlash(path)
	h := e.lookup(method, path, rs)
	if h == nil {
		h = NotFoundHandler
	}
	return h
}

func (e *muxEntry) lookup(method, path []byte, rs *RequestScope) Handler {
	me := e.findPath(path, rs)
	if me == nil {
		return nil
	}
	for _, entry := range me.entries {
		if bytes.Equal(entry.method, method) {
			return entry.handler
		}
	}
	return nil
}

func (e *muxEntry) findPath(path []byte, rs *RequestScope) *muxEntry {
	me := e

	var idx int
	for idx > -1 {
		if me == nil {
			return nil
		}
		idx = bytes.IndexByte(path, '/')
		if idx > 0 {
			me = me.find(path[:idx], rs)
			path = path[idx+1:]
		} else {
			me = me.find(path, rs)
		}
	}
	if me == e {
		me = nil
	}
	return me
}

func (e *muxEntry) find(path []byte, rs *RequestScope) *muxEntry {
	for _, node := range e.nodes {
		if bytes.Equal(node.part, path) {
			return node
		}
	}
	if !bytes.Equal(path, aliasHolder) {
		for _, node := range e.nodes {
			if bytes.Equal(node.part, aliasHolder) {
				if rs != nil {
					rs.SetPathParam(node.alias, bytes.TrimSpace(path))
				}
				return node
			}
		}
	}
	return nil
}

func (e *muxEntry) Add(method, path []byte, handler Handler) {
	path = e.trimSlash(path)
	me := e.add(path)
	for _, entry := range me.entries {
		if bytes.Equal(entry.method, method) {
			panic("muxEntry: add duplicate entry")
		}
	}
	me.entries = append(me.entries, &entry{method, handler})
}

func (e *muxEntry) add(path []byte) *muxEntry {
	var (
		me     = e
		idx    int
		field  []byte
		fields = bytes.Split(path, slash)
	)
	for idx, field = range fields {
		if len(field) > 1 && field[0] == aliasPrefix {
			field = aliasHolder
		}
		m := me.find(field, nil)
		if m == nil {
			idx--
			break
		}
		if bytes.Equal(field, aliasHolder) {
			m.setAlias(fields[idx][1:])
		}
		me = m
	}
	if idx < len(fields)-1 {
		for _, field := range fields[idx+1:] {
			nm := &muxEntry{
				entries: make([]*entry, 0),
				nodes:   make([]*muxEntry, 0),
			}
			if len(field) > 1 && field[0] == aliasPrefix {
				nm.part = aliasHolder
				nm.setAlias(field[1:])
			} else {
				nm.part = field
			}
			me.nodes = append(me.nodes, nm)
			me = nm
		}
	}
	return me
}

func lookUp(part string, e *muxEntry) {
	if part != "" {
		part += "/" + string(e.part)
	} else {
		part = string(e.part)
	}
	if len(e.nodes) == 0 {
		desc := ""
		for _, en := range e.entries {
			desc += string(en.method) + " "
		}
		desc += part
		println(desc)
	}
	for _, node := range e.nodes {
		lookUp(part, node)
	}
}
