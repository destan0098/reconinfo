package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
	"reconinfo/pkg/iphistory"
	"reconinfo/pkg/portscanner"
	"reconinfo/pkg/reversip"
)

type apikey struct {
	APIkey string `yaml:"APIkey,omitempty"`
}

var data *apikey
var website string
var iphi, portsb, revbb int

//var portss Ports

func ReadConfig() *apikey {
	f, err := os.Open("./config/api.yaml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&data)

	if err != nil {
		fmt.Println(err)
	}
	return data
}
func main() {
	apikeys := ReadConfig()
	fmt.Println(`


  _____                                    _____            __         
 |  __ \                                  |_   _|          / _|        
 | |__) |   ___    ___    ___    _ __       | |    _ __   | |_    ___  
 |  _  /   / _ \  / __|  / _ \  | '_ \      | |   | '_ \  |  _|  / _ \ 
 | | \ \  |  __/ | (__  | (_) | | | | |    _| |_  | | | | | |   | (_) |
 |_|  \_\  \___|  \___|  \___/  |_| |_|   |_____| |_| |_| |_|    \___/ 
                                                                       
                                                                       

 
`)
	fmt.Println(color.Colorize(color.Red, "[*] This tool is for training."))
	app := &cli.App{
		Name:  "Recon Info",
		Usage: "This Tool for recon And Just For Training",
		Commands: []*cli.Command{

			{
				Name:    "scan",
				Aliases: []string{"s"},
				Usage:   "Scan command",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "url",
						Value:       "test.com",
						Usage:       "Url Address to scan",
						Aliases:     []string{"u"},
						Destination: &website,
					},
					&cli.BoolFlag{
						Name:    "iphistory",
						Aliases: []string{"i"},
						Usage:   "show ip history for domain",
						Count:   &iphi,
					},
					&cli.BoolFlag{
						Name:    "portscan",
						Aliases: []string{"p"},
						Usage:   "show Ports Sattus",
						Count:   &portsb,
					},
					&cli.BoolFlag{
						Name:    "reverse",
						Aliases: []string{"r"},
						Usage:   "Reverse IP and domain",
						Count:   &revbb,
					},
				},
				Action: func(c *cli.Context) error {
					inputFile := c.String("url")
					if inputFile == "test.com" {

						fmt.Println(color.Colorize(color.Green, "Please Enter Website Address "))
						err := cli.ShowAppHelp(c)
						if err != nil {
							return err
						}
						return nil

					} else {
						if iphi > 0 {
							iph := iphistory.DomainIPhistory("example.com", apikeys.APIkey)
							for _, Records := range iph.Response.Records {
								fmt.Printf(color.Colorize(color.Green, "[+] IP : %s , Location : %s , IP Owner : %s , Last Seen : %s \n"), Records.IP, Records.Location, Records.Owner, Records.Lastseen)
							}
						}
						if portsb > 0 {
							ports := portscanner.PortScanner("example.com", apikeys.APIkey)

							for _, portsss := range ports.Response.Port {
								fmt.Printf(color.Colorize(color.Green, "[+] Port Num: %s , Port Status : %s , Port Service : %s  \n"), portsss.Number, portsss.Status, portsss.Service)
							}
						}
						if revbb > 0 {
							reversedip := reversip.ReverseIP("example.com", apikeys.APIkey)
							for _, reversedipdomain := range reversedip.Response.Domains {
								//	fmt.Println(reversedipdomain.Name)
								fmt.Printf(color.Colorize(color.Green, "[+] Domain name %s  \n"), reversedipdomain.Name)

							}
						}
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
