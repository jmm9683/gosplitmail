package gosplitmail

import (
	"bytes"
	"html/template"

	"gopkg.in/gomail.v2"
)

type EmailData struct {
	ThreadTitle  string
	ThreadList   string
	ImagePaths   []string // paths to images
	ImageCIDs    []string // CIDs for inline embedding
	ThreadFilter map[string]any
}

func SplitEmail(bodyTemplate, subject, from string, groups []EmailData) []*gomail.Message {
	tmpl, err := template.New("email").Parse(bodyTemplate)
	if err != nil {
		panic(err)
	}

	messages := make([]*gomail.Message, len(groups))
	for i, data := range groups {
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			panic(err)
		}
		message := gomail.NewMessage()
		message.SetHeader("From", from)
		message.SetHeader("To", data.ThreadList)
		message.SetHeader("Subject", subject)
		message.SetBody("text/html", buf.String())
		for j := range data.ImagePaths {
			if data.ImagePaths[j] != "" && data.ImageCIDs[j] != "" {
				message.Embed(data.ImagePaths[j], gomail.SetHeader(map[string][]string{
					"Content-ID": {"<" + data.ImageCIDs[j] + ">"},
				}))
			}
		}
		messages[i] = message
	}
	return messages
}
