package gowhisper_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/co0p/gowhisper"
)

type ClientMock struct {
	lastRequest   *http.Request
	checkFn       func(*http.Request, *testing.T)
	t             *testing.T
	StatusCode    int
	StatusMessage string
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	c.checkRequest(req, c.t)
	return &http.Response{
		StatusCode: c.StatusCode,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(c.StatusMessage))),
	}, nil
}

func (c *ClientMock) checkRequest(req *http.Request, t *testing.T) {
	if c.checkFn != nil {
		c.checkFn(req, t)
	}
}

func Test_SendShouldSendAnEmailContainingMessage(t *testing.T) {

	givenUrl := "/some/path"
	client := &ClientMock{t: t, StatusCode: http.StatusOK}

	msg := gowhisper.Message{
		To:      "julian.godesa@googlemail.com",
		From:    "gowhisper@mail.co0p.org",
		Subject: "hello there",
		Text:    "some data",
	}

	client.checkFn = func(req *http.Request, t *testing.T) {

		expHeader := "multipart/form-data; boundary="
		header := req.Header.Get("Content-Type")
		if !strings.HasPrefix(header, expHeader) {
			t.Errorf("expected header to start with '%s', got '%s'", expHeader, header)
		}

		actualUrl := req.URL.String()
		if givenUrl != actualUrl {
			t.Errorf("expected url to be '%s', got '%s'", givenUrl, actualUrl)
		}
		mr, _ := req.MultipartReader()
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				return
			}

			content, _ := ioutil.ReadAll(p)
			formName := p.FormName()
			if formName == "from" && string(content) != msg.From {
				t.Errorf("expected from value to be '%s', got '%s'", msg.From, content)
			}

			if formName == "to" && string(content) != msg.To {
				t.Errorf("expected from value to be '%s', got '%s'", msg.To, content)
			}

			if formName == "subject" && string(content) != msg.Subject {
				t.Errorf("expected from value to be '%s', got '%s'", msg.Subject, content)
			}

			if formName == "text" && string(content) != msg.Text {
				t.Errorf("expected from value to be '%s', got '%s'", msg.Text, content)
			}
		}
	}

	mailer := gowhisper.MailNotifier{ApiURL: givenUrl, Client: client}
	mailer.Send(msg)
}

func Test_SendShouldReturnErrorWithMessageIfStatusCodeWasNotOK(t *testing.T) {

	givenUrl := "/some/path"
	client := &ClientMock{t: t, StatusCode: http.StatusBadRequest}

	msg := gowhisper.Message{
		To:      "julian.godesa@googlemail.com",
		From:    "gowhisper@mail.co0p.org",
		Subject: "hello there",
		Text:    "some data",
	}

	mailer := gowhisper.MailNotifier{ApiURL: givenUrl, Client: client}
	err := mailer.Send(msg)

	if err == nil {
		t.Errorf("expected err not be nil")
	}
}
