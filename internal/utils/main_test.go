package utils

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Utils")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var _ = Describe("CleanURL", func() {
	expected := "https://www.funimation.com/shows/the-rising-of-the-shield-hero/the-slave-girl/uncut"

	When("the input is already clean", func() {
		It("returns the same value", func() {
			actual := CleanURL("https://www.funimation.com/shows/the-rising-of-the-shield-hero/the-slave-girl/uncut")
			Expect(actual).To(Equal(expected))
		})
	})

	When("the input has a / at the end", func() {
		It("returns the url without the /", func() {
			actual := CleanURL("https://www.funimation.com/shows/the-rising-of-the-shield-hero/the-slave-girl/uncut/")
			Expect(actual).To(BeEquivalentTo(expected))
		})
	})

	When("the input has query params", func() {
		It("returns the url without the query params", func() {
			actual := CleanURL("https://www.funimation.com/shows/the-rising-of-the-shield-hero/the-slave-girl/uncut?lang=&a=1&qid=")
			Expect(actual).To(BeEquivalentTo(expected))
		})
	})

	When("the input has both a / at the end and query params", func() {
		It("returns the url without the final / and query params", func() {
			actual := CleanURL("https://www.funimation.com/shows/the-rising-of-the-shield-hero/the-slave-girl/uncut/?lang=&a=1&qid=")
			Expect(actual).To(BeEquivalentTo(expected))
		})
	})
})
