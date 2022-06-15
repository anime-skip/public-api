package utils_test

import (
	"testing"

	"anime-skip.com/public-api/internal/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "String Utils")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var _ = Describe("SanitizeURL", func() {
	DescribeTable("Valid URLs",
		func(input string, options utils.SanitizeURLOptions, expected string) {
			actual, err := utils.SanitizeURL(input, options)
			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		},
		Entry(
			"Valid URLs are returned verbatim by default",
			"https://google.com/some/path",
			utils.SanitizeURLOptions{},
			"https://google.com/some/path",
		),
		Entry(
			"Allowed schemes and hostnames work",
			"http://localhost:3000",
			utils.SanitizeURLOptions{
				AllowedSchemes:   []string{"https", "http"},
				AllowedHostnames: []string{"localhost"},
			},
			"http://localhost",
		),
		Entry(
			"Fragments are removed by default",
			"https://google.com#example",
			utils.SanitizeURLOptions{},
			"https://google.com",
		),
		Entry(
			"Fragments can be kept",
			"https://google.com#example",
			utils.SanitizeURLOptions{
				KeepFragment: true,
			},
			"https://google.com#example",
		),
		Entry(
			"Final unnecessary slashes are removed",
			"http://google.com/",
			utils.SanitizeURLOptions{
				KeepFragment: true,
			},
			"http://google.com/",
		),
		Entry(
			"Query params are removed by default",
			"https://google.com?query=param&test=value",
			utils.SanitizeURLOptions{},
			"https://google.com",
		),
		Entry(
			"Specified query params are kept, others are removed",
			"https://google.com?query=param&test=value",
			utils.SanitizeURLOptions{
				KeepQueryParams: []string{"test"},
			},
			"https://google.com?test=value",
		),
	)

	DescribeTable("Invalid URLs",
		func(input string, options utils.SanitizeURLOptions, expectedError string) {
			actual, err := utils.SanitizeURL(input, options)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(expectedError))
			Expect(actual).To(Equal(""))
		},
		Entry(
			"Passing a non-allowed hostname throws an error",
			"http://google.com",
			utils.SanitizeURLOptions{
				AllowedHostnames: []string{"anime-skip.com", "duckduckgo.com"},
			},
			"SanitizeURL: URL does not have the required hostname (allowed: anime-skip.com, duckduckgo.com | url: http://google.com | hostname: google.com)",
		),
		Entry(
			"Passing a non-allowed scheme throws an error",
			"http://google.com",
			utils.SanitizeURLOptions{
				AllowedSchemes: []string{"file", "https"},
			},
			"SanitizeURL: URL does not have the required scheme (allowed: file, https | url: http://google.com | scheme: http)",
		),
	)
})
