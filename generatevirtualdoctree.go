package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Entry struct {
	Path   string `yaml:"path"`
	Weight int    `yaml:"weight"`
	URL    string `yaml:"url"`
	Title  string `yaml:"title"`
	ID     string `yaml:"id"`
	Source string `yaml:"source"`
}

type DocYaml struct {
	ID          string             `yaml:"id"`
	Title       string             `yaml:"title"`
	Description string             `yaml:"description"`
	Routes      map[string][]Entry `yaml:"routes"`
}

func main() {
	yamlRoot := "data/docs"
	targetRoot := "content/chinese/docs"

	entries, err := os.ReadDir(yamlRoot)
	if err != nil {
		panic(err)
	}

	for _, file := range entries {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}

		yamlPath := filepath.Join(yamlRoot, file.Name())
		fmt.Println("Processing:", yamlPath)
		data, err := ioutil.ReadFile(yamlPath)
		if err != nil {
			fmt.Println("failed to read:", yamlPath)
			continue
		}

		doc := DocYaml{}
		err = yaml.Unmarshal(data, &doc)
		if err != nil {
			fmt.Println("failed to unmarshal:", yamlPath)
			continue
		}

		// 记录已生成的路径，防止重复
		generated := make(map[string]bool)

		// 生成 routes 中的目录结构
		for path, entryList := range doc.Routes {
			if len(entryList) == 0 {
				continue
			}

			segments := strings.Split(strings.Trim(path, "/"), "/")
			for i := 0; i < len(segments); i++ {
				subPath := strings.Join(segments[:i+1], "/")
				if generated[subPath] {
					continue
				}
				generated[subPath] = true

				entriesAtPath, ok := doc.Routes[subPath]
				if !ok || len(entriesAtPath) == 0 {
					continue
				}

				entry := entriesAtPath[0]
				folderPath := filepath.Join(targetRoot, filepath.FromSlash(subPath))
				err := os.MkdirAll(folderPath, os.ModePerm)
				if err != nil {
					fmt.Println("failed to create dir:", folderPath)
					continue
				}

				indexPath := filepath.Join(folderPath, "_index.md")
				content := fmt.Sprintf(`---
title: "%s"
slug: "%s"
docmeta:
  id: "%s"
  path: "%s"
  title: "%s"
  weight: %d
---
`, entry.Title, subPath, entry.ID, subPath, entry.Title, entry.Weight)

				err = ioutil.WriteFile(indexPath, []byte(content), 0644)
				if err != nil {
					fmt.Println("failed to write:", indexPath)
				} else {
					fmt.Println("generated:", indexPath)
				}
			}
		}

		// 同时为根路径创建 _index.md
		rootPath := filepath.Join(targetRoot, doc.ID)
		err = os.MkdirAll(rootPath, os.ModePerm)
		if err != nil {
			fmt.Println("failed to create root dir:", rootPath)
			continue
		}
		rootIndex := filepath.Join(rootPath, "_index.md")
		rootContent := fmt.Sprintf(`---
title: "%s"
slug: "%s"
docmeta:
  id: "%s"
  path: "%s"
  title: "%s"
  weight: 1
---
		`, doc.ID, doc.ID, doc.ID, doc.ID, doc.ID)
		ioutil.WriteFile(rootIndex, []byte(rootContent), 0644)
		fmt.Println("generated root:", rootIndex)
	}
}
