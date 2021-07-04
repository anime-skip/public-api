package rate_limiter

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestBetterVRVService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Better VRV Service")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var client1 = entities.APIClient{
	ID:           "client1",
	RateLimitRPM: utils.UIntPtr(1),
}
var clientN = entities.APIClient{
	ID:           "client1",
	RateLimitRPM: utils.UIntPtr(uint(rand.Int())),
}
var clientInfinity = entities.APIClient{
	ID:           "client1",
	RateLimitRPM: nil,
}

var _ = Describe("Increment", func() {
	BeforeEach(func() {
		delete(rates, client1.ID)
		delete(rates, clientN.ID)
		delete(rates, clientInfinity.ID)
	})

	It("should rate limit based on the client's requests per minute and client id", func() {
		err1 := Increment(client1)
		err2 := Increment(client1)
		err3 := Increment(client1)
		Expect(err1).To(BeNil())
		Expect(err2).To(MatchError("Rate limit exceeded"))
		Expect(err3).To(MatchError("Rate limit exceeded"))
	})

	randomRates := []int{
		rand.Intn(100) + 1,
		rand.Intn(100) + 1,
		rand.Intn(100) + 1,
		rand.Intn(100) + 1,
	}
	DescribeTable("respect random rate limit",
		func(rate int) {
			clientN.RateLimitRPM = utils.UIntPtr(uint(rate))
			i := 0
			for i < rate {
				i++
				err := Increment(clientN)
				Expect(err).To(BeNil())
			}
			err := Increment(clientN)

			Expect(err).To(MatchError("Rate limit exceeded"))
		},
		Entry(fmt.Sprintf("%d RPM", randomRates[0]), randomRates[0]),
		Entry(fmt.Sprintf("%d RPM", randomRates[1]), randomRates[1]),
		Entry(fmt.Sprintf("%d RPM", randomRates[2]), randomRates[2]),
		Entry(fmt.Sprintf("%d RPM", randomRates[3]), randomRates[3]),
	)

	It("should not track clients with no rate limit", func() {
		err := Increment(clientInfinity)
		_, hasTracker := rates[clientInfinity.ID]

		Expect(err).To(BeNil())
		Expect(hasTracker).To(BeFalse())
	})

	It("should return an nil when the timer is expired and the client was exceeding the rate limit", func() {
		rates[client1.ID] = &rateTracker{
			count:   *client1.RateLimitRPM + 1,
			expires: time.Now().Add(time.Second),
		}
		err1 := Increment(client1)
		rates[client1.ID].expires = time.Now().Add(-time.Second)
		err2 := Increment(client1)

		Expect(err1).ToNot(BeNil())
		Expect(err2).To(BeNil())
	})
})
