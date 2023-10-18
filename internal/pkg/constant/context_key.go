package constant

type ContextKey any

var (
	CONTEXT_KEY_USER         = ContextKey("user")
	CONTEXT_KEY_REQUEST_USER = ContextKey("request-user")
)
