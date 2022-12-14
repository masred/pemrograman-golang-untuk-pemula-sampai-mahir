package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is Trace")
	logger.Debug("This is Debuge")
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
}
