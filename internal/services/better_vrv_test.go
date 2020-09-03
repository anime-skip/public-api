package services

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func newIntPointer(value int) *int {
	return &value
}
func newBoolPointer(value bool) *bool {
	return &value
}
func newFloatPointer(value float64) *float64 {
	return &value
}

func TestBetterVRVService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Better VRV Service")
}

var _ = Describe("MapBetterVRVEpisodeToThirdPartyEpisode", func() {
	When("there the episode is nil", func() {
		It("returns `nil`", func() {
		})
	})

	When("there are nothing is known about the episode", func() {
		It("returns `nil`", func() {
		})
	})

	When("there are is only an intro", func() {
		It("returns `nil`", func() {
		})
	})

	When("there are is only an outro", func() {
		It("returns `nil`", func() {
		})
	})

	When("there are is only a post credits scene", func() {
		It("returns `nil`", func() {
		})
	})

	When("there are is only a preview", func() {
		It("returns `nil`", func() {
		})
	})

	When("the preview is 0", func() {
		It("returns the preview as RECAP and the intro afterwards", func() {
		})
	})

	When("the preview is 0 and the intro is at it's end", func() {
		It("returns the preview as RECAP and the intro direcly afterwards", func() {

		})
	})
})

// func TestMapBetterVRVEpisodeToThirdPartyEpisode_NoTimestamps(t *testing.T) {
// 	input := &BetterVRVEpisode{
// 		ID:             "id",
// 		EpisodeTitle:   "Title",
// 		CreatedAt:      "create_at",
// 		UpdatedAt:      "updated_at",
// 		AmbiguosNumber: newIntPointer(1),
// 		Season:         newIntPointer(2),
// 		VRVID:          "advb7r39z",
// 		HasIntro:       newBoolPointer(false),
// 		HasOutro:       newBoolPointer(false),
// 		HasPostCredit:  newBoolPointer(false),
// 		HasPreview:     newBoolPointer(false),
// 	}
// 	var expected *models.ThirdPartyEpisode = nil

// 	actual := MapBetterVRVEpisodeToThirdPartyEpisode(input)

// }
