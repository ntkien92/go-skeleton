package main

import (
	"api-tests/config"
	"api-tests/suites"
	"testing"
)

func TestSuites(t *testing.T) {
	config.Init()

	suites.TestArticleSuite(t)
}
