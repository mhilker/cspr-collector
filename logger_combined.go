package main

type CombinedLogger struct {
	Loggers []Logger
}

func (l *CombinedLogger) Log(data []CSPRequest) {
	for _, logger := range l.Loggers {
		logger.Log(data)
	}
}
