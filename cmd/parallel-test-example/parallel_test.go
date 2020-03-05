package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"lab.weave.nl/weave/presentation/packages/testhelpers"
)

type ParallelSuiteOne struct {
	testhelpers.BaseSuite
}

func (s *ParallelSuiteOne) TestExample() {
	delayedFunc("Parallel suite one")
}

func TestParallelSuiteOne(t *testing.T) {
	t.Parallel()
	suite.Run(t, &ParallelSuiteOne{testhelpers.BaseSuite{ShowLogs: true}})
}

type ParallelSuiteTwo struct {
	testhelpers.BaseSuite
}

func (s *ParallelSuiteTwo) TestExample() {
	delayedFunc("Parallel suite two")

}

func TestParallelSuiteTwo(t *testing.T) {
	t.Parallel()
	suite.Run(t, &ParallelSuiteTwo{testhelpers.BaseSuite{ShowLogs: true}})
}

type NotParallelSuite struct {
	testhelpers.BaseSuite
}

func (s *NotParallelSuite) TestExample() {
	delayedFunc("Suite not running parallel")

}

func TestNotParallelSuite(t *testing.T) {
	suite.Run(t, &NotParallelSuite{testhelpers.BaseSuite{ShowLogs: true}})
}
