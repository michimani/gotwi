package fields

type MediaField string

const (
	DurationMs            MediaField = "duration_ms"
	Height                MediaField = "height"
	MediaKey              MediaField = "media_key"
	PreviewImageURL       MediaField = "preview_image_url"
	Type                  MediaField = "type"
	MediaURL              MediaField = "url"
	Width                 MediaField = "width"
	MediaPublicMetrics    MediaField = "public_metrics"
	MediaNonPublicMetrics MediaField = "non_public_metrics"
	MediaOrganicMetrics   MediaField = "organic_metrics"
	MediaPromotedMetrics  MediaField = "promoted_metrics"
	AltText               MediaField = "alt_text"
)
