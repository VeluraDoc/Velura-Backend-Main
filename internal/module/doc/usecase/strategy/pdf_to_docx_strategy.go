package doc_strategy

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

type PdfToDocx struct{}

func (file *PdfToDocx) Convert(input, outputFilePath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	scriptPath, _ := filepath.Abs("tools/pdf_convertor.py")
	cmd := exec.CommandContext(ctx, "python3", scriptPath, "-i", input, "-o", outputFilePath)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("convert failed: %w; stderr: %s", err, stderr.String())
	}
	return nil
}
