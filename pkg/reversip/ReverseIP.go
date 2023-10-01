package reversip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReverseIPs struct {
	Query struct {
		Tool string `json:"tool,omitempty"`
		Host string `json:"host,omitempty"`
	} `json:"query,omitempty"`
	Response struct {
		DomainCount string `json:"domain_count,omitempty"`
		Domains     []struct {
			Name         string `json:"name,omitempty"`
			LastResolved string `json:"last_resolved,omitempty"`
		} `json:"domains,omitempty"`
	} `json:"response,omitempty"`
}

var data ReverseIPs

func ReverseIP(domain, api string) ReverseIPs {
	urls := fmt.Sprintf("https://api.viewdns.info/reverseip/?host=%s&apikey=%s&output=json", domain, api)
	resp, err := http.Get(urls)
	if err != nil {
		fmt.Println(err.Error())
		recover()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data
	}

	return data
}
