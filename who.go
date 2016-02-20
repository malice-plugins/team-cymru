package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	answers, err := net.LookupTXT("44d88612fea8a8f36de82e1278abb02f.malware.hash.cymru.com")
	fmt.Printf("%#v", answers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answers)
}
