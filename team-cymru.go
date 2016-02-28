package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/crackcomm/go-clitable"
	"github.com/levigross/grequests"
	"github.com/parnurzeal/gorequest"
)

// Version stores the plugin's version
var Version string

// BuildTime stores the plugin's build time
var BuildTime string

// TeamCymru json object
type TeamCymru struct {
	Results ResultsData `json:"team-cymru"`
}

// ResultsData json object
type ResultsData struct {
	Found     bool   `json:"found"`
	LastSeen  string `json:"lastseen"`  // last known GMT timestamp associated with that hash in unix epoch
	Detection string `json:"detection"` // detection percentage across a mix of AV packages
}

type notFound struct {
	Found bool `json:"found"`
}

func getopt(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dfault
	}
	return value
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func hashType(hash string) *grequests.RequestOptions {
	if match, _ := regexp.MatchString("([a-fA-F0-9]{32})", hash); match {
		return &grequests.RequestOptions{
			Params: map[string]string{
				"md5": hash,
			},
		}
	} else if match, _ := regexp.MatchString("([a-fA-F0-9]{40})", hash); match {
		return &grequests.RequestOptions{
			Params: map[string]string{
				"sha1": hash,
			},
		}
	} else if match, _ := regexp.MatchString("([a-fA-F0-9]{64})", hash); match {
		return &grequests.RequestOptions{
			Params: map[string]string{
				"sha256": hash,
			},
		}
	} else if match, _ := regexp.MatchString("([a-fA-F0-9]{128})", hash); match {
		return &grequests.RequestOptions{
			Params: map[string]string{
				"sha512": hash,
			},
		}
	} else {
		return &grequests.RequestOptions{ //, fmt.Errorf("%s is not a valid hash", hash)
		}
	}
}

func parseLookupHashOutput(lookupout []string) ResultsData {

	lookup := ResultsData{Found: false}

	if len(lookupout) > 0 {
		fields := strings.Fields(lookupout[0])

		if len(fields) > 0 {
			lookup.Detection = fields[0]
			s, _ := strconv.ParseInt(fields[1], 10, 64)
			lookup.LastSeen = string(time.Unix(s, 0).Format("20060102"))
		} else {
			log.Fatal(fmt.Errorf("Unable to parse LookupHashOutput: %#v\n", lookupout))
		}
	}

	return lookup
}

// LookupHash retreieves the TeamCymru - Malware Hash Registry for the given hash
func LookupHash(hash string) ResultsData {

	answers, _ := net.LookupTXT(hash + ".malware.hash.cymru.com")
	// fmt.Printf("answers: %#v\n", answers)
	// fmt.Printf("err: %#v\n", err)

	tcResult := parseLookupHashOutput(answers)
	// fmt.Printf("%#v", tcResult)
	return tcResult
}

func printStatus(resp gorequest.Response, body string, errs []error) {
	fmt.Println(resp.Status)
}

func printMarkDownTable(ss TeamCymru) {
	fmt.Println("#### TeamCymru")
	if ss.Results.Found {
		table := clitable.New([]string{"Found", "Detection", "LastSeen"})
		table.AddRow(map[string]interface{}{
			"Found":     ss.Results.Found,
			"Detection": ss.Results.Detection,
			"LastSeen":  ss.Results.LastSeen,
		})
		table.Markdown = true
		table.Print()
	} else {
		fmt.Println(" - Not found")
	}
}

var appHelpTemplate = `Usage: {{.Name}} {{if .Flags}}[OPTIONS] {{end}}COMMAND [arg...]

{{.Usage}}

Version: {{.Version}}{{if or .Author .Email}}

Author:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}
{{if .Flags}}
Options:
  {{range .Flags}}{{.}}
  {{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`

func main() {
	cli.AppHelpTemplate = appHelpTemplate
	app := cli.NewApp()
	app.Name = "team-cymru"
	app.Author = "blacktop"
	app.Email = "https://github.com/blacktop"
	app.Version = Version + ", BuildTime: " + BuildTime
	app.Compiled, _ = time.Parse("20060102", BuildTime)
	app.Usage = "Malice TeamCymru - Malware Hash Registry Plugin"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "post, p",
			Usage:  "POST results to Malice webhook",
			EnvVar: "MALICE_ENDPOINT",
		},
		cli.BoolFlag{
			Name:   "proxy, x",
			Usage:  "proxy settings for Malice webhook endpoint",
			EnvVar: "MALICE_PROXY",
		},
		cli.BoolFlag{
			Name:  "table, t",
			Usage: "output as Markdown table",
		},
	}
	app.ArgsUsage = "MD5/SHA1 hash of file"
	app.Action = func(c *cli.Context) {
		if c.Args().Present() {
			ssReport := LookupHash(c.Args().First())
			ss := TeamCymru{Results: ssReport}
			if c.Bool("table") {
				printMarkDownTable(ss)
			} else {
				ssJSON, err := json.Marshal(ss)
				assert(err)
				fmt.Println(string(ssJSON))
			}
		} else {
			log.Fatal(fmt.Errorf("Please supply a MD5/SHA1 hash to query."))
		}
	}

	err := app.Run(os.Args)
	assert(err)
}
