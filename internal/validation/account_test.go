package validation

import (
	"testing"

	"anime-skip.com/public-api/internal/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validation")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var _ = Describe("AccountUsername", func() {
	DescribeTable("when the username is less than 3 characters long",
		func(inputUsername string) {
			actual := AccountUsername(inputUsername)
			Expect(actual).To(MatchError("Username must be at least 3 characters long"))
		},
		Entry("should return an error for 'x'", utils.RandomString(1)),
		Entry("should return an error for 'x '", utils.RandomString(1)+" "),
		Entry("should return an error for ' x'", " "+utils.RandomString(1)),
		Entry("should return an error for ' x '", " "+utils.RandomString(1)+" "),
		Entry("should return an error for 'xx'", utils.RandomString(2)),
		Entry("should return an error for 'xx '", utils.RandomString(2)+" "),
		Entry("should return an error for ' xx'", " "+utils.RandomString(2)),
		Entry("should return an error for ' xx '", " "+utils.RandomString(2)+" "),
		Entry("should return an error for 'xx\\n'", utils.RandomString(2)+"\n"),
	)

	DescribeTable("any username with more than 3 characters",
		func(inputUsername string) {
			actual := AccountUsername(inputUsername)
			Expect(actual).To(BeNil())
		},
		Entry("should be valid for 'xxx'", utils.RandomString(3)),
		Entry("should be valid for 'xxxx'", utils.RandomString(4)),
		Entry("should be valid for 'xxxxx'", utils.RandomString(5)),
		Entry("should be valid for 'xxxxxx'", utils.RandomString(6)),
	)
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
