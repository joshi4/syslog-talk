package main

import (
	"crypto/tls"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// START OMIT
	sw, err := tls.Dial("tcp", "logs4.papertrailapp.com:27193", nil)
	if err != nil {
		log.Println(err)
	} else {
		log.SetPrefix("gophercon2016: ")
		log.SetOutput(io.MultiWriter(sw, os.Stderr))
	}
	//END OMIT
	for {
		log.Println("Hello Gophers!")
		time.Sleep(1 * time.Second)
	}
}
