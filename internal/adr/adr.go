package adr

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/corani/adr/internal/config"
	"github.com/gosimple/slug"
	"gopkg.in/yaml.v2"
)

type Status string

const (
	StatusProposed   Status = "proposed"
	StatusAccepted   Status = "accepted"
	StatusDeprecated Status = "deprecated"
	StatusSuperseded Status = "superseded"
)

type Number int

type Adr struct {
	Filename string `yaml:"-"`
	Type     string `yaml:"type"`
	Number   Number `yaml:"number"`
	Title    string `yaml:"title"`
	Date     string `yaml:"date"`
	Status   Status `yaml:"status"`
	Link     Number `yaml:"link,omitempty"`
	Body     []byte `yaml:"-"`
}

type Adrs map[Number]*Adr

func Parse(path string) (*Adr, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var adr Adr

	body, err := frontmatter.Parse(f, &adr)
	if err != nil {
		return nil, err
	}

	adr.Filename = filepath.Base(path)
	adr.Body = body

	return &adr, nil
}

func Index(conf *config.Config) (Adrs, error) {
	root := filepath.Join(conf.Project, conf.Root)

	files, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	list := Adrs{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if ok, _ := regexp.MatchString(`[0-9]+\-.*?\.md`, file.Name()); ok {
			if v, err := Parse(filepath.Join(root, file.Name())); err == nil {
				list[v.Number] = v
			} else {
				log.Printf("failed to parse %v: %v", file.Name(), err)
			}
		}
	}

	return list, nil
}

func Create(conf *config.Config, title string) (*Adr, error) {
	list, err := Index(conf)
	if err != nil {
		return nil, err
	}

	var id Number

	for k := range list {
		if k > id {
			id = k
		}
	}

	id++

	adr := Adr{
		Filename: fmt.Sprintf("%04d-%s.md", id, slug.Make(title)),
		Number:   id,
		Title:    title,
		Status:   StatusProposed,
		Date:     time.Now().Format("2006-01-02"),
		Link:     0,
	}

	tmpl, err := template.ParseFiles(filepath.Join(conf.Project, conf.Template))
	if err != nil {
		return nil, err
	}

	f, err := os.Create(filepath.Join(conf.Project, conf.Root, adr.Filename))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tmpl.Execute(f, adr)

	return &adr, nil
}

func Update(conf *config.Config, adr *Adr) error {
	front, err := yaml.Marshal(adr)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(conf.Project, conf.Root, adr.Filename))
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("---\n")
	f.Write(front)
	f.WriteString("---\n")
	f.Write(adr.Body)

	return nil
}
