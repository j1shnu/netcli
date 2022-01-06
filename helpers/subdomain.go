package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type domainData struct {
	NameValue  string   `json:"name_value"`
	Subdomains []string `json:"subdomains"`
}

func domainValidate(domain string) bool {
	pattern := `^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\\.[a-zA-Z]{2,3})$`
	regExp := regexp.MustCompile(pattern)
	return regExp.MatchString(domain)
}

func subdomainValidate(subdomain, domain string) bool {
	pattern := fmt.Sprintf("([-a-z0-9]+).%v", domain)
	regExp := regexp.MustCompile(pattern)
	return regExp.MatchString(subdomain)

}

func GetSubdomains(domain string) []string {
	if !domainValidate(domain) {
		log.Fatal("Invalid Domain Name.")
	}
	subDomains := fetchSubdomains(domain)
	return removeDuplicates(subDomains, domain)
}

func getData(baseUrl string) []byte {
	resp, err := http.Get(baseUrl)
	ErrorHandler(err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	ErrorHandler(err)
	return data
}

func fetchSubdomains(domain string) []string {
	data := getData(fmt.Sprintf("https://crt.sh/?q=%v&output=json", domain))
	crtDomainDatas := make([]domainData, 0)
	json.Unmarshal([]byte(string(data)), &crtDomainDatas)
	var crtSubDomains []string
	for _, subdomains := range crtDomainDatas {
		split_subd := strings.Split(subdomains.NameValue, "\n")
		crtSubDomains = append(crtSubDomains, split_subd...)
	}
	api := getData("https://pastebin.com/raw/9nYue4Dh")
	baseURL := "https://www.virustotal.com/vtapi/v2/domain/report?apikey=" + string(api) + "&domain=" + domain
	data1 := getData(baseURL)
	var vtDomainDatas domainData
	json.Unmarshal(data1, &vtDomainDatas)
	return append(crtSubDomains, vtDomainDatas.Subdomains...)
}

func removeDuplicates(subDomains []string, domain string) []string {
	keys := make(map[string]bool)
	var subdomains []string
	for _, subdomain := range subDomains {
		if _, val := keys[subdomain]; !val {
			keys[subdomain] = true
			if subdomainValidate(subdomain, domain) {
				subdomains = append(subdomains, subdomain)
			}
		}
	}
	return subdomains
}
