package helpers

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// HandleFileUpload extracts a file from the request, validates it, and saves it to the specified directory.
// It returns the generated filename, the full file path, and an error if any occurs.
func HandleFileUpload(c *fiber.Ctx, formField string, uploadDir string) (string, string, error) {
	// 1. Get file from the request
	file, err := c.FormFile(formField)
	if err != nil {
		return "", "", fmt.Errorf("failed to get file from request: %w", err)
	}

	// 2. Validate file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	if !allowedExtensions[ext] {
		return "", "", errors.New("invalid file format: only images (.jpg, .jpeg, .png, .gif, .webp) are allowed")
	}

	// 3. Generate a unique filename using timestamp and UUID
	uniqueID := uuid.New().String()[:8]
	newFilename := fmt.Sprintf("%d-%s%s", time.Now().Unix(), uniqueID, ext)

	// 4. Construct the full path where the file will be saved
	savePath := filepath.Join(uploadDir, newFilename)

	// 5. Save the file to the disk
	if err := c.SaveFile(file, savePath); err != nil {
		return "", "", fmt.Errorf("failed to save the file to disk: %w", err)
	}

	// 6. Return the raw filename and the path (relative or absolute depending on uploadDir)
	return newFilename, savePath, nil
}
