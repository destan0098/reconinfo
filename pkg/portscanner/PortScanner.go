package portscanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Portscan struct {
	Query struct {
		Tool string `json:"tool,omitempty"`
		Host string `json:"host,omitempty"`
	} `json:"query,omitempty"`
	Response struct {
		Port []struct {
			Number  string `json:"number,omitempty"`
			Service string `json:"service,omitempty"`
			Status  string `json:"status,omitempty"`
		} `json:"port,omitempty"`
	} `json:"response,omitempty"`
}

var data Portscan

func PortScanner(domain, api string) Portscan {
	urls := fmt.Sprintf("https://api.viewdns.info/portscan/?host=%s&apikey=%s&output=json", domain, api)
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
