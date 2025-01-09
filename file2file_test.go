package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFile2File(t *testing.T) {
	// prepare temp dir
	tempDir, err := os.MkdirTemp("", "still2cloud-test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if cleanupErr := os.RemoveAll(tempDir); cleanupErr != nil {
			t.Errorf("failed to remove temp dir: %v", cleanupErr)
		}
	}()

	// prepare test data location
	sample, err := os.ReadFile("testdata/1x1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	srcPath := filepath.Join(tempDir, "src.jpg")
	if err = os.WriteFile(srcPath, sample, 0644); err != nil {
		t.Fatal(err)
	}
	destPathLayout := filepath.Join(tempDir, "dest-20060102-150405.jpg")
	latestPath := filepath.Join(tempDir, "latest.txt")

	// create config
	var config Config
	config.Source.Type = SourceTypeFile
	config.Source.Path = srcPath
	config.Destination.Type = DestinationTypeFile
	config.Destination.PathLayout = destPathLayout
	config.Destination.LayoutMode = LayoutModeFileName
	config.Destination.CreateLatestFile = true
	config.Destination.LatestFilePath = latestPath

	// run
	targetDate := time.Date(2025, 12, 31, 23, 58, 59, 0, time.UTC)
	expectedDestPath := filepath.Join(tempDir, "dest-20251231-235859.jpg")
	if err = still2cloud(config, targetDate); err != nil {
		t.Fatal(err)
	}
	if _, err = os.Stat(expectedDestPath); err != nil {
		t.Error(err)
	}
	if _, err = os.Stat(latestPath); err != nil {
		t.Error(err)
	}
	latestContent, err := os.ReadFile(latestPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(latestContent) != expectedDestPath {
		t.Errorf("unexpected latest file content: %s", string(latestContent))
	}
}
