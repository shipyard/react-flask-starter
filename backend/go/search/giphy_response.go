package search

type GiphyResponse struct {
	Data       []Data     `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Meta       Meta       `json:"meta,omitempty"`
}
type Original struct {
	Height   string `json:"height,omitempty"`
	Width    string `json:"width,omitempty"`
	Size     string `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
	Mp4Size  string `json:"mp4_size,omitempty"`
	Mp4      string `json:"mp4,omitempty"`
	WebpSize string `json:"webp_size,omitempty"`
	Webp     string `json:"webp,omitempty"`
	Frames   string `json:"frames,omitempty"`
	Hash     string `json:"hash,omitempty"`
}
type Downsized struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type DownsizedLarge struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type DownsizedMedium struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type DownsizedSmall struct {
	Height  string `json:"height,omitempty"`
	Width   string `json:"width,omitempty"`
	Mp4Size string `json:"mp4_size,omitempty"`
	Mp4     string `json:"mp4,omitempty"`
}
type DownsizedStill struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type FixedHeight struct {
	Height   string `json:"height,omitempty"`
	Width    string `json:"width,omitempty"`
	Size     string `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
	Mp4Size  string `json:"mp4_size,omitempty"`
	Mp4      string `json:"mp4,omitempty"`
	WebpSize string `json:"webp_size,omitempty"`
	Webp     string `json:"webp,omitempty"`
}
type FixedHeightDownsampled struct {
	Height   string `json:"height,omitempty"`
	Width    string `json:"width,omitempty"`
	Size     string `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
	WebpSize string `json:"webp_size,omitempty"`
	Webp     string `json:"webp,omitempty"`
}
type FixedHeightSmall struct {
	Height   string `json:"height,omitempty"`
	Width    string `json:"width,omitempty"`
	Size     string `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
	Mp4Size  string `json:"mp4_size,omitempty"`
	Mp4      string `json:"mp4,omitempty"`
	WebpSize string `json:"webp_size,omitempty"`
	Webp     string `json:"webp,omitempty"`
}
type FixedHeightSmallStill struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type FixedHeightStill struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type FixedWidth struct {
	Height   string `json:"height,omitempty"`
	Width    string `json:"width,omitempty"`
	Size     string `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
	Mp4Size  string `json:"mp4_size,omitempty"`
	Mp4      string `json:"mp4,omitempty"`
	WebpSize string `json:"webp_size,omitempty"`
	Webp     string `json:"webp,omitempty"`
}
type FixedWidthDownsampled struct {
	Height   string `json:"height,omitempty"`
	Width    string `json:"width,omitempty"`
	Size     string `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
	WebpSize string `json:"webp_size,omitempty"`
	Webp     string `json:"webp,omitempty"`
}
type FixedWidthSmall struct {
	Height   string `json:"height,omitempty"`
	Width    string `json:"width,omitempty"`
	Size     string `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
	Mp4Size  string `json:"mp4_size,omitempty"`
	Mp4      string `json:"mp4,omitempty"`
	WebpSize string `json:"webp_size,omitempty"`
	Webp     string `json:"webp,omitempty"`
}
type FixedWidthSmallStill struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type FixedWidthStill struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type Looping struct {
	Mp4Size string `json:"mp4_size,omitempty"`
	Mp4     string `json:"mp4,omitempty"`
}
type OriginalStill struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type OriginalMp4 struct {
	Height  string `json:"height,omitempty"`
	Width   string `json:"width,omitempty"`
	Mp4Size string `json:"mp4_size,omitempty"`
	Mp4     string `json:"mp4,omitempty"`
}
type Preview struct {
	Height  string `json:"height,omitempty"`
	Width   string `json:"width,omitempty"`
	Mp4Size string `json:"mp4_size,omitempty"`
	Mp4     string `json:"mp4,omitempty"`
}
type PreviewGif struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type PreviewWebp struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type Four80WStill struct {
	Height string `json:"height,omitempty"`
	Width  string `json:"width,omitempty"`
	Size   string `json:"size,omitempty"`
	URL    string `json:"url,omitempty"`
}
type Images struct {
	Original               Original               `json:"original,omitempty"`
	Downsized              Downsized              `json:"downsized,omitempty"`
	DownsizedLarge         DownsizedLarge         `json:"downsized_large,omitempty"`
	DownsizedMedium        DownsizedMedium        `json:"downsized_medium,omitempty"`
	DownsizedSmall         DownsizedSmall         `json:"downsized_small,omitempty"`
	DownsizedStill         DownsizedStill         `json:"downsized_still,omitempty"`
	FixedHeight            FixedHeight            `json:"fixed_height,omitempty"`
	FixedHeightDownsampled FixedHeightDownsampled `json:"fixed_height_downsampled,omitempty"`
	FixedHeightSmall       FixedHeightSmall       `json:"fixed_height_small,omitempty"`
	FixedHeightSmallStill  FixedHeightSmallStill  `json:"fixed_height_small_still,omitempty"`
	FixedHeightStill       FixedHeightStill       `json:"fixed_height_still,omitempty"`
	FixedWidth             FixedWidth             `json:"fixed_width,omitempty"`
	FixedWidthDownsampled  FixedWidthDownsampled  `json:"fixed_width_downsampled,omitempty"`
	FixedWidthSmall        FixedWidthSmall        `json:"fixed_width_small,omitempty"`
	FixedWidthSmallStill   FixedWidthSmallStill   `json:"fixed_width_small_still,omitempty"`
	FixedWidthStill        FixedWidthStill        `json:"fixed_width_still,omitempty"`
	Looping                Looping                `json:"looping,omitempty"`
	OriginalStill          OriginalStill          `json:"original_still,omitempty"`
	OriginalMp4            OriginalMp4            `json:"original_mp4,omitempty"`
	Preview                Preview                `json:"preview,omitempty"`
	PreviewGif             PreviewGif             `json:"preview_gif,omitempty"`
	PreviewWebp            PreviewWebp            `json:"preview_webp,omitempty"`
	Four80WStill           Four80WStill           `json:"480w_still,omitempty"`
}
type Onload struct {
	URL string `json:"url,omitempty"`
}
type Onclick struct {
	URL string `json:"url,omitempty"`
}
type Onsent struct {
	URL string `json:"url,omitempty"`
}
type Analytics struct {
	Onload  Onload  `json:"onload,omitempty"`
	Onclick Onclick `json:"onclick,omitempty"`
	Onsent  Onsent  `json:"onsent,omitempty"`
}
type Data struct {
	Type                     string    `json:"type,omitempty"`
	ID                       string    `json:"id,omitempty"`
	URL                      string    `json:"url,omitempty"`
	Slug                     string    `json:"slug,omitempty"`
	BitlyGifURL              string    `json:"bitly_gif_url,omitempty"`
	BitlyURL                 string    `json:"bitly_url,omitempty"`
	EmbedURL                 string    `json:"embed_url,omitempty"`
	Username                 string    `json:"username,omitempty"`
	Source                   string    `json:"source,omitempty"`
	Title                    string    `json:"title,omitempty"`
	Rating                   string    `json:"rating,omitempty"`
	ContentURL               string    `json:"content_url,omitempty"`
	SourceTld                string    `json:"source_tld,omitempty"`
	SourcePostURL            string    `json:"source_post_url,omitempty"`
	IsSticker                int       `json:"is_sticker,omitempty"`
	ImportDatetime           string    `json:"import_datetime,omitempty"`
	TrendingDatetime         string    `json:"trending_datetime,omitempty"`
	Images                   Images    `json:"images,omitempty"`
	AnalyticsResponsePayload string    `json:"analytics_response_payload,omitempty"`
	Analytics                Analytics `json:"analytics,omitempty"`
}
type Pagination struct {
	TotalCount int `json:"total_count,omitempty"`
	Count      int `json:"count,omitempty"`
	Offset     int `json:"offset,omitempty"`
}
type Meta struct {
	Status     int    `json:"status,omitempty"`
	Msg        string `json:"msg,omitempty"`
	ResponseID string `json:"response_id,omitempty"`
}
