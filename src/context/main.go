package context

import (
	"audit/src/utils"
	"context"
	"fmt"
)

type ctxKeyType string

const (
	jsonKey    ctxKeyType = "json"
	sessionKey            = "sid"
	// ...
)

// Context for middlewares
type Context struct {
	context.Context
}

// NewContext constructor
func NewContext(ctx context.Context) Context {
	return Context{ctx}
}

// WithJSON writes json data to contex
func (ctx Context) WithJSON(data *utils.StringMap) Context {
	return Context{context.WithValue(ctx, jsonKey, data)}
}

// JSON get json data from context
func (ctx Context) JSON() *utils.StringMap {
	val, ok := ctx.Value(jsonKey).(*utils.StringMap)

	if !ok {
		panic(fmt.Errorf("Failed to get *utils.StringMap from context %v", val))
	}

	return val
}
