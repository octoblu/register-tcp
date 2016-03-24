package healthchecker

import (
	"fmt"
	"log"
	"net"
	"net/url"
)

// Healthy returns true if the TCP request
// succeeds, false otherwise
func Healthy(uriStr string) bool {
	uri, err := url.Parse(uriStr)
	if err != nil {
		log.Fatalln("Error parsing url", err.Error())
	}

	conn, err := net.Dial("tcp", uri.Host)
	if err != nil {
		fmt.Println("net.Dial err:", err.Error())
		return false
	}
	conn.Close()
	return true
}
