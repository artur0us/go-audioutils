package audioutils

import (
	"net/http"
	"os"
)

// isFileExists : ...
func (_a *AudioUtils) isFileExists(locationType int, fileLocation string) bool {
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
