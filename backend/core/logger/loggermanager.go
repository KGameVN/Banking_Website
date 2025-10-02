package logger


type LoggerManager struct {
	dispatcher 	*Dispatcher
	Sink *sink
}

func GetLogger() Logger {  
	return NewLogger()
}

