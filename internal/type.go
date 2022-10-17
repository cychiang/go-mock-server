package internal

type Config struct {
	Name    string
	Routers []Router
}

type Router struct {
	Path        string
	Method      string
	ContentType string
	StatusCode  int
	Body        string
}
