package main

import (
	"context"
	"flag"
	"log"
	"os"
	"sitemap"
)

var (
	maxDepth   = flag.Uint("max-depth", 0, "max depth of url navigation recursion")
	outputFile = flag.String("output-file", "", "output file path")
	parallel   = flag.Uint("parallel", 1, "number of parallel workers to navigate through site")
	url        = flag.String("url", "", "start url")
)

func init() {
	flag.Parse()
}

func main() {
	switch "" {
	case *outputFile, *url:
		flag.PrintDefaults()
		os.Exit(1)
	}

	errorHandler := func(err error) {
		log.Println(err)
	}

	file, err := os.OpenFile(*outputFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		errorHandler(err)
		os.Exit(1)
	}
	defer file.Close()

	transport := sitemap.NewTransport(
		context.Background(),
		sitemap.Parse,
		errorHandler,
		*maxDepth,
		*parallel,
	)
	defer transport.Close()

	urls := transport.Fetch(*url)

	sitemap.Build(urls, file)
}
