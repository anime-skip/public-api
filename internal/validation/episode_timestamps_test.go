package validation_test

import (
	"fmt"
	"testing"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/config"
	"anime-skip.com/public-api/internal/validation"
	"github.com/gofrs/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/lo"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
}

var canon = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_CANON))
var credits = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_CREDITS))
var filler = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_FILLER))
var intro = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_INTRO))
var mixedIntro = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_MIXED_INTRO))
var newIntro = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_NEW_INTRO))
var preview = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_PREVIEW))
var titleCard = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_TITLE_CARD))
var transition = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_TRANSITION))
var unknown = lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_UNKNOWN))

var _ = Describe("Episode Timestamps Validation", func() {
	It("should create an unknown timestamp at 0 when there's not a timestamp before or at 0", func() {
		timestamps, isValid, errs := validation.EpisodeTimestamps(
			internal.Episode{},
			[]internal.Timestamp{
				{TypeID: canon, At: 6},
			},
		)
		Expect(isValid).To(BeTrue())
		Expect(errs).To(BeNil())
		Expect(timestamps).To(Equal([]internal.Timestamp{
			{TypeID: unknown, At: 0, Source: internal.TimestampSourceAnimeSkip},
			{TypeID: canon, At: 6},
		}))
	})

	It("should not create an unknown timestamp at 0 there's already a timestamp at 0", func() {
		timestamps, isValid, errs := validation.EpisodeTimestamps(
			internal.Episode{},
			[]internal.Timestamp{
				{TypeID: canon, At: 0},
			},
		)
		Expect(isValid).To(BeTrue())
		Expect(errs).To(BeNil())
		Expect(timestamps).To(Equal([]internal.Timestamp{
			{TypeID: canon, At: 0},
		}))
	})

	It("should not create an unknown timestamp at 0 there's a timestamp before 0", func() {
		timestamps, isValid, errs := validation.EpisodeTimestamps(
			internal.Episode{},
			[]internal.Timestamp{
				{TypeID: canon, At: -2},
			},
		)
		Expect(isValid).To(BeTrue())
		Expect(errs).To(BeNil())
		Expect(timestamps).To(Equal([]internal.Timestamp{
			{TypeID: canon, At: -2},
		}))
	})

	It("should merge timestamps of the same type at the same time", func() {
		timestamps, isValid, errs := validation.EpisodeTimestamps(
			internal.Episode{},
			[]internal.Timestamp{
				{TypeID: canon, At: 0},
				{TypeID: canon, At: 0},
			},
		)
		Expect(isValid).To(BeTrue())
		Expect(errs).To(BeNil())
		Expect(timestamps).To(Equal([]internal.Timestamp{
			{TypeID: canon, At: 0},
		}))
	})

	It("should not allow two of the same timestamp types in a row", func() {
		timestamps, isValid, errs := validation.EpisodeTimestamps(
			internal.Episode{},
			[]internal.Timestamp{
				{TypeID: canon, At: 0},
				{TypeID: canon, At: 30},
			},
		)
		Expect(isValid).To(BeFalse())
		Expect(errs).To(Equal([]error{
			fmt.Errorf(
				"The same timestamp type should not be used twice in a row: %s",
				config.TIMESTAMP_ID_CANON,
			),
		}))
		Expect(timestamps).To(Equal([]internal.Timestamp{
			{TypeID: canon, At: 0},
			{TypeID: canon, At: 30},
		}))
	})

	It("should not allow timestamps of different types at the same time", func() {
		timestamps, isValid, errs := validation.EpisodeTimestamps(
			internal.Episode{},
			[]internal.Timestamp{
				{TypeID: canon, At: 0},
				{TypeID: filler, At: 0},
				{TypeID: unknown, At: 52},
				{TypeID: filler, At: 52},
			},
		)
		Expect(isValid).To(BeFalse())
		Expect(errs).To(Equal([]error{
			fmt.Errorf(
				"Timestamps cannot be at the same time: %s+%s@0.00, %s+%s@52.00",
				config.TIMESTAMP_ID_CANON, config.TIMESTAMP_ID_FILLER,
				config.TIMESTAMP_ID_UNKNOWN, config.TIMESTAMP_ID_FILLER,
			),
		}))
		Expect(timestamps).To(Equal([]internal.Timestamp{
			{TypeID: canon, At: 0},
			{TypeID: filler, At: 0},
			{TypeID: unknown, At: 52},
			{TypeID: filler, At: 52},
		}))
	})

	It("should not allow more than 1 intro", func() {
		timestamps, isValid, errs := validation.EpisodeTimestamps(
			internal.Episode{},
			[]internal.Timestamp{
				{TypeID: newIntro, At: 0},
				{TypeID: intro, At: 90},
			},
		)
		Expect(isValid).To(BeFalse())
		Expect(errs).To(Equal([]error{
			fmt.Errorf("At most, there should be 1 intro(s): 2 intro(s) found"),
		}))
		Expect(timestamps).To(Equal([]internal.Timestamp{
			{TypeID: newIntro, At: 0},
			{TypeID: intro, At: 90},
		}))
	})

	It("should count new and mixed intros as intros", func() {
		timestamps, isValid, errs := validation.EpisodeTimestamps(
			internal.Episode{},
			[]internal.Timestamp{
				{TypeID: newIntro, At: 0},
				{TypeID: mixedIntro, At: 90},
			},
		)
		Expect(isValid).To(BeFalse())
		Expect(errs).To(Equal([]error{
			fmt.Errorf("At most, there should be 1 intro(s): 2 intro(s) found"),
		}))
		Expect(timestamps).To(Equal([]internal.Timestamp{
			{TypeID: newIntro, At: 0},
			{TypeID: mixedIntro, At: 90},
		}))
	})

	Describe("Intros", func() {
		DescribeTable(
			"valid durations",
			func(duration float64) {
				timestamps, isValid, errs := validation.EpisodeTimestamps(
					internal.Episode{},
					[]internal.Timestamp{
						{TypeID: intro, At: 0},
						{TypeID: canon, At: duration},
					},
				)
				Expect(isValid).To(BeTrue())
				Expect(errs).To(BeNil())
				Expect(timestamps).To(Equal([]internal.Timestamp{
					{TypeID: intro, At: 0},
					{TypeID: canon, At: duration},
				}))
			},
			Entry("should accept 1m30s intros", 90.0),
			Entry("should accept 50s intros as the shortest", 50.0),
			Entry("should accept 2m40s intros as the longest", 160.0),
		)

		DescribeTable(
			"invalid durations",
			func(duration float64) {
				timestamps, isValid, errs := validation.EpisodeTimestamps(
					internal.Episode{},
					[]internal.Timestamp{
						{TypeID: intro, At: 0},
						{TypeID: canon, At: duration},
					},
				)
				Expect(isValid).To(BeFalse())
				Expect(errs).To(Equal([]error{
					fmt.Errorf(
						"If intros exists, they should be between %.2f-%.2fs long: %.2fs",
						validation.INTRO_DURATION_MIN, validation.INTRO_DURATION_MAX,
						duration,
					),
				}))
				Expect(timestamps).To(Equal([]internal.Timestamp{
					{TypeID: intro, At: 0},
					{TypeID: canon, At: duration},
				}))
			},
			Entry("should not allow 1s intros", 1.0),
			Entry("should not allow 49s intros", 49.0),
			Entry("should not allow 101s intros", 161.0),
		)

		It("should use the episode duration as the next timestamp if there isn't one", func() {
			timestamps, isValid, errs := validation.EpisodeTimestamps(
				internal.Episode{
					BaseDuration: lo.ToPtr(10.0 + 170.0),
				},
				[]internal.Timestamp{
					{TypeID: canon, At: 0},
					{TypeID: newIntro, At: 10},
				},
			)
			Expect(isValid).To(BeFalse())
			Expect(errs).To(Equal([]error{
				fmt.Errorf(
					"If intros exists, they should be between %.2f-%.2fs long: %.2fs",
					validation.INTRO_DURATION_MIN, validation.INTRO_DURATION_MAX,
					170.0,
				),
			}))
			Expect(timestamps).To(Equal([]internal.Timestamp{
				{TypeID: canon, At: 0},
				{TypeID: newIntro, At: 10},
			}))
		})

		It("should ignore the intro if there's not another timestamp and the episode duration is not known", func() {
			timestamps, isValid, errs := validation.EpisodeTimestamps(
				internal.Episode{},
				[]internal.Timestamp{
					{TypeID: canon, At: 0},
					{TypeID: newIntro, At: 10},
				},
			)
			Expect(isValid).To(BeTrue())
			Expect(errs).To(BeNil())
			Expect(timestamps).To(Equal([]internal.Timestamp{
				{TypeID: canon, At: 0},
				{TypeID: newIntro, At: 10},
			}))
		})
	})

	It("should accept a standard episode sturcture", func() {
		input := []internal.Timestamp{
			{TypeID: canon, At: 0.00},
			{TypeID: intro, At: 273.50},
			{TypeID: titleCard, At: 364.00},
			{TypeID: canon, At: 369.00},
			{TypeID: transition, At: 757.00},
			{TypeID: canon, At: 762.00},
			{TypeID: credits, At: 1211.00},
			{TypeID: canon, At: 1301.00},
			{TypeID: preview, At: 1417.00},
		}
		timestamps, isValid, errs := validation.EpisodeTimestamps(
			internal.Episode{},
			input,
		)
		Expect(isValid).To(BeTrue())
		Expect(errs).To(BeNil())
		Expect(timestamps).To(Equal(input))
	})
})
