# This is a test issue 

## Usage example
```bash
make build
chmod u+x ./sitemap 
./sitemap -url https://example.com -parallel 10 -output-file sitemap.xml max-depth 1
```

## Required parameters
`-url` - start url

`-output-file` - output file path

## Optional parameters
`-max-depth` - max depth of url navigation recursion. Default is 0 (include only current url)

`-parallel` - number of parallel workers to navigate through site. Default is 1

# Testing
### Unit tests
`make test`

### Unit race tests
`make test-race`

### Integration test
`make integration-test`

### Tests coverage
`make coverage`

### Lint
`make lint`

## Issue text
```
Sitemap generator

Implement simple sitemap (https://www.sitemaps.org) generator as command line tool.

It should:
	•	accept start url as argument
	•	recursively navigate by site pages in parallel 
	•	should not use any external dependencies, only standard golang library
	•	extract page urls only from <a> elements and take in account <base> element if declared
	•	should be well tested (automated testing)

Suggested program options:
-parallel=  			number of parallel workers to navigate through site
-output-file= 			output file path
-max-depth= 			max depth of url navigation recursion
```