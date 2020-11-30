package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

// Return application command line to execute
func Generate2() *cli.App {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	commands := []*cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs",
			Flags:  flags,
			Action: searchIPV2,
		},
		{
			Name:   "servers",
			Usage:  "Search name server in internet",
			Flags:  flags,
			Action: searchServerV2,
		},
	}
	app := &cli.App{
		Name:     "Command Line Application",
		Usage:    "Search for IPs and Server Names on the Internet",
		Flags:    flags,
		Commands: commands,
	}
	return app
}

func searchIPV2(c *cli.Context) error {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

func searchServerV2(c *cli.Context) error {
	host := c.String("host")

	servers, erro := net.LookupNS(host) // name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
	return nil
}
