package iphistory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type iphistory struct {
	Query struct {
		Tool   string `json:"tool,omitempty"`
		Domain string `json:"domain,omitempty"`
	} `json:"query,omitempty"`
	Response struct {
		Records []struct {
			IP       string `json:"ip,omitempty"`
			Location string `json:"location,omitempty"`
			Owner    string `json:"owner,omitempty"`
			Lastseen string `json:"lastseen,omitempty"`
		} `json:"records,omitempty"`
	} `json:"response,omitempty"`
}

var data iphistory

func DomainIPhistory(domain, api string) iphistory {
	urls := fmt.Sprintf("https://api.viewdns.info/iphistory/?domain=%s&apikey=%s&output=json", domain, api)
	resp, err := http.Get(urls)
	if err != nil {
		fmt.Println(err.Error())
		recover()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &data)

	return data
}
