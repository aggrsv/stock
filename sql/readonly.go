package sql

import (
	"strings"
	"sync"

	"shendu.com/etc"
	"shendu.com/etc/conf"
)

var (
	readOnlyConfPath = etc.String("sql", "readonly_conf")
)

type ReadOnly struct {
	Weight       int
	MaxOpenConns int
	MaxIdelConns int
	Domain       string
	replaced     bool
}

func (r *ReadOnly) Replace(dsn, key string) {
	if r.replaced {
		return
	}
	r.Domain = strings.Replace(dsn, key, r.Domain, 1)
	r.replaced = true
}

var (
	m sync.Mutex
	c map[string][]*ReadOnly
)

func wrapDsn(dsn, key string, l []*ReadOnly) int {
	weight := 0
	for _, item := range l {
		weight += item.Weight
		item.Replace(dsn, key)
	}
	return weight
}

func Slaves(dsn string) ([]*ReadOnly, int, error) {
	m.Lock()
	defer m.Unlock()
	if c == nil {
		c = make(map[string][]*ReadOnly)
		if err := conf.ReadJSON(readOnlyConfPath, &c); err != nil {
			return nil, 0, err
		}
	}
	for key, l := range c {
		if strings.Contains(dsn, key) {
			dup := make([]*ReadOnly, 0)
			for _, i := range l {
				d := *i
				dup = append(dup, &d)
			}
			weight := wrapDsn(dsn, key, dup)
			return dup, weight, nil
		}
	}
	return nil, 0, nil
}
