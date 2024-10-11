package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

func main() {

	wac, err := whatsapp.NewConn(20 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	qrChan := make(chan string)
	go func() {
		fmt.Printf("qr code: %v\n", <-qrChan)
		//show qr code or save it somewhere to scan
	}()

	sess, err := wac.Login(qrChan)
	if err != nil {
		log.Fatal(err)
	}

	newSess, err := wac.RestoreWithSession(sess)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newSess.Wid)

	// from := "pabloarmando.macbeath@gmail.com"
	// pass := "jxyp kmnl ijxv gvpc"
	// to := "binizavazquez@gmail.com"

	// msg := "From: " + from + "\n" +
	// 	"To: " + to + "\n" +
	// 	"Subject: Hello there\n\n" +
	// 	"Hola biniza"

	// err := smtp.SendMail("smtp.gmail.com:587",
	// 	smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
	// 	from, []string{to}, []byte(msg))

	// if err != nil {
	// 	log.Printf("smtp error: %s", err)
	// 	return
	// }

	// fmt.Println("Sent")
}
