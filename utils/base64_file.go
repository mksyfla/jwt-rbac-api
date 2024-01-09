package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Base64ToFile(filetype string, image string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return "", err
	}

	uploadDir := fmt.Sprintf("public/%s", filetype)

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, 0755)
	}

	randomFilename := generateRandomFilename(filetype)
	filename := filepath.Join(uploadDir, randomFilename)

	err = os.WriteFile(filename, decodedData, 0644)
	if err != nil {
		return "", err
	}

	return filename, err
}

func generateRandomFilename(extension string) string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	uniqueID := strconv.Itoa(int(rand.Uint32()))

	return fmt.Sprintf("%d_%s.%s", timestamp, uniqueID, extension)
}
