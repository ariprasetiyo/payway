package util

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func IsEmptyString(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

func IsErrorDoPanic(e error) {
	if e != nil {
		logrus.Panicln(e)
	}
}

func IsErrorDoPanicWithMessage(customMessage string, e error) {
	if e != nil {
		logrus.Panicln(customMessage, e)
	}
}

func IsErrorDoPrint(e error) {
	if e != nil {
		logrus.Error(e)
	}
}

func IsErrorDoPrintWithMessage(customMessage *string, e *error) {
	if e != nil {
		logrus.Error(customMessage, e)
	}
}
