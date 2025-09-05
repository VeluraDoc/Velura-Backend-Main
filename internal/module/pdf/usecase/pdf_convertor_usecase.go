package pdf_usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func PdfToDocx(input, output string) error {
	if input == "" || output == "" {
		return errors.New("empty input or output")
	}

	if _, err := os.Stat(input); err != nil {
		return errors.New("input file does not exists")
	}

	filePath, err := filepath.Abs(input)

	if err != nil {
		return errors.New("input file does not exists")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	scriptPath, _ := filepath.Abs("tools/pdf_convertor.py")
	cmd := exec.CommandContext(ctx, "python3", scriptPath, "-i", filePath, "-o", output+".docx")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("convert failed: %w; stderr: %s", err, stderr.String())
	}
	return nil
}
