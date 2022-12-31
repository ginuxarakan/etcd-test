package logger

import "github.com/sirupsen/logrus"

var Logrus *logrus.Logger

func init() {
	Logrus = logrus.New()
}
