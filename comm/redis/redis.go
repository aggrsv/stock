package redis

import (
	"log"
	"time"

	"stock/etc"

	"github.com/garyburd/redigo/redis"
)

type Client struct {
	*redis.Pool
}

func (c *Client) borrow() redis.Conn {
	if c.Pool == nil {
		log.Fatal("must invoke init() before borrow")
	}
	return c.Get()
}

func (c *Client) Do(f func(c redis.Conn) error) error {
	conn := c.borrow()
	defer conn.Close()
	return f(conn)
}

func (c *Client) Close() error {
	return c.Pool.Close()
}

func New(opt *InitOption, dailOpts ...redis.DialOption) *Client {
	if opt == nil {
		opt = defaultInitOption
	}
	p := &redis.Pool{
		MaxIdle:     opt.MaxIdle,
		IdleTimeout: time.Hour,
		Dial: func() (redis.Conn, error) {
			dailOpts = append(dailOpts, redis.DialPassword(opt.Password))
			c, err := redis.Dial("tcp", opt.Addr, dailOpts...)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxActive: 50,
		Wait:      true,
	}

	c := &Client{p}
	return c
}

type InitOption struct {
	Addr     string
	Password string
	MaxIdle  int
}

var defaultInitOption = &InitOption{
	Addr:     etc.String("redis2", "addr"),
	Password: etc.String("redis2", "password"),
	MaxIdle:  10,
}
