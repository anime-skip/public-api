package models

type BetterVRVEpisode struct {
	ID             string              `json:"objectId"`
	CreatedAt      string              `json:"createdAt"`
	UpdatedAt      string              `json:"updatedAt"`
	VRVID          string              `json:"episodeId"`
	EpisodeTitle   string              `json:"episodeTitle"`
	Season         *int                `json:"seasonNumber"`
	AmbiguosNumber *int                `json:"episodeNumber"`
	Series         BetterVRVSeriesLink `json:"series"`
	// When HasIntro=true, IntroStart and IntroEnd may exist, but they don't have to
	HasIntro   *bool    `json:"hasIntro"`
	IntroStart *float64 `json:"introStart"`
	IntroEnd   *float64 `json:"introEnd"`
	// When HasOutro=true, OutroStart and OutroEnd may exist, but they don't have to
	HasOutro   *bool    `json:"hasOutro"`
	OutroStart *float64 `json:"outroStart"`
	OutroEnd   *float64 `json:"outroEnd"`
	// When HasPostCredit=true, PostCreditStart and PostCreditEnd may exist, but they don't have to
	HasPostCredit   *bool    `json:"hasPostScene"`
	PostCreditStart *float64 `json:"postSceneStart"`
	PostCreditEnd   *float64 `json:"postSceneEnd"`
	// When HasPreview=true, PreviewStart and PreviewEnd may exist, but they don't have to
	HasPreview   *bool    `json:"hasPreview"`
	PreviewStart *float64 `json:"previewStart"`
	PreviewEnd   *float64 `json:"previewEnd"`
}

type BetterVRVSeriesLink struct {
	Type      string `json:"__type"`
	ClassName string `json:"className"`
	ObjectId  string `json:"objectId"`
}
