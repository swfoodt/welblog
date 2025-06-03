package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

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
	contentDir := "./content/chinese/blog" 
	outputDir := "./data/docs"

	docMap := make(map[string]*DocYAML)

	err := filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		fmt.Printf("Processing file: %s\n", path)

		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("Failed to read %s: %v", path, err)
			return nil
		}

		var fm struct {
			Title    string   `yaml:"title"`
			Slug     string   `yaml:"slug"`
			DocMeta  DocMeta  `yaml:"docmeta"`
		}

		_, err = frontmatter.Parse(bytes.NewReader(data), &fm)
		if err != nil {
			log.Printf("Failed to parse frontmatter for %s: %v", path, err)
			return nil
		}

		if fm.DocMeta.ID == "" || fm.DocMeta.Path == "" {
			fmt.Printf("Skipped: missing docmeta.id or docmeta.path in %s\n", path)
			return nil
		}

		relURL := "/blog/" + strings.TrimSuffix(filepath.Base(path), ".md") + "/"
		article := Article{
			Title:  fm.DocMeta.Title,
			URL:    relURL,
			Weight: fm.DocMeta.Weight,
			Path:   fm.DocMeta.Path,
			ID:     fm.DocMeta.ID,
			Source: path,
		}

		doc, ok := docMap[fm.DocMeta.ID]
		if !ok {
			doc = &DocYAML{
				ID:     fm.DocMeta.ID,
				Routes: make(map[string][]Article),
			}
			docMap[fm.DocMeta.ID] = doc
		}

		fmt.Printf("Collected: ID=%s, Path=%s, URL=%s\n", fm.DocMeta.ID, fm.DocMeta.Path, relURL)
		doc.Routes[fm.DocMeta.Path] = append(doc.Routes[fm.DocMeta.Path], article)
		return nil
	})

	if err != nil {
		log.Fatalf("walk error: %v", err)
	}

	// sort and write
	for id, doc := range docMap {
		for path := range doc.Routes {
			articles := doc.Routes[path]
			sort.SliceStable(articles, func(i, j int) bool {
				return articles[i].Weight < articles[j].Weight
			})
			doc.Routes[path] = articles
		}

		outPath := filepath.Join(outputDir, id+".yaml")
		outData, err := yaml.Marshal(doc)
		if err != nil {
			log.Printf("YAML marshal error for %s: %v", id, err)
			continue
		}
		if err := ioutil.WriteFile(outPath, outData, 0644); err != nil {
			log.Printf("Failed to write %s: %v", outPath, err)
		} else {
			fmt.Println("Generated:", outPath)
		}
	}
}