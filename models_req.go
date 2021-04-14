package audioutils

// SrcAudioToHLSRequest : ...
type SrcAudioToHLSRequest struct {
	SrcLocationType int    `json:"src_location_type"`
	SrcLocation     string `json:"src_location"`
}

// AudioFileDurationRequest : ...
type AudioFileDurationRequest struct {
	SrcLocationType int    `json:"src_location_type"`
	SrcLocation     string `json:"src_location"`
}
