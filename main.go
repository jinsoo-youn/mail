package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
)

func main() {

	// content, err := ioutil.ReadFile("C:/Users/jinsoo_youn/go/src/github.com/jinsoo-youn/mail/config.yaml")
	content, err := ioutil.ReadFile("https://github.com/jinsoo-youn/mail/config.yaml")
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}

	// pw := []byte("52382fbb9879096561ec0a61845c0a394ffdc0396bbe97c45236d7e6dcb93b2fdea879db30")

	// d := gomail.NewDialer("mail.tmax.co.kr", 587, "jinsoo_youn@tmax.co.kr", string(pwr))

	d := gomail.NewDialer(config.SmtpInfo.SmtpHost, config.SmtpInfo.SmtpPort, config.SmtpInfo.SmtpId, config.SmtpInfo.SmtpPw)

	// d.TLSConfig = &tls.Config{
	// 	InsecureSkipVerify: true,
	// 	// ServerName:         "587",
	// }

	s, err := d.Dial()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.From)
	m.SetHeader("To", config.To...)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", config.Subject)
	m.SetBody("text/html", config.Body)
	// m.Attach("/home/Alex/lolcat.jpg")

	if err := gomail.Send(s, m); err != nil {
		panic(err)
	}
	m.Reset()

	// s, err := d.Dial()
	// if err != nil {
	// 	panic(err)
	// }
	// gomail.Send(s, m)

}