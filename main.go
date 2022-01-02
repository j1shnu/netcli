package main

import (
	"fmt"
	"os"

	"net-cli/helpers"
	"net-cli/speedtest"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "Net-CLI",
		Usage:   "Tool to get net infos from CLI",
		Version: "v0.1",
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

	app.Commands = []*cli.Command{
		{
			Name:    "ns",
			Aliases: []string{"nameserver"},
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
			Name:    "mx",
			Aliases: []string{"mailserver"},
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
			Name:  "speed",
			Usage: "Do an internet speed test",
			Flags: speedTestFlags,
			Action: func(c *cli.Context) error {
				speedtest.SpeedTest(c.Bool("saving-mode"), c.Bool("json"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	helpers.ErrorHandler(err)
}
