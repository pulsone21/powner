package middleware

import (
	"net/http"
)

// Middleware represents the type signature of a middleware
// function.
type Chain struct {
	middleware []Middleware
}

type Middleware interface {
	apply(http.Handler) http.Handler
}

func New(mws ...Middleware) Chain {
	return Chain{middleware: mws}
}

func (c Chain) Apply(final http.Handler) http.Handler {
	if len(c.middleware) < 1 {
		return final
	}

	mwApplied := final

	for i := len(c.middleware) - 1; i > -1; i-- {
		mwApplied = c.middleware[i].apply(mwApplied)
	}

	return mwApplied
}
