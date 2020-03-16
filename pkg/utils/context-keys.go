// Package utils contains misc utilities
package utils

// ContextKey defines a type for context keys shared in the app
type ContextKey string

// ContextKeys holds the context keys throughout the project
type ContextKeys struct {
	ProviderCtxKey ContextKey // Provider in Auth
	UserCtxKey     ContextKey // User db object in Auth
}

// ProjectContextKeys the project's context keys
func ProjectContextKeys() ContextKeys {
	return ContextKeys{
		ProviderCtxKey: "provider",
		UserCtxKey:     "auth-user",
	}
}
