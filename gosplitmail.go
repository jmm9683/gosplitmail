package gosplitmail

import (
	"gopkg.in/gomail.v2"
)

func SplitEmail(body, subject, from string, groups []string) []*gomail.Message {
	messages := make([]*gomail.Message, len(groups))
	for i := range len(groups) {
		message := gomail.NewMessage()
		message.SetHeader("From", from)
		message.SetHeader("To", groups[i])
		message.SetHeader("Subject", subject)
		message.SetBody("text/html", body)
		messages[i] = message
	}
	return messages

}
