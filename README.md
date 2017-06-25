![TC logo](https://raw.githubusercontent.com/maliceio/malice-team-cymru/master/logo.png)
# team-cymru [Deprecated] due to the service not working anymore :cry:

[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)
[![Docker Stars](https://img.shields.io/docker/stars/malice/team-cymru.svg)][hub]
[![Docker Pulls](https://img.shields.io/docker/pulls/malice/team-cymru.svg)][hub]
[![Image Size](https://img.shields.io/imagelayers/image-size/malice/team-cymru/latest.svg)](https://imagelayers.io/?images=malice/team-cymru:latest)
[![Image Layers](https://img.shields.io/imagelayers/layers/malice/team-cymru/latest.svg)](https://imagelayers.io/?images=malice/team-cymru:latest)

Malice TeamCymru - Malware Hash Registry Plugin

This repository contains a **Dockerfile** of TeamCymru's [Malware Hash Registry](http://www.team-cymru.org/MHR.html) for [Docker](https://www.docker.io/)'s [trusted build][hub] published to the public [DockerHub](https://hub.docker.com).

> **WARNING:** TeamCymru - Malware Hash Registry doesn't seem to be working on their end.  Please contact them and ask to get it back up and running

### Dependencies

* [gliderlabs/alpine](https://index.docker.io/_/gliderlabs/alpine/)


### Installation

1. Install [Docker](https://www.docker.io/).
2. Download [trusted build](https://hub.docker.com/r/malice/team-cymru/) from public [DockerHub](https://hub.docker.com): `docker pull malice/team-cymru`

### Usage

    docker run --rm malice/team-cymru (MD5|SHA1)

```bash
Usage: team-cymru [OPTIONS] COMMAND [arg...]

Malice TeamCymru - Malware Hash Registry Plugin

Version: v0.1.0, BuildTime: 20160228

Author:
  blacktop - <https://github.com/blacktop>

Options:
  --post, -p	POST results to Malice webhook [$MALICE_ENDPOINT]
  --proxy, -x	proxy settings for Malice webhook endpoint [$MALICE_PROXY]
  --table, -t	output as Markdown table
  --help, -h	show help
  --version, -v	print the version

Commands:
  help	Shows a list of commands or help for one command

Run 'team-cymru COMMAND --help' for more information on a command.
```

This will output to stdout and POST to malice results API webhook endpoint.

### Sample Output JSON:
```json
{
  "team-cymru": {
    "found": false,
    "lastseen": "",
    "detection": ""
  }
}
```
### Sample Output STDOUT (Markdown Table):
---
#### TeamCymru
| Detection   | LastSeen |
| ----------- | -------- |
| 85%         | 20091112 |
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
