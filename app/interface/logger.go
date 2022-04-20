package interfaces

type Logger interface {
	Log(...interface{})
	LogError(string, ...interface{})
	LogAccess(string, ...interface{})
}
