package umami

// Configuration holds Umami options,
// like instance URL and website ID.
// It is used to initialize handler, middleware, etc.
type Configuration struct {
	Href    string // Umami URL (f.e. https://umami.foo.bar)
	Website string // Umami website ID

	LogInf bool // Log information (like tracking events)
	LogErr bool // Log errors
}
