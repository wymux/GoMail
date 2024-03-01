package main

import (
	"net/smtp"
	"testing"
)

type Mailer interface {
	SendMail(addr string, a smtp.Auth, from string, to[] string, msg []byte) error
}

type RealMailer struct{}

func (m *RealMailer) SendMailer(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

type MockMailer struct {
	Addr string
	Auth smtp.Auth
	From string
	To   []string
	Msg  []byte
}

func (m *MockMailer) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	m.Addr = addr
	m.Auth = a
	m.From = from
	m.To   = to
	m.Msg  = msg
	return nil
}

func TestSendMail(t *testing.T) {
	mailer := &MockMailer{}
	sendMail(mailer)
	if mailer.From != "wymux@shoash.com" {
		t.Errorf("Expected wymux@shoash.com, got %s")
	}
}

func sendMail(mailer Mailer) {
	mailer.SendMail(mailer)
}
