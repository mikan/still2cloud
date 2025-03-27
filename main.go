package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var ver = "v0.0.0" // filled by "make build"

func main() {
	configPath := flag.String("c", "still2cloud.json", "path to configuration file")
	printVersion := flag.Bool("v", false, "print version and exit")
	flag.Parse()
	if *printVersion {
		fmt.Printf("still2cloud %s compiled by %s\n", ver, runtime.Version())
		os.Exit(0)
	}
	now := time.Now().UTC()
	rawConfig, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("[CONF] failed to read %s: %v", *configPath, err)
	}
	var config Config
	if err = json.Unmarshal(rawConfig, &config); err != nil {
		log.Fatalf("[CONF] failed to parse %s: %v", *configPath, err)
	}
	if err = still2cloud(config, now); err != nil {
		log.Fatal(err)
	}
}

func still2cloud(config Config, now time.Time) error {
	// Source
	var err error
	var content []byte
	switch config.Source.Type {
	case SourceTypeHTTP:
		content, err = fetchHTTP(config)
	case SourceTypeFile:
		content, err = readFile(config.Source.Path)
	case SourceTypeRTSP:
		content, err = readRTSP(config.Source.URL, config.Source.Path)
	case SourceTypeRPi:
		content, err = readRPi(config.Source.Path)
	default:
		return fmt.Errorf("[SRC] unsupported source type: %s", config.Source.Type)
	}
	if err != nil {
		return fmt.Errorf("[SRC] failed to transport source: %w", err)
	}

	// Convert
	if config.Convert.Width > 0 && config.Convert.Height > 0 {
		log.Print("[CONV] resizing is not implemented yet")
	}
	if config.Convert.Format != "" {
		log.Print("[CONV] formatting is not implemented yet")
	}

	// Destination
	path := formatLayout(config.Destination.PathLayout, now, config.Destination.LayoutMode)
	switch config.Destination.Type {
	case DestinationTypeS3:
		err = putS3Object(context.Background(), config, path, content, false)
	case DestinationTypeFile:
		err = writeFile(path, content)
	default:
		return fmt.Errorf("[DEST] unsupported destination type: %s", config.Destination.Type)
	}
	if err != nil {
		return fmt.Errorf("[DEST] failed to transport destination: %v", err)
	}

	// Latest file
	if config.Destination.CreateLatestFile {
		if config.Destination.LatestFilePath == "" {
			return errors.New("[DEST] latest_file_path is required")
		}
		switch config.Destination.Type {
		case DestinationTypeS3:
			err = putS3Object(context.Background(), config, config.Destination.LatestFilePath, []byte(path), true)
		case DestinationTypeFile:
			err = writeFile(config.Destination.LatestFilePath, []byte(path))
		}
		if err != nil {
			log.Fatalf("[DEST] failed to create %s: %v", config.Destination.LatestFilePath, err)
		}
	}
	return nil
}

func formatLayout(layout string, timestamp time.Time, mode LayoutMode) string {
	switch mode {
	default:
		fallthrough
	case LayoutModeAll:
		return timestamp.Format(layout)
	case LayoutModeDisable:
		return layout
	case LayoutModeFileName:
		base := filepath.Base(layout)
		return strings.Replace(layout, base, timestamp.Format(base), 1)
	}
}
