package domain

// Handler defines the interface for HTTP handlers
type Handler interface {
	// RegisterRoutes registers all routes for the handler
	RegisterRoutes(router Router)
}

// Router defines the interface for HTTP routing
type Router interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	PUT(path string, handler HandlerFunc)
	DELETE(path string, handler HandlerFunc)
	Use(middleware ...Middleware)
}

// HandlerFunc represents a function that handles HTTP requests
type HandlerFunc func(ctx Context) error

// Context represents the HTTP request context
type Context interface {
	JSON(code int, obj interface{}) error
	Bind(obj interface{}) error
	Param(key string) string
	Query(key string) string
	GetHeader(key string) string
	SetHeader(key, value string)
}

// Middleware represents HTTP middleware
type Middleware func(HandlerFunc) HandlerFunc

// Repository defines the interface for data access
type Repository interface {
	// Add your repository methods here
}

// Service defines the interface for business logic
type Service interface {
	// Add your service methods here
}

// UseCase defines the interface for application use cases
type UseCase interface {
	// Add your use case methods here
}
