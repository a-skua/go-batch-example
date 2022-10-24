package main

import (
	"flag"
	"github.com/a-skua/go-batch-example/driver/file"
	"github.com/a-skua/go-batch-example/driver/repository"
	"github.com/a-skua/go-batch-example/pkg/user/upload"
	"log"
	"net/url"
	"os"
)

var (
	filename = flag.String("file", "", "choose upload file")
	rawURL   = flag.String("url", "http://localhost:8080", "base url")
)

func init() {
	flag.Parse()

	if *filename == "" {
		log.Println("-file is empty")
		os.Exit(1)
	}
}

func main() {
	data, err := file.NewUserData(*filename)
	if err != nil {
		log.Fatal(err)
	}

	url, err := url.Parse(*rawURL)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.New(url)

	err = upload.Upload(data, repo)
	if err != nil {
		log.Fatal(err)
	}
}
