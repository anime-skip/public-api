package validation

import (
	"fmt"
	"testing"

	"anime-skip.com/backend/internal/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Account Validation")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var _ = Describe("AccountUsername", func() {
	Describe("when the username is less than 3 characters long", func() {
		testCases := []string{
			utils.RandomString(1),
			utils.RandomString(1) + " ",
			" " + utils.RandomString(1),
			" " + utils.RandomString(1) + " ",
			utils.RandomString(2),
			utils.RandomString(2) + " ",
			" " + utils.RandomString(2),
			" " + utils.RandomString(2) + " ",
			utils.RandomString(2) + "\n",
		}
		for _, testCase := range testCases {
			It(fmt.Sprintf("'%s' should return an error", testCase), func() {
				actual := AccountUsername(testCase)
				Expect(actual).To(MatchError("Username must be at least 3 characters long"))
			})
		}
	})

	Describe("any username with more than 3 characters", func() {
		testCases := []string{
			utils.RandomString(3),
			utils.RandomString(4),
			utils.RandomString(5),
			utils.RandomString(6),
		}
		for _, testCase := range testCases {
			It(fmt.Sprintf("should be valid for '%s'", testCase), func() {
				actual := AccountUsername(testCase)
				Expect(actual).To(BeNil())
			})
		}
	})
})

var _ = Describe("AccountEmail", func() {
	DescribeTable("valid email addresses",
		func(email string) {
			err := AccountEmail(email)
			Expect(err).To(BeNil())
		},
		Entry("''", "a@test.com"),
		Entry("b@gmail.com", "b@gmail.com"),
		Entry("some.email@test.com", "some.email@test.com"),
		Entry("some-email@test.com", "some-email@test.com"),
	)

	DescribeTable("invalid email addresses",
		func(email string) {
			err := AccountEmail(email)
			Expect(err).To(MatchError("Email is not valid"))
		},
		Entry("'' is invalid", ""),
		Entry("'some-email' is invalid", "some-email"),
		Entry("'some-email@test' is invalid", "some-email@test"),
	)

	DescribeTable("valid email addresses",
		func(email string) {
			err := AccountEmail(email)
			Expect(err).To(BeNil())
		},
		Entry("'a@test.com' is valid", "a@test.com"),
		Entry("'b@gmail.com' is valid", "b@gmail.com"),
		Entry("'some.email@test.com' is valid", "some.email@test.com"),
		Entry("'some-email@test.com' is valid", "some-email@test.com"),
	)
})
