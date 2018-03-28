# amionline
[![GitHub release](http://img.shields.io/github/release/jakewarren/amionline.svg?style=flat-square)](https://github.com/jakewarren/amionline/releases])
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/amionline/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/amionline)](https://goreportcard.com/report/github.com/jakewarren/amionline)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)

> Check network connectivity and report via exit code

## Install
### Option 1: Binary

Download the latest release from [https://github.com/jakewarren/amionline/releases/latest](https://github.com/jakewarren/amionline/releases/latest)

### Option 2: From source

```
go get github.com/jakewarren/amionline
```

## Usage

Usage in a bash one-liner:
```
amionline -http && echo "connected" || echo "offline"
```

```
❯ amionline -h
Usage of amionline:
  -dns
    	check dns resolution
  -domain string
    	domain name to query (default "httpbin.org")
  -http
    	check url for 200 status code
  -verbose
    	verbose output
  -version
    	display version
```

## Background

I created this utility for use in automated jobs that require internet connectivity to work; this program makes it easy to add connectivity checking to other programs.

## Changes

All notable changes to this project will be documented in the [changelog].

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this project adheres to [Semantic Versioning](http://semver.org/).

## License

MIT © 2018 Jake Warren

[changelog]: https://github.com/jakewarren/amionline/blob/master/CHANGELOG.md
