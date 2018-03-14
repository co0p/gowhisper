package gowhisper

import (
	"bytes"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"time"
)

type MailNotifier struct {
	apiURL string
	client *http.Client
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

func NewMailNotifier(uri string) MailNotifier {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	return MailNotifier{apiURL: uri, client: client}
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

	req, err := http.NewRequest("POST", m.apiURL, buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := m.client.Do(req)

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
