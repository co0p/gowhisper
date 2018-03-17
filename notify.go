package gowhisper

import (
	"bytes"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MailNotifier struct {
	ApiURL string
	Client HttpClient
}

type Notifier interface {
	Send(Message) error
}

type Message struct {
	From    string
	To      string
	Subject string
	Text    string
}

func (m *MailNotifier) Send(msg Message) error {

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	from, _ := w.CreateFormField("from")
	from.Write([]byte(msg.From))

	to, _ := w.CreateFormField("to")
	to.Write([]byte(msg.To))

	subject, _ := w.CreateFormField("subject")
	subject.Write([]byte(msg.Subject))

	text, _ := w.CreateFormField("text")
	text.Write([]byte(msg.Text))

	w.Close()

	req, err := http.NewRequest("POST", m.ApiURL, buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := m.Client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		return errors.New("failed to send mail: " + bodyString)
	}

	return nil
}
