package envfs

import "context"

type customEnvironmentContextKey struct{}

// ContextWithEnvironmentGetter returns a new context with the given EnvironmentGetter
func ContextWithEnvironmentGetter(ctx context.Context, getter EnvironmentGetter) context.Context {
	return context.WithValue(ctx, customEnvironmentContextKey{}, getter)
}

// EnvironmentGetterFromContext retrieves the EnvironmentGetter from the context,
// returning the default system implementation if none is found
func EnvironmentGetterFromContext(ctx context.Context) EnvironmentGetter {
	if getter, ok := ctx.Value(customEnvironmentContextKey{}).(EnvironmentGetter); ok {
		return getter
	}

	return defaultSystemEnvironmentGetter
}