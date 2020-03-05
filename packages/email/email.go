package email

import "github.com/stretchr/testify/mock"

// Client interface represent the email client
type Client interface {
	SendMail(string, string) error
}

// NewClient creates a new email client
func NewClient() Client {
	return &SomeClient{}
}

// SomeClient implementation of some email client
type SomeClient struct {
}

// SendMail implementation of the sendmail function of the SomeClient
func (c *SomeClient) SendMail(receiver, message string) error {
	// Calls some external API
	return nil
}

// ClientMock mock implementation of the email client using the testify mock struct
type ClientMock struct {
	mock.Mock
}

// SendMail mocked send mail implementation
func (c *ClientMock) SendMail(receiver, message string) error {
	args := c.MethodCalled("SendMail", receiver, message)
	return args.Error(0)
}
