package interfaces

type Logger interface {
	Info(...any)
	Infof(string, ...any)
}
