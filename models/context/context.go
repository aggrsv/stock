package context

import (
	"golang.org/x/net/context"
)

const (
	userKey  = "User"
	tokenKey = "Token"
	appIdKey = "AppId"
)

type (
	User struct {
		Id   int64
		Name string
	}
)

func WithUser(ctx context.Context, uid int64, username string) context.Context {
	return context.WithValue(ctx, userKey, &User{uid, username})
}

func Uid(ctx context.Context) int64 {
	return ctx.Value(userKey).(*User).Id
}

func Username(ctx context.Context) string {
	return ctx.Value(userKey).(*User).Name
}

func WithToken(ctx context.Context, tokenId string) context.Context {
	return context.WithValue(ctx, tokenKey, tokenId)
}

func Token(ctx context.Context) string {
	return ctx.Value(tokenKey).(string)
}

func WithAppId(ctx context.Context, appId int64) context.Context {
	return context.WithValue(ctx, appIdKey, appId)
}

func AppId(ctx context.Context) int64 {
	return ctx.Value(appIdKey).(int64)
}
