package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/pkg/errors"
)

var version = "(ﾉ☉ヮ⚆)ﾉ ⌒*:･ﾟ✧"

type Config struct {
	dns     bool
	http    bool
	domain  string
	timeout time.Duration
	verbose bool
}

func main() {
	var app Config

	flag.BoolVar(&app.dns, "dns", false, "check dns resolution")
	flag.BoolVar(&app.http, "http", false, "check url for 200 status code")
	flag.StringVar(&app.domain, "domain", "httpbin.org", "domain name to query")
	flag.DurationVar(&app.timeout, "timeout", time.Second*30, "http timeout")

	displayVersion := flag.Bool("version", false, "display version")
	flag.BoolVar(&app.verbose, "verbose", false, "verbose output")
	flag.Parse()

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000-0700"
	if app.verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	if *displayVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	var err error

	if app.http {
		err = app.httpCheck()
	} else { // default to a DNS check
		err = app.dnsCheck()
	}

	if err != nil {
		if app.verbose {
			log.Error().Err(err).Msg("received error")
		}
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func (c *Config) dnsCheck() error {
	ctx := context.Background()
	resolver := &net.Resolver{
		Dial:     cloudflareDNSDialer,
		PreferGo: true,
	}
	log.Debug().Str("domain", c.domain).Msg("issuing DNS request")
	ips, err := resolver.LookupHost(ctx, c.domain)
	log.Debug().Str("domain", c.domain).Strs("ips", ips).Msg("received DNS response")
	return err
}

// cloudflareDNSDialer Dialer pointing to the DNS resolver from cloudflare.
func cloudflareDNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "udp", "1.1.1.1:53")
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

	log.Debug().Msg("issuing HTTP request")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	log.Debug().Int("status code", resp.StatusCode).Msg("received response")

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s (%d)", "Invalid status code", resp.StatusCode))
	}
}
