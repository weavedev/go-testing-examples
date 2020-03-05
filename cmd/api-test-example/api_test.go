package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"lab.weave.nl/weave/presentation/packages/email"
	"lab.weave.nl/weave/presentation/packages/testhelpers"
)

type APITestSuite struct {
	testhelpers.BaseSuite

	basicAuth      string
	defaultRequest *http.Request

	api MyWonderFullAPI
}

func (s *APITestSuite) SetupTest() {
	s.BaseSuite.SetupTest()

	// Create new request with headers
	req, err := http.NewRequest("GET", "mywonderfulltest.com", nil)
	s.NoErrorWithFail(err)
	req.Header.Add("Authorization", s.basicAuth)
	req.Header.Add("x-forwarded-for", "127.0.0.1")
	s.defaultRequest = req

	// Initialize the email mock
	emailClient := &email.ClientMock{}
	emailClient.On("SendMail", mock.Anything, mock.Anything).Return(nil)

	// Create a new API with the emailClient mock
	s.api = CreateNewAPI(emailClient)
}

func (s *APITestSuite) TearDownTest() {
	s.BaseSuite.TearDownTest()
}

func (s *APITestSuite) TestSignInAPICall() {
	s.NoErrorWithFail(s.api.SignInAPICall(s.defaultRequest))
}

func (s *APITestSuite) TestSignInAPICallEmailError() {
	// Overwrite the email mock behaviour such that it returns an error
	emailClient := &email.ClientMock{}
	emailClient.On("SendMail", "error@weave.nl", mock.Anything).Return(fmt.Errorf("unable to send email"))
	emailClient.On("SendMail", "wim@weave.nl", mock.Anything).Return(nil)
	s.api.EmailClient = emailClient
	// Since the default requests contains the email wim@weave.nl we don't expect an error
	s.NoErrorWithFail(s.api.SignInAPICall(s.defaultRequest))
	failingReq, err := http.NewRequest("GET", "mywonderfulltest.com", nil)
	s.NoErrorWithFail(err)
	// Basic auth representing: error@weave.nl:superSecret
	failingReq.Header.Add("Authorization", "Basic ZXJyb3JAd2VhdmUubmw6c3VwZXJTZWNyZXQ=")
	failingReq.Header.Add("x-forwarded-for", "127.0.0.1")
	// Since the email client returns an error we also expect the API route to return an error
	s.Error(s.api.SignInAPICall(failingReq))
}

func (s *APITestSuite) TestSignInAPICallVerifyEmailClientCall() {
	// Overwrite the mock behaviour
	emailClient := &email.ClientMock{}
	emailClient.On("SendMail", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		// Validate that the email client is called with the correct parameters
		email := args.Get(0).(string)
		s.Equal("wim@weave.nl", email)
		emailMessage := args.Get(1).(string)
		s.Equal(fmt.Sprintf("Signin from a new IP: %s", "127.0.0.1"), emailMessage)
	})
	s.api.EmailClient = emailClient
	// Call the API route
	s.NoErrorWithFail(s.api.SignInAPICall(s.defaultRequest))
	// Assert that the send email function is called
	emailClient.AssertNumberOfCalls(s.T(), "SendMail", 1)
}

func (s *APITestSuite) TestSignInAPICallVerifyEmailClientCallAsync() {
	// Overwrite the mock behaviour
	emailClient := &email.ClientMock{}
	var wg sync.WaitGroup
	emailClient.On("SendMail", "wim@weave.nl", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		// Validate that the email client is called with the correct parameters
		email := args.Get(0).(string)
		s.Equal("wim@weave.nl", email)
		emailMessage := args.Get(1).(string)
		s.Equal(fmt.Sprintf("Signin from a new IP: %s", "127.0.0.1"), emailMessage)
		wg.Done()
		s.Log("End of mock call")
	})
	s.api.EmailClient = emailClient
	wg.Add(1)
	s.NoErrorWithFail(s.api.SignInAPICallAsyncEmail(s.defaultRequest))
	s.Log("Waiting for async method call to finish")
	wg.Wait()
	s.Log("Async mock call finished")

	// Assert that the send email function is called
	emailClient.AssertNumberOfCalls(s.T(), "SendMail", 1)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, &APITestSuite{
		// Basic auth representing wim@weave.nl:OpenSesame
		basicAuth: "Basic d2ltQHdlYXZlLm5sOk9wZW5TZXNhbWU=",
	})
}
