package kitchen

// Logger is the logging surface made available to recipes.
type Logger interface {
	Info(message string)
	Success(message string)
	Warning(message string)
	Error(message string)
	Debug(message string)
}
