package testlogger

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type TestLogger interface {
	logger.Logger
	TestError(args ...interface{})                   // equal to logger.Error() + test.Fail()
	TestErrorw(msg string, fields logger.Fields)     // equal to logger.Errorw() + test.Fail()
	TestErrorf(template string, args ...interface{}) // equal to logger.Errorf() + test.Fail()
	TestErr(msg string, err error)                   // equal to logger.Err() + test.Fail()
	TestFatal(args ...interface{})                   // equal to logger.Fatal() + test.FailNow()
	TestFatalf(template string, args ...interface{}) //	equal to logger.Fatalf() + test.FailNow()
}
