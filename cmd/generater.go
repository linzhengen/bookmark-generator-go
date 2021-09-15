package cmd

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/urfave/cli/v2"
)

//go:embed bookmarks.tpl
var bookMarkTpl string

type chromeBookMarks struct {
	Folders []struct {
		Name      string    `yaml:"name"`
		Bookmarks bookMarks `yaml:"bookmarks"`
	} `yaml:"folders"`
	Bookmarks bookMarks `yaml:"bookmarks"`
}

type bookMarks []bookMark

type bookMark struct {
	URL     string `yaml:"url"`
	Name    string `yaml:"name"`
	AddDate string `yaml:"addDate"`
	IconURL string `yaml:"iconUrl"`
}

func GenerateCmd() *cli.Command {
	return &cli.Command{
		Name:  "gen",
		Usage: "Generates chrome bookmarks from yaml",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "yaml file",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"0"},
				Usage:   "Write to file instead of stdout",
			},
		},
		Action: func(c *cli.Context) error {
			return generateAction(c)
		},
	}
}

func generateAction(c *cli.Context) error {
	bm := readYaml(c.String("file"))
	render(bm)
	return nil
}

func readYaml(path string) chromeBookMarks {
	if _, err := os.Stat(path); err != nil {
		log.Fatalf("not found file: %s", path)
	}
	fb, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	var bm chromeBookMarks
	if err := yaml.Unmarshal(fb, &bm); err != nil {
		log.Fatalf("parse yaml error: %v", err)
	}
	return bm
}

func render(bm chromeBookMarks) string {
	t, err := template.New("").Parse(bookMarkTpl)
	if err != nil {
		log.Fatalf("parse tpl error: %v", err)
	}
	var html bytes.Buffer
	if err := t.Execute(&html, bm); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(html.String())
	return html.String()
}
