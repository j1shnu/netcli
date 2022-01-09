package helpers

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/likexian/whois"
)

func ErrorHandler(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Get data from URL or API
func GetData(baseUrl string) []byte {
	resp, err := http.Get(baseUrl)
	ErrorHandler(err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	ErrorHandler(err)
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, string(data))
		os.Exit(2)
	}
	return data
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
	myIP := GetData("http://ifconfig.me")
	return string(myIP)
}

func Whois(host string) string {
	result, err := whois.Whois(host)
	ErrorHandler(err)
	return result
}

func UrlShorten(longUrl string) string {
	baseUrl := "http://is.gd/api.php?longurl=" + url.QueryEscape(longUrl)
	isGDUrl := GetData(baseUrl)
	return string(isGDUrl)
}
