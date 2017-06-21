package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWebServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WebServer Suite")
}
