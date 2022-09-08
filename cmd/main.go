package main

import (
	"csv2tsv/utils"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func csv2tsv(csvFile string) {
	base := path.Base(csvFile)
	suffix := path.Ext(csvFile)
	tsvFile := base[0:len(base)-len(suffix)] + ".tsv"
	command := fmt.Sprintf("sed -E 's/(\"([^\"]*)\")?,/\\2\\t/g' %s > %s", csvFile, tsvFile)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		log.Printf("%s convert to tsv error: %v", csvFile, err)
	}
}

func main() {
	dirPath := flag.String("path", "", "csv directory path")
	flag.Usage = func() {
		_, _ = fmt.Fprint(os.Stderr,
			`Usage: ./csv2tsv  OPTIONS [arg...]
       ./csv2tsv [ -path /tmp/csv | -help ]
Options:
  -path            csv directory path。
  -help            Print usage
`)
	}
	flag.Parse()

	if !utils.PathExists(*dirPath) || !utils.IsDir(*dirPath) {
		log.Fatal("the path does not exist")
	}

	var files []string
	err := filepath.Walk(*dirPath, utils.ListFiles(&files))
	if err != nil {
		log.Fatal("filepath walk failed:", err)
	}

	for _, file := range files {
		csv2tsv(file)
	}
	log.Printf("finished csv2tsv")
}
