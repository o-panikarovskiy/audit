package middlewares

import (
	"audit/src/config"
	"audit/src/utils"
	"context"
	"fmt"
	"log"
	"net/http"
)

type ctxKeyType int

const (
	jsonKey ctxKeyType = iota
	configKey
	sessionKey
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

// GetContext returns typed context from request
func GetContext(r *http.Request) Context {
	ctx, ok := r.Context().(Context)
	if !ok {
		log.Panic("Failed to get custom Context from request")
		return ctx
	}
	return ctx
}

// WithJSON put json data to contex
func (ctx Context) WithJSON(data *utils.StringMap) Context {
	return Context{context.WithValue(ctx, jsonKey, data)}
}

// JSON get json data from context
func (ctx Context) JSON() *utils.StringMap {
	val, ok := ctx.Value(jsonKey).(*utils.StringMap)

	if !ok {
		panic(fmt.Errorf("Failed to get value from context %v by key %v", val, jsonKey))
	}

	return val
}

// WithAppConfig put data to context
func (ctx Context) WithAppConfig(cfg *config.AppConfig) Context {
	return Context{context.WithValue(ctx, configKey, cfg)}
}

// Config get data from context
func (ctx Context) Config() *config.AppConfig {
	val, ok := ctx.Value(configKey).(*config.AppConfig)

	if !ok {
		panic(fmt.Errorf("Failed to get value from context %v by key %v", val, jsonKey))
	}

	return val
}
