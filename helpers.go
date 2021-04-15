package audioutils

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// isFileExists : ...
func isFileExists(locationType int, fileLocation string) bool {
	if locationType == AudioFileSrcLocationTypeLocal {
		if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
			return false
		}
	} else if locationType == AudioFileSrcLocationTypeURL {
		resp, err := http.Head(fileLocation)
		if err != nil {
			return false
		}
		if resp.StatusCode != http.StatusOK {
			return false
		}
	}
	return true
}

// addStrBeforeFileNameInM3U8 : ...
func addStrBeforeFileNameInM3U8(m3u8FileAbsPath string, selectedLinesPatterns []string, appendingStr string) error {
	// selectedLinesPatterns = []string{"hls_", ".ts"}
	// selectedLinesPatterns = []string{"result_file_", ".m4a"}

	m3u8FileBytes, err := ioutil.ReadFile(m3u8FileAbsPath)
	if err != nil {
		return err
	}

	var m3u8FileStrLines []string = strings.Split(string(m3u8FileBytes), "\n")
	if len(m3u8FileStrLines) < 1 {
		return errors.New("specified M3U8 file is empty")
	}

	var processedM3U8Lines []string = []string{}

	for _, oneLine := range m3u8FileStrLines {
		var detectedPatternsCount int = 0
		for _, onePattern := range selectedLinesPatterns {
			if strings.Contains(strings.ToLower(oneLine), onePattern) {
				detectedPatternsCount++
			}
		}

		if len(selectedLinesPatterns) == detectedPatternsCount {
			processedM3U8Lines = append(processedM3U8Lines, appendingStr+oneLine)
			continue
		}

		processedM3U8Lines = append(processedM3U8Lines, oneLine)
	}

	// Truncate current file
	srcFile, err := os.OpenFile(m3u8FileAbsPath, os.O_TRUNC, 0755)
	if err != nil {
		log.Println(err)
		return err
	}
	if err = srcFile.Close(); err != nil {
		log.Println(err)
		return err
	}

	// Write new content to file
	destFile, err := os.OpenFile(m3u8FileAbsPath, os.O_CREATE|os.O_WRONLY, 0644)
	defer destFile.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	dataWriter := bufio.NewWriter(destFile)
	for _, oneStrLine := range processedM3U8Lines {
		_, _ = dataWriter.WriteString(oneStrLine + "\n")
	}
	dataWriter.Flush()

	return nil
}

// IsAudioFile : ...
// func (_a *AudioUtils) IsAudioFile(locationType int, fileLocation string) bool {
// 	if locationType == AudioFileSrcLocationTypeLocal {
// 		f, err := os.Open(fileLocation)
// 		if err != nil {
// 			return false
// 		}
// 		defer f.Close()

// 		buffer := make([]byte, 512)
// 		if _, err = f.Read(buffer); err != nil {
// 			return false
// 		}
// 		contentType := http.DetectContentType(buffer)
// 		log.Println(contentType)

// 		// TODO: ...

// 		return true
// 	} else if locationType == AudioFileSrcLocationTypeURL {
// 		// TODO: not implemented
// 		return false
// 	}

// 	return false
// }
