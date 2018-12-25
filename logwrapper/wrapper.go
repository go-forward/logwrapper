package logwrapper

import "log"

type Logger struct {
	rawLogger *log.Logger
}

func (l *Logger) SetRawLogger(rl *log.Logger) {
	l.rawLogger = rl
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.rawLogger.Printf("[Debug] " + format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.rawLogger.Printf("[Info] " + format, v...)

}

func (l *Logger) Noticef(format string, v ...interface{}) {
	l.rawLogger.Printf("[Notice] " + format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.rawLogger.Printf("[Warning] " + format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.rawLogger.Printf("[Error] " + format, v...)
}

func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.rawLogger.Printf("[Critical] " + format, v...)
}

func (l *Logger) Alertf(format string, v ...interface{}) {
	l.rawLogger.Printf("[Alert] " + format, v...)
}

func (l *Logger) Emergencyf(format string, v ...interface{}) {
	l.rawLogger.Printf("[Emergency] " + format, v...)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.rawLogger.Printf("[Panic] " + format, v...)
}