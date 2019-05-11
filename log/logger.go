package log

import "github.com/sirupsen/logrus"

func Debug(v ...interface{}) {
	logrus.Debug(v)
}

func Debugf(format string, v ...interface{}) {
	logrus.Debugf(format, v)
}

func Info(v ...interface{}) {
	logrus.Info(v)
}

func Infof(format string, v ...interface{}) {
	logrus.Infof(format, v)
}

func Warn(v ...interface{}) {
	logrus.Warn(v)
}

func Warnf(format string, v ...interface{}) {
	logrus.Warnf(format, v)
}

func Error(v ...interface{}) {
	logrus.Error(v)
}

func Errorf(format string, v ...interface{}) {
	logrus.Errorf(format, v)
}

func Fatal(v ...interface{}) {
	logrus.Fatal(v)
}

func Fatalf(format string, v ...interface{}) {
	logrus.Fatalf(format, v)
}