package http

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"stock/comm/http/proto"

	"stock/comm/http/mux"
)

type Context struct {
	context.Context

	rs *mux.RequestScope

	body []byte

	err   error
	reply *proto.Reply
}

func NewContext(ctx context.Context) *Context {
	rs := mux.Scope(ctx)
	r, w := rs.Request(), rs.ResponseWriter()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	return &Context{
		Context: ctx,
		rs:      rs,
	}
}

func Background() *Context {
	emptyCtx := context.Background()
	return &Context{
		Context: emptyCtx,
	}
}

func (c *Context) Error(err error) {
	c.err = err
}

func (c *Context) Scope() *mux.RequestScope {
	return c.rs
}

func (c *Context) Reply(reply interface{}) {
	c.reply = proto.NewReply(reply)
}

func (c *Context) ReplyCode(p proto.Proto) {
	c.reply = proto.ReplyCode(p)
}

func (c *Context) FillReply() {
	if c.err != nil {
		log.Println(c.err)
		if c.reply == nil {
			c.reply = proto.ReplyCode(proto.ServiceNotAvailable)
		}
	} else {
		if c.reply == nil {
			c.reply = proto.ReplyCode(proto.Ok)
		}
	}
	c.rs.SetReply(c.reply.Json())
}

func (c *Context) parseBody() error {
	if c.body != nil {
		return nil
	}
	r := c.rs.Request()
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	c.body = body
	return err
}

func (c *Context) Json(i interface{}) error {
	if err := c.parseBody(); err != nil {
		return err
	}
	if err := json.Unmarshal(c.body, i); err != nil {
		c.ReplyCode(proto.InvalidDataFormat)
		return err
	}
	return nil
}

func (c *Context) FormInt(key string) int64 {
	i, _ := strconv.ParseInt(c.formValue(key), 10, 64)
	return i
}

func (c *Context) FormIntWithSafe(key string) (int64, error) {
	i, err := strconv.ParseInt(c.formValue(key), 10, 64)
	if err != nil {
		c.ReplyCode(proto.InvalidParameter)
		return 0, err
	}
	return i, nil
}

func (c *Context) FormString(key string) string {
	return c.formValue(key)
}

func (c *Context) FormMultiIntWithComma(key string) ([]int64, error) {
	ints := make([]int64, 0)
	for _, field := range strings.FieldsFunc(c.FormString(key), func(r rune) bool { return r == ',' }) {
		i, err := strconv.ParseInt(field, 10, 64)
		if err != nil {
			c.ReplyCode(proto.InvalidParameter)
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func (c *Context) FormMultiStringWithComma(key string) ([]string, error) {
	fields := strings.FieldsFunc(c.FormString(key), func(r rune) bool { return r == ',' })
	for _, field := range fields {
		if field == "" {
			c.ReplyCode(proto.InvalidParameter)
			return nil, errors.New("empty string exists")
		}
	}
	return fields, nil
}

func (c *Context) formValue(key string) string {
	r := c.rs.Request()
	return strings.TrimSpace(r.FormValue(key))
}

func (c *Context) VarsInt(key string) int64 {
	i, _ := strconv.ParseInt(c.vars(key), 10, 64)
	return i
}

func (c *Context) VarsIntWithSafe(key string) (int64, error) {
	i, err := strconv.ParseInt(c.vars(key), 10, 64)
	if err != nil {
		c.ReplyCode(proto.InvalidParameter)
		return 0, err
	}
	return i, nil
}

func (c *Context) VarsString(key string) string {
	return c.vars(key)
}

func (c *Context) vars(key string) string {
	return string(c.rs.PathParam([]byte(key)))
}
