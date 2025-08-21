package main

import (
	"fmt"

	"github.com/jakemorrissey/go-splitmail/gosplitmail"
)

func main() {
	// Email template using all fields
	bodyTemplate := `
    <h2>{{.ThreadTitle}}</h2>
    <p>Recipients: {{.ThreadList}}</p>
    {{if .ThreadFilter.Family}}<p>Hi fam!!</p>{{end}}
    {{if .ThreadFilter.Friends}}<p>Hey frens!</p>{{end}}
    {{if .ThreadFilter.Colleague}}<p>Dear Colleagues, </p>{{end}}
    {{range $i, $cid := .ImageCIDs}}
        <img src="cid:{{$cid}}" alt="Photo {{$i}}">
    {{end}}
    `

	// Example EmailData for each group
	groups := []gosplitmail.EmailData{
		{
			ThreadTitle: "To the family",
			ThreadList:  "dad@example.com",
			ImagePaths:  []string{"./Photos/dad1.jpg", "./Photos/dad2.jpg"},
			ImageCIDs:   []string{"dadphoto1", "dadphoto2"},
			ThreadFilter: map[string]any{
				"Dad": true,
			},
		},
		{
			ThreadTitle: "To the friends",
			ThreadList:  "friend@example.com",
			ImagePaths:  []string{"./Photos/friend1.jpg"},
			ImageCIDs:   []string{"friendphoto1"},
			ThreadFilter: map[string]any{
				"Friend": true,
			},
		},
		{
			ThreadTitle: "To the colleagues",
			ThreadList:  "colleague@example.com",
			ImagePaths:  []string{"./Photos/work1.jpg", "./Photos/work2.jpg"},
			ImageCIDs:   []string{"workphoto1", "workphoto2"},
			ThreadFilter: map[string]any{
				"Colleague": true,
			},
		},
	}

	subject := "Weekly Recap"
	from := "me@example.com"

	messages := gosplitmail.SplitEmail(bodyTemplate, subject, from, groups)

	// Example: Print the email bodies to stdout
	for i, msg := range messages {
		fmt.Printf("Email #%d to %s:\n%s\n\n", i+1, groups[i].ThreadList, msg.GetBody("text/html"))
	}

	// To actually send, use gomail.Dialer (not shown here)
	// d := gomail.NewDialer("smtp.example.com", 587, "user", "pass")
	// if err := d.DialAndSend(messages...); err != nil {
	//     fmt.Println("Send error:", err)
	//     os.Exit(1)
}
