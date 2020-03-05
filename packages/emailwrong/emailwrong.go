package emailwrong

func NewClient() *SomeClient {
	return &SomeClient{}
}

type SomeClient struct {
}

func (c *SomeClient) SendMail(receiver, message string) error {
	// Calls some external API
	return nil
}
