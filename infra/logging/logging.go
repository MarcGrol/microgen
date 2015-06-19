package logging

import (
	"log"
	"log/syslog"
)

type Logging struct {
	logger *log.Logger
}

func NewLogging(prio syslog.Priority, progName string) *Logging {
	logging := new(Logging)
	logwriter, err := syslog.New(prio, progName)
	if err == nil {
		log.SetOutput(logwriter)
	}
	return logging
}

func (l *Logging) Fatalf(format string, args ...interface{}) {
	l.logger.Panicf(format, args)
}

func (l *Logging) Errorf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args)
}

func (l *Logging) Warningf(format string, args ...interface{}) {
	l.logger.Printf(format, args)
}

func (l *Logging) Infof(format string, args ...interface{}) {
	l.logger.Printf(format, args)
}

func (l *Logging) Debugf(format string, args ...interface{}) {
	l.logger.Printf(format, args)
}
