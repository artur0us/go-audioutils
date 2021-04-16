package audioutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// audiowaveform binary path
const (
	audiowaveformPath string = "audiowaveform"
)

// audiowaveformGetWaveformData : ...
func (_a *AudioUtils) audiowaveformGetWaveformData(input audiowaveformWaveformDataInput) (*audiowaveformJSON, error) {
	// 3 points per second:
	// audiowaveform -i source_file.mp3 --pixels-per-second 3 -b 8 -o source_file.json
	// 20 points per second:
	// audiowaveform -i source_file.mp3 --pixels-per-second 20 -b 8 -o source_file.json

	var err error = nil

	var downloadedTempFilePath string = fmt.Sprintf("./%v.mp3", genRandStr(32))
	var waveformDataTempJSONFilePath string = fmt.Sprintf("./%v.json", genRandStr(32))
	downloadedTempFilePath, err = filepath.Abs(downloadedTempFilePath)
	if err != nil {
		_a.ErrLogger.Printf("downloaded temp file absolute path parse failed: %v\n", err)
		return nil, fmt.Errorf("downloaded temp file absolute path parse failed: %v", err)
	}
	waveformDataTempJSONFilePath, err = filepath.Abs(waveformDataTempJSONFilePath)
	if err != nil {
		_a.ErrLogger.Printf("waveform data temp JSON file absolute path parse failed: %v\n", err)
		return nil, fmt.Errorf("waveform data temp JSON file absolute path parse failed: %v", err)
	}
	defer func() {
		os.RemoveAll(downloadedTempFilePath)
		os.RemoveAll(waveformDataTempJSONFilePath)
	}()

	_a.InfoLogger.Printf("validating input: %v\n", input)
	if err := input.validate(); err != nil {
		_a.ErrLogger.Printf("input validation failed: %v\n", err)
		return nil, err
	}

	var fileLocalPath string = input.SrcLocation

	// 1. Download file if source location is URL
	if isUrl(input.SrcLocation) {
		_a.InfoLogger.Printf("web file location detected: %v\n", input.SrcLocation)

		if !isFileExists(AudioFileSrcLocationTypeURL, input.SrcLocation) {
			_a.ErrLogger.Printf("specified web file is not found: %v\n", input.SrcLocation)
			return nil, fmt.Errorf("specified web file is not found: %v", input.SrcLocation)
		}

		_a.InfoLogger.Printf("downloading file: %v\n", input.SrcLocation)
		err := downloadFileFromWeb(input.SrcLocation, downloadedTempFilePath)
		if err != nil {
			_a.ErrLogger.Printf("failed to download file: %v\n", input.SrcLocation)
			return nil, fmt.Errorf("failed to download file: %v", input.SrcLocation)
		}

		fileLocalPath = downloadedTempFilePath
	} else {
		fileLocalPath, err = filepath.Abs(fileLocalPath)
		if err != nil {
			_a.ErrLogger.Printf("local file absolute path parse failed: %v\n", err)
			return nil, fmt.Errorf("local file absolute path parse failed: %v", err)
		}
	}

	pointsPerSecond := 3
	if input.PointsPerSecond != nil {
		pointsPerSecond = *input.PointsPerSecond
	}

	dataBitsCount := 8
	if input.DataBitsCount != nil {
		dataBitsCount = *input.DataBitsCount
	}

	_a.InfoLogger.Printf("audiowaveformGetWaveformData internal variables -> downloadedTempFilePath: %v\n", downloadedTempFilePath)
	_a.InfoLogger.Printf("audiowaveformGetWaveformData internal variables -> waveformDataTempJSONFilePath: %v\n", waveformDataTempJSONFilePath)
	_a.InfoLogger.Printf("audiowaveformGetWaveformData internal variables -> fileLocalPath: %v\n", fileLocalPath)
	_a.InfoLogger.Printf("audiowaveformGetWaveformData internal variables -> pointsPerSecond: %v\n", pointsPerSecond)
	_a.InfoLogger.Printf("audiowaveformGetWaveformData internal variables -> dataBitsCount: %v\n", dataBitsCount)

	// 2. Call "audiowaveform"
	_a.InfoLogger.Println("building new audiowaveform process config...")
	ctx := context.Background()
	args := []string{
		"-i", fileLocalPath,
		"--pixels-per-second", strconv.Itoa(pointsPerSecond),
		"-b", strconv.Itoa(dataBitsCount),
		"-o", waveformDataTempJSONFilePath,
	}

	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer

	cmd := exec.CommandContext(ctx, audiowaveformPath, args...)
	cmd.SysProcAttr = nil
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	_a.InfoLogger.Println("running audiowaveform process...")
	if err := cmd.Run(); err != nil {
		_a.ErrLogger.Printf("failed to start audiowaveform process: %v\n", err)
		_a.ErrLogger.Printf("audiowaveform stdout:\n%v\n", stdoutBuf.String())
		_a.ErrLogger.Printf("audiowaveform stderr:\n%v\n", stderrBuf.String())
		return nil, fmt.Errorf("failed to start audiowaveform process: %w", err)
	}

	// 3. Answer preparation
	_a.InfoLogger.Println("audiowaveform process done, reading file with result...")
	audiowaveformAnswer := &audiowaveformJSON{}
	waveformDataTempJSONFileBytes, err := ioutil.ReadFile(waveformDataTempJSONFilePath)
	if err != nil {
		_a.ErrLogger.Printf("failed to read waveform data JSON file: %v\n", err)
		return nil, fmt.Errorf("failed to read waveform data JSON file: %v", err)
	}
	_a.InfoLogger.Println("parsing result...")
	err = json.Unmarshal(waveformDataTempJSONFileBytes, &audiowaveformAnswer)
	if err != nil {
		_a.ErrLogger.Printf("failed to parse waveform data: %v\n", err)
		return nil, fmt.Errorf("failed to parse waveform data: %v", err)
	}

	return audiowaveformAnswer, nil
}
