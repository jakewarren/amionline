package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"

	"github.com/pkg/errors"
	"time"
)

var version string

type Config struct {
	dns     bool
	http    bool
	domain  string
	timeout time.Duration
}

func main() {

	var app Config

	flag.BoolVar(&app.dns, "dns", false, "check dns resolution")
	flag.BoolVar(&app.http, "http", false, "check url for 200 status code")
	flag.StringVar(&app.domain, "domain", "httpbin.org", "domain name to query")
	flag.DurationVar(&app.timeout, "timeout", time.Second*30, "http timeout")

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

	req, err := http.NewRequest("GET", domain, nil)
	if err != nil {
		return err
	}

	httpClient := &http.Client{
		Timeout: c.timeout,
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s (%d)", "Invalid status code", resp.StatusCode))
	}

}
