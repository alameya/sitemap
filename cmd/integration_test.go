package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M)  {
	if testing.Short() {
		return
	}

	*maxDepth   = 1
	*outputFile = "tmp_sitemap.xml"
	*parallel   = 10
	*url        = "https://www.sitemaps.org/index.html"

	exitCode := m.Run()

	os.Remove(*outputFile)

	os.Exit(exitCode)
}

func TestFull(t *testing.T)  {
	if testing.Short() {
		return
	}

	main()

	checkFile(t)
}

func checkFile(t *testing.T)  {
	t.Helper()

	fileInfo, err := os.Stat(*outputFile)
	assert.NoError(t, err)
	assert.True(t, fileInfo.Size() > 100)
}