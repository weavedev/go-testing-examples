package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"lab.weave.nl/weave/presentation/packages/email"
)

func setFormatter() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}

func main() {
	setFormatter()
	api := CreateNewAPI(email.NewClient())
	req, err := http.NewRequest("GET", "mywonderfulltest.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Basic d2ltQHdlYXZlLm5sOk9wZW5TZXNhbWU=")
	req.Header.Add("x-forwarded-for", "127.0.0.1")
	err = api.SignInAPICall(req)
	if err != nil {
		panic(err)
	}
	log.Info("User signed in")
}

type MyWonderFullAPI struct {
	EmailClient email.Client
}

func CreateNewAPI(emailClient email.Client) MyWonderFullAPI {
	return MyWonderFullAPI{
		EmailClient: emailClient,
	}
}

func (a *MyWonderFullAPI) SignInAPICall(req *http.Request) error {
	user, _, ok := req.BasicAuth()
	if !ok {
		return fmt.Errorf("incorrect basic auth provided")
	}
	ip := GetIP(req)
	lastIPUsed := "0.0.0.0"
	// Check if the signin request is send from a new IP
	if ip != lastIPUsed {
		err := a.EmailClient.SendMail(user, fmt.Sprintf("Signin from a new IP: %s", ip))
		if err != nil {
			log.WithError(err).Errorf("couldnt send email")
			return err
		}
	}
	return nil
}

func (a *MyWonderFullAPI) SignInAPICallAsyncEmail(req *http.Request) error {
	user, _, ok := req.BasicAuth()
	if !ok {
		return fmt.Errorf("incorrect basic auth provided")
	}
	ip := GetIP(req)
	lastIPUsed := "0.0.0.0"
	// Check if the signin request is send from a new IP
	if ip != lastIPUsed {
		go func() {
			time.Sleep(time.Second * 1)
			err := a.EmailClient.SendMail(user, fmt.Sprintf("Signin from a new IP: %s", ip))
			if err != nil {
				log.WithError(err).Errorf("couldnt send email")
			}
		}()
	}
	return nil
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
