![TC logo](https://raw.githubusercontent.com/maliceio/malice-team-cymru/master/logo.png)
# malice-team-cymru

[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)
[![Docker Stars](https://img.shields.io/docker/stars/malice/teamcymru.svg)][hub]
[![Docker Pulls](https://img.shields.io/docker/pulls/malice/teamcymru.svg)][hub]
[![Image Size](https://img.shields.io/imagelayers/image-size/malice/teamcymru/latest.svg)](https://imagelayers.io/?images=malice/teamcymru:latest)
[![Image Layers](https://img.shields.io/imagelayers/layers/malice/teamcymru/latest.svg)](https://imagelayers.io/?images=malice/teamcymru:latest)

Malice TeamCymru Hash Lookup Plugin

This repository contains a **Dockerfile** of **malice/team-cymru** for [Docker](https://www.docker.io/)'s [trusted build](https://index.docker.io/u/malice/teamcymru/) published to the public [DockerHub](https://index.docker.io/).

> **WARNING:** Work in progress.  Not ready yet.

### Dependencies

* [gliderlabs/alpine:3.3](https://index.docker.io/_/gliderlabs/alpine/)


### Installation

1. Install [Docker](https://www.docker.io/).
2. Download [trusted build](https://hub.docker.com/r/malice/team-cymru/) from public [DockerHub](https://hub.docker.com): `docker pull malice/team-cymru`

### Usage

    docker run --rm malice/team-cymru (MD5|SHA1)

```bash
Usage: team-cymru [OPTIONS] COMMAND [arg...]

Malice TeamCymru Hash Lookup Plugin

Version: v0.1.0, BuildTime: 20160214

Author:
  blacktop - <https://github.com/blacktop>

Options:
  --table, -t	output as Markdown table
  --post, -p	POST results to Malice webhook [$MALICE_ENDPOINT]
  --proxy, -x	proxy settings for Malice webhook endpoint [$MALICE_PROXY]
  --help, -h	show help
  --version, -v	print the version

Commands:
  help		Shows a list of commands or help for one command

Run 'team-cymru COMMAND --help' for more information on a command.
```

This will output to stdout and POST to malice results API webhook endpoint.

### Sample Output JSON:
```json
{
  "team-cymru": {

  }
}
```
### Sample Output STDOUT (Markdown Table):
---
#### teamcymru
| Ratio   | Scanned                |
| ------- | ---------------------- |
| 85%     | Sun 2016Feb14 14:00:50 |
---
### To Run on OSX
 - Install [Homebrew](http://brew.sh)

```bash
$ brew install caskroom/cask/brew-cask
$ brew cask install virtualbox
$ brew install docker
$ brew install docker-machine
$ docker-machine create --driver virtualbox malice
$ eval $(docker-machine env malice)
```

### Documentation

### Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/maliceio/malice-team-cymru/issues/new) and I'll get right on it.

### Credits

### License
MIT Copyright (c) 2016 **blacktop**

[hub]: https://hub.docker.com/r/malice/team-cymru/
