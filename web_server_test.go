package main_test

import (
	// . "github.com/APShirley/web-server"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("WebServer", func() {
	It("Responds to requests", func() {
		Expect(1).To(Equal(1))
	})
})
