package main

import (
	"net"
	"os"
	"flag"
	"fmt"
	"net/http"
	"regexp"
	"github.com/pkg/errors"
)

type Config struct{
	dns bool
	http bool
	domain string
}

func main() {

	var app Config

	flag.BoolVar(&app.dns,"dns",false,"check dns resolution")
	flag.BoolVar(&app.http,"http",false,"check url for 200 status code")
	flag.StringVar(&app.domain,"domain","httpbin.org","domain name to query")
	flag.Parse()

    var err error

    if app.dns{
    	err = app.dnsCheck()
	}else if app.http {
		err = app.httpCheck()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func (c *Config) dnsCheck() error{
	_, err := net.LookupHost(c.domain)
	return err
}

func (c *Config) httpCheck() error{
	domain := c.domain


	if ok,_ := regexp.MatchString(`https?://`,domain);!ok{
		domain = "https://"+domain
	}


	resp, err := http.Get(domain)
	if err != nil {
		return err
	}


	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	} else {
		return errors.New("Invalid status code")
	}

}
