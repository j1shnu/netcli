package helpers

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetNS(url string) []*net.NS {
	ips, err := net.LookupNS(url)
	ErrorHandler(err)
	return ips
}

func GetMX(url string) []*net.MX {
	ips, err := net.LookupMX(url)
	ErrorHandler(err)
	return ips
}

func GetA(url string) []net.IP {
	ips, err := net.LookupIP(url)
	ErrorHandler(err)
	return ips
}

func GetCNAME(url string) string {
	host, err := net.LookupCNAME(url)
	ErrorHandler(err)
	return host
}

func GetMyIP() string {
	res, err := http.Get("http://ifconfig.me")
	ErrorHandler(err)
	defer res.Body.Close()
	myIP, err := ioutil.ReadAll(res.Body)
	ErrorHandler(err)
	return string(myIP)
}
