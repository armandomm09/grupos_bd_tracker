package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {

	from := "pabloarmando.macbeath@gmail.com"
	pass := "jxyp kmnl ijxv gvpc"
	to := "binizavazquez@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		"Hola biniza"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("Sent")
}
