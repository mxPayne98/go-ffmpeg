package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func compressVideo(inputPath, outputPath string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	// Acquire semaphore
	sem <- struct{}{}
	defer func() { <-sem }()

	cmd := exec.Command(
		"ffmpeg",
		"-i", inputPath,
		"-vf", "scale=1280:-1",
		"-c:v", "libx264",
		"-preset", "veryslow",
		"-crf", "28",
		"-c:a", "aac",
		"-b:a", "128k",
		outputPath,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Failed to compress %s: %v\n", inputPath, err)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run compress.go [inputDir] [outputDir]")
		return
	}

	inputDir := os.Args[1]
	outputDir := os.Args[2]

	// Number of concurrent goroutines. Adjust as necessary.
	concurrency := 4
	sem := make(chan struct{}, concurrency)

	var wg sync.WaitGroup

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".mp4" {
			outputPath := filepath.Join(outputDir, info.Name())
			outputPath = outputPath[:len(outputPath)-len(".mp4")] + ".mp4"
			fmt.Printf("Compressing %s to %s...\n", path, outputPath)

			wg.Add(1)
			go compressVideo(path, outputPath, &wg, sem)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	}

	wg.Wait()
}
