package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"gopkg.in/gomail.v2"
)

func main() {

	photoPath := filepath.Join("./Photos")
	jpgPattern := filepath.Join(photoPath, "*.jpg")

	photos, err := filepath.Glob(jpgPattern)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uniqueImageMap := map[int64]bool{} //tbd: sqlite

	for _, photo := range photos {
		imageDescription := ""
		f, err := os.Open(photo)
		if err != nil {
			log.Fatal(err)
		}

		exif.RegisterParsers(mknote.All...)

		x, err := exif.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		tm, _ := x.DateTime()
		if uniqueImageMap[tm.Unix()] {
			fmt.Println("Duplicate image: ", photo, tm.Unix())
			continue
		} else {
			uniqueImageMap[tm.Unix()] = true //tbd: sqlite
		}
		year, month, day := tm.Date()
		fmt.Printf("Taken on %v %v, %v\n", month, day, year)
		imageDescriptionTiff, _ := x.Get(exif.ImageDescription)
		imageDescription = strings.ReplaceAll(imageDescriptionTiff.String(), "\"", "")
		fmt.Println(imageDescription)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "SENDER@SENDER.COM")
	m.SetHeader("To", "RECIPIENT@RECIPIENT.COM")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	// d := gomail.NewDialer("smtp.mail.me.com", 587, "MEICLOUD@ME.COM", "TOlKENTIME")

	// // Send the email to Bob, Cora and Dan.
	// if err := d.DialAndSend(m); err != nil {
	// 	panic(err)
	// }

}
