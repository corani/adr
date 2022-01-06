package adr

import (
	"errors"
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

var ErrForEachStop = errors.New("stop iterating")

func ForEach(conf *config.Config, callback func(*Adr) error) error {
	list, err := List(conf)
	if err != nil {
		return err
	}

	for _, adr := range list {
		if err := callback(adr); err != nil {
			if errors.Is(err, ErrForEachStop) {
				break
			}

			return err
		}
	}

	return nil
}

func ByID(conf *config.Config, number Number) (*Adr, error) {
	list, err := List(conf)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		if v.Number == number {
			return v, nil
		}
	}

	return nil, os.ErrNotExist
}

func Parse(path string) (*Adr, error) {
	adrFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w: parse: %v", ErrAdr, err)
	}
	defer adrFile.Close()

	var adr Adr

	body, err := frontmatter.Parse(adrFile, &adr)
	if err != nil {
		return nil, fmt.Errorf("%w: parse: %v", ErrAdr, err)
	}

	adr.Filename = filepath.Base(path)
	adr.Body = body

	return &adr, nil
}

func List(conf *config.Config) (Adrs, error) {
	root := filepath.Join(conf.Project, conf.Root)

	files, err := os.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("%w: list: %v", ErrAdr, err)
	}

	list := Adrs{}

	check := regexp.MustCompile(`[0-9]+\-.*?\.md`)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if ok := check.MatchString(file.Name()); ok {
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
	var id Number

	err := ForEach(conf, func(v *Adr) error {
		if v.Number > id {
			id = v.Number
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	id++

	adr := Adr{
		Filename: fmt.Sprintf("%04d-%s.md", id, slug.Make(title)),
		Number:   id,
		Title:    title,
		Status:   StatusProposed,
		Date:     time.Now().Format("2006-01-02"),
		Link:     0,
		Type:     "",
		Body:     nil,
	}

	log.Printf("creating ADR: %v", filepath.Join(conf.Root, adr.Filename))

	tmpl, err := template.ParseFiles(filepath.Join(conf.Project, conf.AdrTemplate))
	if err != nil {
		return nil, fmt.Errorf("%w: create: %v", ErrAdr, err)
	}

	out, err := os.Create(filepath.Join(conf.Project, conf.Root, adr.Filename))
	if err != nil {
		return nil, fmt.Errorf("%w: create: %v", ErrAdr, err)
	}
	defer out.Close()

	if err := tmpl.Execute(out, adr); err != nil {
		return nil, fmt.Errorf("%w: create: %v", ErrAdr, err)
	}

	return &adr, nil
}

func Index(conf *config.Config) error {
	log.Printf("updating index: %v", filepath.Join(conf.Root, "README.md"))

	list, err := List(conf)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(filepath.Join(conf.Project, conf.IndexTemplate))
	if err != nil {
		return fmt.Errorf("%w: index: %v", ErrAdr, err)
	}

	out, err := os.Create(filepath.Join(conf.Project, conf.Root, "README.md"))
	if err != nil {
		return fmt.Errorf("%w: index: %v", ErrAdr, err)
	}
	defer out.Close()

	if err := tmpl.Execute(out, list); err != nil {
		return fmt.Errorf("%w: index: %v", ErrAdr, err)
	}

	return nil
}

func Update(conf *config.Config, adr *Adr) error {
	log.Printf("updating ADR: %v", filepath.Join(conf.Root, adr.Filename))

	front, err := yaml.Marshal(adr)
	if err != nil {
		return fmt.Errorf("%w: update: %v", ErrAdr, err)
	}

	out, err := os.Create(filepath.Join(conf.Project, conf.Root, adr.Filename))
	if err != nil {
		return fmt.Errorf("%w: update: %v", ErrAdr, err)
	}
	defer out.Close()

	//nolint:errcheck
	{
		out.WriteString("---\n")
		out.Write(front)
		out.WriteString("---\n")
		out.Write(adr.Body)
	}

	return nil
}
