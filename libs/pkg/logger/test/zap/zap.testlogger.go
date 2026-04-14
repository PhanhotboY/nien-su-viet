package zaptest

import (
	"testing"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"
)

type zapTestLogger struct {
	t *testing.T
	logger.Logger
}

type ZapTestLogger interface {
	testlogger.TestLogger
}

// NewZapLogger create new zap logger
func NewZapTestLogger(
	t *testing.T,
	l logger.Logger,
) ZapTestLogger {
	zapTestLogger := &zapTestLogger{t: t, Logger: l}

	return zapTestLogger
}

// Error uses fmt.Sprint to construct and log a message.
func (z *zapTestLogger) TestError(args ...interface{}) {
	z.Logger.Error(args...)
	z.t.Fail()
}

// Errorw logs a message with some additional context.
func (z *zapTestLogger) TestErrorw(msg string, fields logger.Fields) {
	z.Logger.Errorw(msg, fields)
	z.t.Fail()
}

// Errorf uses fmt.Sprintf to log a templated message.
func (z *zapTestLogger) TestErrorf(template string, args ...interface{}) {
	z.Logger.Errorf(template, args...)
	z.t.Fail()
}

// Err uses error to log a message.
func (z *zapTestLogger) TestErr(msg string, err error) {
	z.Logger.Err(msg, err)
	z.t.Fail()
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (z *zapTestLogger) TestFatal(args ...interface{}) {
	z.Logger.Fatal(args...)
	z.t.FailNow()
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (z *zapTestLogger) TestFatalf(template string, args ...interface{}) {
	z.Logger.Fatalf(template, args...)
	z.t.FailNow()
}
