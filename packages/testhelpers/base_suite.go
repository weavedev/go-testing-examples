package testhelpers

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

// BaseSuite base suite defines the default suite witih helper functions
// for testing
type BaseSuite struct {
	suite.Suite
	start    time.Time
	ShowLogs bool
}

// TearDownTest ran before executing a test case
func (s *BaseSuite) SetupTest() {
	if !s.ShowLogs || !testing.Verbose() {
		logrus.SetOutput(ioutil.Discard)
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	}

	s.start = time.Now()
}

// TearDownTest ran after the suite is done
func (s *BaseSuite) TearDownTest() {
	s.Logf("Test ran for: %s", time.Since(s.start))
}

// Log test logging
func (s *BaseSuite) Log(args ...interface{}) {
	s.T().Log(args...)
}

// Logf test logging using formatting
func (s *BaseSuite) Logf(format string, args ...interface{}) {
	s.T().Logf(format, args...)
}

// NoErrorWithFail checks if the error is nil and fails the test otherwise
func (s *BaseSuite) NoErrorWithFail(err error, args ...interface{}) {
	if err != nil {
		s.FailNow("Unexpected error:"+err.Error(), args...)
	}
}
