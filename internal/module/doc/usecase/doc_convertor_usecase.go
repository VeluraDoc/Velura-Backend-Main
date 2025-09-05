package doc_usecase

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type Strategy interface {
	Convert(inputFile, outputFilePath string) error
}

var cpuThreads int = (runtime.NumCPU() + 1) / 2

func Convertor(selectedStrategy Strategy, inputFile string) error {
	if strings.TrimSpace(inputFile) == "" {
		return errors.New("empty input list")
	}

	inputFiles := strings.Split(inputFile, ",")

	ch := make(chan struct{}, cpuThreads)

	var wg sync.WaitGroup

	var mu sync.Mutex

	var errs int

	for _, f := range inputFiles {
		file := strings.TrimSpace(f)
		if file == "" {
			continue
		}

		if !isFileExist(file) {
			fmt.Fprintf(os.Stderr, "file does not exist: %s\n", file)
			continue
		}

		inputFilePath, err := filepath.Abs(file)

		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get absolute path: %s\n", file)
			continue
		}

		ch <- struct{}{}
		wg.Add(1)

		go func(file string) {
			defer wg.Done()
			defer func() { <-ch }()

			if err := selectedStrategy.Convert(inputFilePath, outputFilePath(inputFilePath)); err != nil {
				mu.Lock()
				errs++
				mu.Unlock()
				fmt.Fprintf(os.Stderr, "failed to convert %s: %v\n", inputFilePath, err)
			}
		}(file)
	}

	wg.Wait()

	fmt.Printf("failed to convert %d files\n", errs)

	return nil
}

func isFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func outputFilePath(file string) string {
	dir := filepath.Dir(file)
	base := filepath.Base(file)
	nameNoExt := strings.TrimSuffix(base, filepath.Ext(base))
	return filepath.Join(dir, nameNoExt+".docx")
}
