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

	var outputBuf bytes.Buffer
	var stdErr bytes.Buffer

	cmd := exec.CommandContext(ctx, FFProbePath, args...)
	cmd.SysProcAttr = nil
	cmd.Stdout = &outputBuf
	cmd.Stderr = &stdErr
	// cmd.Stdin = reader

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to start ffprobe process: %w", err)
	}

	answer := &FFProbeProbeData{}
	if err := json.Unmarshal(outputBuf.Bytes(), answer); err != nil {
		return answer, fmt.Errorf("failed to parse ffprobe answer: %w", err)
	}

	return answer, nil
}
