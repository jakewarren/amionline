package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

var version string

type Config struct {
	dns    bool
	http   bool
	domain string
}

func main() {

	var app Config

	flag.BoolVar(&app.dns, "dns", false, "check dns resolution")
	flag.BoolVar(&app.http, "http", false, "check url for 200 status code")
	flag.StringVar(&app.domain, "domain", "httpbin.org", "domain name to query")
	displayVersion := flag.Bool("version", false, "display version")
	verbose := flag.Bool("verbose", false, "verbose output")
	flag.Parse()

	if *displayVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	var err error

	if app.dns {
		err = app.dnsCheck()
	} else if app.http {
		err = app.httpCheck()
	}

	if err != nil {
		if *verbose {
			fmt.Println(err)
		}
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func (c *Config) dnsCheck() error {
	_, err := net.LookupHost(c.domain)
	return err
}

func (c *Config) httpCheck() error {
	domain := c.domain

	if ok, _ := regexp.MatchString(`https?://`, domain); !ok {
		domain = "https://" + domain
	}

	resp, err := http.Get(domain)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s (%d)", "Invalid status code", resp.StatusCode))
	}

}
