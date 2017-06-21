package main_test

import (
	. "github.com/APShirley/web-server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		StartServer()

		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("firefox"))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should handle requests", func() {
		Expect(page.Navigate("http://localhost:3000")).To(Succeed())
		Expect(page).To(HaveURL("http://localhost:3000"))
		// TODO Something like this: Expect(page).To(ContainSubstring("Hello world!"))
	})
})
