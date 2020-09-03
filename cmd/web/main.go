package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"

	"github.com/jinsoo-youn/mail/password"
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
)

func main() {

	var key = []byte("tmaxcloudck2-1js")
	a, _ := password.NewAesCipher(key)
	pw := a.DecryptString("yFgOiKqHci9M8a29i9HFgx-81mBy8XwUBA==")
	fmt.Println("real password=", pw)

	content, err := ioutil.ReadFile("./static/config/config.yml")
	// content, err := ioutil.ReadFile("https://raw.githubusercontent.com/jinsoo-youn/mail/master/config.yaml")
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	d := gomail.NewDialer(config.SmtpInfo.SmtpHost, config.SmtpInfo.SmtpPort, config.SmtpInfo.SmtpId, a.DecryptString(config.SmtpInfo.SmtpPw))

	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
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

	m.SetBody("text/plain", config.Body)
	m.Attach(config.Attach)

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
