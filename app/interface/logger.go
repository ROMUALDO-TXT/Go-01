package interfaces

type Logger interface {
	Log(string, ...interface{})
	LogError(string, ...interface{})
	LogAccess(string, ...interface{})
}
