package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FibTestSuite struct {
	suite.Suite

	cases []struct {
		Name   string
		Input  int64
		Output int64
	}
}

func (s *FibTestSuite) SetupTest() {
	s.T().Log("Setting up test suite")
	// Create test cases
	s.cases = []struct {
		Name   string
		Input  int64
		Output int64
	}{
		{
			Name:   "Test Fib in: 0, out: 0",
			Input:  0,
			Output: 0,
		},
		{
			Name:   "Test Fib in: 1, out: 1",
			Input:  1,
			Output: 1,
		},
		{
			Name:   "Test Fib in: 10, out: 55",
			Input:  10,
			Output: 55,
		},
		{
			Name:   "Test Fib in: 20, out: 6765",
			Input:  20,
			Output: 6765,
		},
		{
			Name:   "Test Fib in: 30, out: 832040",
			Input:  30,
			Output: 832040,
		},
	}
}

func (s *FibTestSuite) TearDownTest() {
	s.T().Log("Tearing down test suite")
}

func (s *FibTestSuite) TestFibRecBlockTests() {
	s.runFibCases(fibRec)
}

func (s *FibTestSuite) TestFibIterationBlockTests() {
	s.runFibCases(fibIteration)
}

func (s *FibTestSuite) TestFibDynamicBlockTests() {
	s.runFibCases(fibDynamicMap)
}

func (s *FibTestSuite) TestFibDynamicArrBlockTests() {
	s.runFibCases(fibDynamicArr)
}

func (s *FibTestSuite) runFibCases(f fibFunc) {
	for _, test := range s.cases {
		s.Run(test.Name, func() {
			output := f(test.Input)
			s.Equal(test.Output, output)
		})
	}
}

func TestFibTestSuite(t *testing.T) {
	suite.Run(t, new(FibTestSuite))
}
