package main

import (
	"fmt"
	"os"

	"netcli/helpers"
	"netcli/speedtest"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "Net-CLI",
		Usage:   "A lightweight network tool",
		Version: "v0.1.4",
		Authors: []*cli.Author{
			{
				Name:  "Jishnu Prasad K P",
				Email: "jishnu.prasad4@gmail.com",
			},
		},
	}
	// A common flag use under required commands
	defaultFlag := []cli.Flag{
		&cli.StringFlag{
			Name:        "host",
			Usage:       "Enter hostname",
			Required:    true,
			Value:       "",
			DefaultText: "github.com",
		},
	}

	speedTestFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "saving-mode",
			Usage: "Using less memory (≒10MB), though low accuracy (especially > 30Mbps).",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "json",
			Usage: "Output results in json format",
			Value: false,
		},
	}

	domainFlag := []cli.Flag{
		&cli.StringFlag{
			Name:        "domain",
			Required:    true,
			Value:       "",
			Aliases:     []string{"d"},
			Usage:       "Enter the domain name. Egs:- google.com",
			DefaultText: "google.com",
		},
	}

	urlFlag := []cli.Flag{
		&cli.StringFlag{
			Name:     "url",
			Required: true,
			Value:    "",
			Aliases:  []string{"u"},
			Usage:    "Enter the URL to shorten, Use quotes if your shell not parsing the URL.",
		},
	}

	emailFlag := []cli.Flag{
		&cli.StringFlag{
			Name:     "id",
			Required: true,
			Value:    "",
			Usage:    "Enter the Email address to validate",
		},
	}

	subnetFlag := []cli.Flag{
		&cli.StringFlag{
			Name:     "cidr",
			Required: true,
			Value:    "",
			Usage:    "Enter the CIDR block(e.g., 192.168.1.0/24)",
		},
		&cli.StringFlag{
			Name:     "ip",
			Required: false,
			Value:    "",
			Usage:    "Enter IP Address to check if it falls under provided CIDR",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "nameserver",
			Aliases: []string{"ns"},
			Usage:   "Print the nameserver of given hostname.(NS RECORD)",
			Flags:   defaultFlag,
			Action: func(c *cli.Context) error {
				nsIPs := helpers.GetNS(c.String("host"))
				for _, ip := range nsIPs {
					fmt.Println(ip.Host)
				}
				return nil
			},
		},
		{
			Name:    "mailserver",
			Aliases: []string{"mx"},
			Usage:   "Print the mail server of given hostname.(MX RECORD)",
			Flags:   defaultFlag,
			Action: func(c *cli.Context) error {
				mxIPs := helpers.GetMX(c.String("host"))
				for _, ip := range mxIPs {
					fmt.Println(ip.Host)
				}
				return nil
			},
		},
		{
			Name:  "a",
			Usage: "Print the host ip of given hostname.(A RECORD)",
			Flags: defaultFlag,
			Action: func(c *cli.Context) error {
				IPs := helpers.GetA(c.String("host"))
				for _, ip := range IPs {
					fmt.Println(ip)
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Print the canonical name of given hostname.(CNAME)",
			Flags: defaultFlag,
			Action: func(c *cli.Context) error {
				cname := helpers.GetCNAME(c.String("host"))
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "myip",
			Usage: "Print your public IP Address.",
			Action: func(c *cli.Context) error {
				fmt.Println("Your Public IP:- ", helpers.GetMyIP())
				return nil
			},
		},
		{
			Name:    "speedtest",
			Aliases: []string{"speed"},
			Usage:   "Do an internet speed test.",
			Flags:   speedTestFlags,
			Action: func(c *cli.Context) error {
				speedtest.SpeedTest(c.Bool("saving-mode"), c.Bool("json"))
				return nil
			},
		},
		{
			Name:    "subdomain",
			Aliases: []string{"subd"},
			Usage:   "Scans an entire domain to find as many subdomains as possible.",
			Flags:   domainFlag,
			Action: func(c *cli.Context) error {
				subDomains := helpers.GetSubdomains(c.String("domain"))
				for _, i := range subDomains {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:  "whois",
			Usage: "Get whois information of Domain Name or IP Addres.",
			Flags: defaultFlag,
			Action: func(c *cli.Context) error {
				fmt.Println(helpers.Whois(c.String("host")))
				return nil
			},
		},
		{
			Name:    "shorten",
			Aliases: []string{"short"},
			Usage:   "URL shortener to reduce a long link.",
			Flags:   urlFlag,
			Action: func(c *cli.Context) error {
				fmt.Println(helpers.UrlShorten(c.String("url")))
				return nil
			},
		},
		{
			Name:  "email",
			Usage: "Check if email address is valid.",
			Flags: emailFlag,
			Action: func(c *cli.Context) error {
				result := helpers.ValidateEmail(c.String("id"))
				fmt.Printf("Email: %s\n", result["email"])
				fmt.Printf("  Valid Format: %t\n", result["validFormat"])
				fmt.Printf("  Has MX Records: %t\n", result["hasMX"])
				fmt.Printf("  SMTP Check: %v\n", result["smtpCheck"])
				fmt.Printf("  Potentially Reachable: %t\n\n", result["reachable"])
				return nil
			},
		},
		{
			Name:  "subnet",
			Usage: "Calculate details about provided CIDR block.",
			Flags: subnetFlag,
			Action: func(c *cli.Context) error {
				result := helpers.CalculateSubnet(c.String("cidr"))
				fmt.Println("\n=== CIDR IP Range Analysis ===")
				fmt.Printf("CIDR Notation: %s\n", result["cidr"])
				fmt.Printf("Network Address: %s\n", result["networkAddr"])
				fmt.Printf("Usable Host IP Range: %s - %s\n", result["firstUsableIP"], result["lastUsableIP"])
				fmt.Printf("Broadcast Address: %s\n", result["broadcast"])
				fmt.Printf("Total Number of Hosts: %.0f\n", result["totalHosts"])
				fmt.Printf("Number of Usable Hosts: %.0f\n", result["usableHosts"])
				fmt.Printf("Subnet Mask: %s\n", result["subnetMask"])
				fmt.Printf("IP Class: %s\n", result["ipClass"])
				fmt.Printf("IP Type: %s\n", result["ipType"])
				if c.String("ip") != "" {
					ipCheck := helpers.IsIPInCIDR(c.String("ip"), c.String("cidr"))
					if ipCheck {
						fmt.Printf("\nIP Check Result: \n\t%s is WITHIN the %s range.\n", c.String("ip"), c.String("cidr"))
					} else {
						fmt.Printf("\nIP Check Result: \n\t%s is NOT within the %s range.\n", c.String("ip"), c.String("cidr"))
					}
				}
				return nil
			},
		},
		{
			Name:    "header",
			Usage:   "Enter the web address to check http header",
			Aliases: []string{"head"},
			Flags:   urlFlag,
			Action: func(c *cli.Context) error {
				helpers.GetHeader(c.String("url"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	helpers.ErrorHandler(err)
}
