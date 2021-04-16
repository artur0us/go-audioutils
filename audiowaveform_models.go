package audioutils

import (
	"errors"
	"strings"
)

// audiowaveformWaveformDataInput : ...
type audiowaveformWaveformDataInput struct {
	// Required
	SrcLocation string `json:"file_location"`

	// Optional
	PointsPerSecond *int `json:"points_per_second"`
	DataBitsCount   *int `json:"data_bits_count"`
}

// validate : ...
func (_a *audiowaveformWaveformDataInput) validate() error {
	if _a == nil {
		return errors.New("input is nil")
	}

	// [ Required fields ]
	// SrcLocation
	if strings.ReplaceAll(_a.SrcLocation, " ", "") == "" {
		return errors.New("source location is empty")
	}

	// [ Optional fields ]
	// PointsPerSecond
	if _a.PointsPerSecond != nil {
		if *_a.PointsPerSecond < 1 || *_a.PointsPerSecond > 100 {
			return errors.New("points per seconds value is invalid")
		}
	}
	// DataBitsCount
	if _a.DataBitsCount != nil {
		if *_a.DataBitsCount != 8 && *_a.DataBitsCount != 16 {
			return errors.New("data bits count is invalid")
		}
	}

	return nil
}

// ----------------------------------------- //

// audiowaveformJSON : ...
type audiowaveformJSON struct {
	Version         int   `json:"version"`
	Channels        int   `json:"channels"`
	SampleRate      int   `json:"sample_rate"`
	SamplesPerPixel int   `json:"samples_per_pixel"`
	Bits            int   `json:"bits"`
	Length          int   `json:"length"`
	Data            []int `json:"data"`
}
