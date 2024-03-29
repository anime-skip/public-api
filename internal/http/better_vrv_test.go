package http

import (
	"testing"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/utils"
	. "github.com/onsi/ginkgo/v2"
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
func newStringPointer(value string) *string {
	return &value
}

func TestBetterVRVService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HTTP")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var _ = Describe("parseBetterVRVEpisode", func() {
	var input betterVRVEpisode
	BeforeEach(func() {
		input = betterVRVEpisode{
			ID:              "",
			VRVID:           "",
			CreatedAt:       "",
			UpdatedAt:       "",
			AmbiguousNumber: newIntPointer(123),
			EpisodeTitle:    "title",
			Season:          newIntPointer(21),
			HasIntro:        nil,
			IntroStart:      nil,
			IntroEnd:        nil,
			HasOutro:        nil,
			OutroStart:      nil,
			OutroEnd:        nil,
			HasPostCredit:   nil,
			PostCreditStart: nil,
			PostCreditEnd:   nil,
			HasPreview:      nil,
			PreviewStart:    nil,
			PreviewEnd:      nil,
		}
	})

	Context("No Data", func() {
		When("there are nothing is known about the episode", func() {
			It("returns nil", func() {
				input.HasIntro = newBoolPointer(false)
				input.HasOutro = newBoolPointer(false)
				input.HasPostCredit = newBoolPointer(false)
				input.HasPreview = newBoolPointer(false)

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(BeNil())
			})
		})
	})

	Context("Basic Cases", func() {
		When("there are is only an intro", func() {
			It("returns (zero, start, and end) when both exist", func() {
				input.HasIntro = newBoolPointer(true)
				input.IntroStart = newFloatPointer(10)
				input.IntroEnd = newFloatPointer(20)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     10,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_INTRO),
						},
						{
							At:     20,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns (zero, start, start+90) when only start exists since intros are generally 90s", func() {
				input.HasIntro = newBoolPointer(true)
				input.IntroStart = newFloatPointer(10)
				input.IntroEnd = nil

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     10,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_INTRO),
						},
						{
							At:     100,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns (zero, end-90, end) when only end exists since intros are generally 90s", func() {
				input.HasIntro = newBoolPointer(true)
				input.IntroStart = nil
				input.IntroEnd = newFloatPointer(120)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     30,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_INTRO),
						},
						{
							At:     120,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns (zero, end) when only end exists and is less than 90 so that no timestamp is less than 0", func() {
				input.HasIntro = newBoolPointer(true)
				input.IntroStart = nil
				input.IntroEnd = newFloatPointer(80)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_INTRO),
						},
						{
							At:     80,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})
		})

		When("there are is only an outro", func() {
			It("returns (zero, start, end) when both exist", func() {
				input.HasOutro = newBoolPointer(true)
				input.OutroStart = newFloatPointer(10)
				input.OutroEnd = newFloatPointer(20)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     10,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_CREDITS),
						},
						{
							At:     20,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns (zero, start, start+90) when only start exists since outros are generally 90s", func() {
				input.HasOutro = newBoolPointer(true)
				input.OutroStart = newFloatPointer(10)
				input.OutroEnd = nil

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     10,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_CREDITS),
						},
						{
							At:     100,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns (zero, end-90, end) when only end exists since outros are generally 90s", func() {
				input.HasOutro = newBoolPointer(true)
				input.OutroStart = nil
				input.OutroEnd = newFloatPointer(80)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     -10,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_CREDITS),
						},
						{
							At:     80,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})
		})

		When("there are is only a post credits scene", func() {
			It("returns (zero, start, end) when both exist", func() {
				input.HasPostCredit = newBoolPointer(true)
				input.PostCreditStart = newFloatPointer(10)
				input.PostCreditEnd = newFloatPointer(20)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     10,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_CANON),
						},
						{
							At:     20,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns (zero, start) when only start exists since post credits are variable length", func() {
				input.HasPostCredit = newBoolPointer(true)
				input.PostCreditStart = newFloatPointer(10)
				input.PostCreditEnd = nil

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     10,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_CANON),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns (zero, end) when only end exists since post credits are variable length", func() {
				input.HasPostCredit = newBoolPointer(true)
				input.PostCreditStart = nil
				input.PostCreditEnd = newFloatPointer(20)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     20,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})
		})

		When("there are is only a preview", func() {
			It("returns (zero, start, and end) when both exist", func() {
				input.HasPreview = newBoolPointer(true)
				input.PreviewStart = newFloatPointer(10)
				input.PreviewEnd = newFloatPointer(20)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     10,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_PREVIEW),
						},
						{
							At:     20,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns (zero, start) when only start exists since post credits are variable length", func() {
				input.HasPreview = newBoolPointer(true)
				input.PreviewStart = newFloatPointer(10)
				input.PreviewEnd = nil

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     10,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_PREVIEW),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns (zero, end) when only end exists since post credits are variable length", func() {
				input.HasPreview = newBoolPointer(true)
				input.PreviewStart = nil
				input.PreviewEnd = newFloatPointer(20)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     20,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})
		})
	})

	Context("Sections combined", func() {
		When("the preview is a recap (at=0)", func() {
			It("returns the preview as RECAP and the intro afterwards when they are directly after each other", func() {
				input.HasPreview = newBoolPointer(true)
				input.PreviewStart = newFloatPointer(0)
				input.PreviewEnd = newFloatPointer(30)
				input.HasIntro = newBoolPointer(true)
				input.IntroStart = newFloatPointer(29)
				input.IntroEnd = newFloatPointer(110)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_RECAP),
						},
						{
							At:     29,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_INTRO),
						},
						{
							At:     110,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns the preview as RECAP and the intro afterwards when the recap doesn't have an end", func() {
				input.HasPreview = newBoolPointer(true)
				input.PreviewStart = newFloatPointer(0)
				input.PreviewEnd = nil
				input.HasIntro = newBoolPointer(true)
				input.IntroStart = newFloatPointer(90)
				input.IntroEnd = newFloatPointer(180)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_RECAP),
						},
						{
							At:     90,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_INTRO),
						},
						{
							At:     180,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns the preview as RECAP and the intro afterwards when they are not directly after each other", func() {
				input.HasPreview = newBoolPointer(true)
				input.PreviewStart = newFloatPointer(0)
				input.PreviewEnd = newFloatPointer(30)
				input.HasIntro = newBoolPointer(true)
				input.IntroStart = newFloatPointer(90)
				input.IntroEnd = newFloatPointer(180)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_RECAP),
						},
						{
							At:     30,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     90,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_INTRO),
						},
						{
							At:     180,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns no unknowns when everything chains together", func() {
				input.HasIntro = newBoolPointer(true)
				input.IntroStart = newFloatPointer(0)
				input.IntroEnd = newFloatPointer(90)
				input.HasOutro = newBoolPointer(true)
				input.OutroStart = newFloatPointer(91)
				input.OutroEnd = newFloatPointer(180)
				input.HasPostCredit = newBoolPointer(true)
				input.PostCreditStart = newFloatPointer(179)
				input.PostCreditEnd = newFloatPointer(190)
				input.HasPreview = newBoolPointer(true)
				input.PreviewStart = newFloatPointer(190)
				input.PreviewEnd = newFloatPointer(200)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_INTRO),
						},
						{
							At:     90,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_CREDITS),
						},
						{
							At:     179,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_CANON),
						},
						{
							At:     190,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_PREVIEW),
						},
						{
							At:     200,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})

			It("returns only the post credit when the recap is at the same time", func() {
				input.HasPostCredit = newBoolPointer(true)
				input.PostCreditStart = newFloatPointer(1000)
				input.PostCreditEnd = newFloatPointer(1020)
				input.HasPreview = newBoolPointer(true)
				input.PreviewStart = newFloatPointer(999)
				input.PreviewEnd = newFloatPointer(1019)

				expected := &internal.ThirdPartyEpisode{
					Name:           &input.EpisodeTitle,
					Number:         nil,
					AbsoluteNumber: newStringPointer("123"),
					Season:         newStringPointer("21"),
					Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
					Timestamps: []*internal.ThirdPartyTimestamp{
						{
							At:     0,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
						{
							At:     1000,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_CANON),
						},
						{
							At:     1020,
							TypeID: utils.Ptr(internal.TIMESTAMP_ID_UNKNOWN),
						},
					},
				}

				actual := parseBetterVRVEpisode(input)

				Expect(actual).To(Equal(expected))
			})
		})
	})

})
