package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"regexp"
	"strings"
)

func main() {
	host := "life.ru"
	addr := fmt.Sprintf("%s%s%d", host, ":", 443)
	header := ""

	var s net.Conn
	var err error

	cfg := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host, //simple fix
	}
	if s, err = tls.Dial("tcp", addr, cfg); err == nil {
		request := ""
		request += "GET /"
		header += " HTTP/1.1\r\nHost: "
		header += addr + "\r\n"

		request += header + "\r\n"

		if _, err := s.Write([]byte(request)); err == nil {
			tmp := make([]byte, 256)
			if _, err := s.Read(tmp); err == nil {
				s := string(tmp[:])
				fmt.Printf("Response: %s\n", s)
				strings.Contains("something", "some")
				if matched, _ := regexp.MatchString("HTTP/\\d\\.\\d\\s2\\d+", s); matched == true {
					fmt.Println("OK")
				} else {
					fmt.Println("BAD")
				}
			}
		}
		s.Close()
	}
}
