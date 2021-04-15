package audioutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

// ffprobe binary path
const (
	FFProbePath string = "ffprobe"
)

// ffprobeGetAudioFileInfo : ...
// Returns: ffprobe answer model, error
func (_a *AudioUtils) ffprobeGetAudioFileInfo(fileLocation string) (*FFProbeProbeData, error) {
	// Local:
	// ffprobe -loglevel fatal -print_format json -show_format -show_streams -v quiet -show_error -show_chapters source_file.mp3
	// URL:
	// ffprobe -loglevel fatal -print_format json -show_format -show_streams -v quiet -show_error -show_chapters https://localhost/source_file.mp3

	_a.InfoLogger.Println("building new ffprobe process config...")
	ctx := context.Background()
	args := []string{
		"-loglevel", "fatal",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		"-v", "quiet",
		"-show_error",
		"-show_chapters",
	}
	args = append(args, fileLocation)

	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer

	cmd := exec.CommandContext(ctx, FFProbePath, args...)
	cmd.SysProcAttr = nil
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	// cmd.Stdin = reader

	_a.InfoLogger.Println("running ffprobe process...")
	if err := cmd.Run(); err != nil {
		_a.ErrLogger.Printf("failed to start ffprobe process: %v\n", err)
		_a.ErrLogger.Printf("ffprobe stdout:\n%v\n", stdoutBuf.String())
		_a.ErrLogger.Printf("ffprobe stderr:\n%v\n", stderrBuf.String())
		return nil, fmt.Errorf("failed to start ffprobe process: %w", err)
	}

	answer := &FFProbeProbeData{}
	if err := json.Unmarshal(stdoutBuf.Bytes(), answer); err != nil {
		_a.ErrLogger.Printf("failed to parse ffprobe answer: %v", err)
		return answer, fmt.Errorf("failed to parse ffprobe answer: %w", err)
	}

	return answer, nil
}
