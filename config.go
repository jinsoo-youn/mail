package main

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SmtpInfo `yaml:"smtpInfo"`
	Message  `yaml:"message"`
}

type SmtpInfo struct {
	SmptNo       int    `yaml:"smtp_no"`
	SmtpHost     string `yaml:"smtp_host"`
	SmtpPort     int    `yaml:"smtp_port"`
	SmtpProtocol string `yaml:"smtp_protocol"`
	SmtpTls      string `yaml:"smtp_tls"`
	SmtpTimeout  string `yaml:"smtp_timeout`
	SmtpId       string `yaml:"smtp_id"`
	SmtpPw       string `yaml:"smtp_pw"`
	SmtpIdentity string `yaml:"smtp_identity"`
}

type Message struct {
	From    string   `yaml:"from"`
	To      []string `yaml:"to"`
	Subject string   `yaml:"subject"`
	Body    string   `yaml:"body"`
	Attach  []string `yaml:"attach"`
}

func SetFlagsFromConfig(fs *flag.FlagSet, filename string) (err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	config := Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	addSmtpInfo(fs, &config.SmtpInfo)
	addMessage(fs, &config.Message)

	return nil
}

func addSmtpInfo(fs *flag.FlagSet, smtpInfo *SmtpInfo) (err error) {
	// if smtpInfo.SmptNo != 0 {
	// 	fs.Set("", smtpInfo.SmptNo, "dfdf")
	// 	fs.Set("test",)
	// 	fs.set
	// }

	return nil
}

func addMessage(fs *flag.FlagSet, message *Message) {
	if len(message.To) > 0 {

	}

}
