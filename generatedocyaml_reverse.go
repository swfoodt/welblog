package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"github.com/adrg/frontmatter"
)

// DocMeta represents the docmeta field in frontmatter
type DocMeta struct {
	ID    string `yaml:"id"`
	Path  string `yaml:"path"`
	Title string `yaml:"title"`
	Weight int    `yaml:"weight"`
}

// Article represents a collected article with full info
type Article struct {
	Title  string `yaml:"title"`
	URL    string `yaml:"url"`
	Weight int    `yaml:"weight"`
	Path   string `yaml:"path"`
	ID     string `yaml:"id"`
	Source string `yaml:"source"`
}

// DocYAML is the structure written to data/docs/{id}.yaml
type DocYAML struct {
	ID     string                         `yaml:"id"`
	Routes map[string][]Article          `yaml:"routes"`
}

func main() {
	dataDir := "./data/docs"    // where YAML files are stored
	writeBack := true            // set to false if only collecting

	yamlFiles, err := filepath.Glob(filepath.Join(dataDir, "*.yaml"))
	if err != nil {
		log.Fatalf("Failed to read yaml files: %v", err)
	}

	for _, yamlPath := range yamlFiles {
		fmt.Println("Parsing:", yamlPath)
		data, err := ioutil.ReadFile(yamlPath)
		if err != nil {
			log.Printf("Failed to read yaml: %v", err)
			continue
		}

		var doc DocYAML
		if err := yaml.Unmarshal(data, &doc); err != nil {
			log.Printf("Failed to parse yaml: %v", err)
			continue
		}

		for _, articles := range doc.Routes {
			for _, a := range articles {
				fmt.Printf("Updating frontmatter for %s\n", a.Source)

				srcData, err := ioutil.ReadFile(a.Source)
				if err != nil {
					log.Printf("Failed to read %s: %v", a.Source, err)
					continue
				}

				var fm map[string]interface{}
				body, err := frontmatter.Parse(bytes.NewReader(srcData), &fm)
				if err != nil {
					log.Printf("Failed to parse frontmatter for %s: %v", a.Source, err)
					continue
				}

				fm["docmeta"] = map[string]interface{}{
					"id":     a.ID,
					"path":   a.Path,
					"title":  a.Title,
					"weight": a.Weight,
				}

				front, err := yaml.Marshal(fm)
				if err != nil {
					log.Printf("Failed to re-marshal frontmatter for %s: %v", a.Source, err)
					continue
				}

				var buf bytes.Buffer
				buf.WriteString("---\n")
				buf.Write(front)
				buf.WriteString("---\n\n")
				buf.Write(body)

				if writeBack {
					err = ioutil.WriteFile(a.Source, buf.Bytes(), 0644)
					if err != nil {
						log.Printf("Failed to write updated file %s: %v", a.Source, err)
					} else {
						fmt.Println("✅ Updated:", a.Source)
					}
				}
			}
		}
	}
}