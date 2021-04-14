package audioutils

// FFProbeProbeData : ...
type FFProbeProbeData struct {
	Streams  []*FFProbeStream `json:"streams"`
	Chapters []interface{}    `json:"chapters"`
	Format   *FFProbeFormat   `json:"format"`
	Error    *FFProbeError    `json:"error"`
}

// FFProbeError : ...
type FFProbeError struct {
	Code   int    `json:"code"`
	String string `json:"string"`
}

// FFProbeFormat : ...
type FFProbeFormat struct {
	Filename         string             `json:"filename"`
	NBStreams        int                `json:"nb_streams"`
	NBPrograms       int                `json:"nb_programs"`
	FormatName       string             `json:"format_name"`
	FormatLongName   string             `json:"format_long_name"`
	StartTimeSeconds float64            `json:"start_time,string"`
	DurationSeconds  float64            `json:"duration,string"`
	Size             string             `json:"size"`
	BitRate          string             `json:"bit_rate"`
	ProbeScore       int                `json:"probe_score"`
	Tags             *FFProbeFormatTags `json:"tags"`
}

// FFProbeFormatTags : ...
type FFProbeFormatTags struct {
	MajorBrand       string `json:"major_brand"`
	MinorVersion     string `json:"minor_version"`
	CompatibleBrands string `json:"compatible_brands"`
	CreationTime     string `json:"creation_time"`
}

// FFProbeStream : ...
type FFProbeStream struct {
	Index              int                      `json:"index"`
	ID                 string                   `json:"id"`
	CodecName          string                   `json:"codec_name"`
	CodecLongName      string                   `json:"codec_long_name"`
	CodecType          string                   `json:"codec_type"`
	CodecTimeBase      string                   `json:"codec_time_base"`
	CodecTagString     string                   `json:"codec_tag_string"`
	CodecTag           string                   `json:"codec_tag"`
	RFrameRate         string                   `json:"r_frame_rate"`
	AvgFrameRate       string                   `json:"avg_frame_rate"`
	TimeBase           string                   `json:"time_base"`
	StartPts           int                      `json:"start_pts"`
	StartTime          string                   `json:"start_time"`
	DurationTs         uint64                   `json:"duration_ts"`
	Duration           string                   `json:"duration"`
	BitRate            string                   `json:"bit_rate"`
	BitsPerRawSample   string                   `json:"bits_per_raw_sample"`
	NbFrames           string                   `json:"nb_frames"`
	Disposition        FFProbeStreamDisposition `json:"disposition,omitempty"`
	Tags               FFProbeStreamTags        `json:"tags,omitempty"`
	Profile            string                   `json:"profile,omitempty"`
	Width              int                      `json:"width"`
	Height             int                      `json:"height"`
	HasBFrames         int                      `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string                   `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string                   `json:"display_aspect_ratio,omitempty"`
	PixFmt             string                   `json:"pix_fmt,omitempty"`
	Level              int                      `json:"level,omitempty"`
	ColorRange         string                   `json:"color_range,omitempty"`
	ColorSpace         string                   `json:"color_space,omitempty"`
	SampleFmt          string                   `json:"sample_fmt,omitempty"`
	SampleRate         string                   `json:"sample_rate,omitempty"`
	Channels           int                      `json:"channels,omitempty"`
	ChannelLayout      string                   `json:"channel_layout,omitempty"`
	BitsPerSample      int                      `json:"bits_per_sample,omitempty"`
}

// FFProbeStreamDisposition : ...
type FFProbeStreamDisposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
}

// FFProbeStreamTags : ...
type FFProbeStreamTags struct {
	Rotate       int    `json:"rotate,string,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
	Language     string `json:"language,omitempty"`
	Title        string `json:"title,omitempty"`
	Encoder      string `json:"encoder,omitempty"`
	Location     string `json:"location,omitempty"`
}
