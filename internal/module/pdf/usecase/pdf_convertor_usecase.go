package pdf_usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func PdfToDocx(inputFile string) error {
	if strings.TrimSpace(inputFile) == "" {
		return errors.New("empty input list")
	}

	inputFiles := strings.Split(inputFile, ",")

	var wg sync.WaitGroup
	wg.Add(len(inputFiles))

	var mu sync.Mutex

	var errs int

	for _, file := range inputFiles {
		go func(file string) {
			defer wg.Done()
			if err := converter(file); err != nil {
				mu.Lock()
				errs++
				mu.Unlock()
				fmt.Fprintf(os.Stderr, "failed to convert %s: %v\n", file, err)
			}
		}(file)
	}

	wg.Wait()

	fmt.Printf("failed to convert %d files\n", errs)

	return nil
}

func converter(input string) error {
	if _, err := os.Stat(input); err != nil {
		return errors.New("input file does not exists")
	}

	inputFilePath, err := filepath.Abs(input)

	if err != nil {
		return errors.New("input file does not exists")
	}

	dir := filepath.Dir(inputFilePath)
	base := filepath.Base(inputFilePath)
	nameNoExt := strings.TrimSuffix(base, filepath.Ext(base))
	outputFilePath := filepath.Join(dir, nameNoExt+".docx")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	scriptPath, _ := filepath.Abs("tools/pdf_convertor.py")
	cmd := exec.CommandContext(ctx, "python3", scriptPath, "-i", inputFilePath, "-o", outputFilePath)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("convert failed: %w; stderr: %s", err, stderr.String())
	}
	return nil
}
