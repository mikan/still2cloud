package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func readRTSP(url, path string) ([]byte, error) {
	if path == "" {
		path = "tmp.jpg"
	}
	if err := runFFMPEG(url, path); err != nil {
		return nil, fmt.Errorf("[SRC] failed to exec ffmpeg: %v", err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("[SRC] failed to read %s: %v", path, err)
	}
	if err = os.Remove(path); err != nil {
		return nil, fmt.Errorf("[SRC] failed to remove %s: %v", path, err)
	}
	return content, nil
}

func runFFMPEG(url, path string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	cmd := exec.CommandContext(ctx, "ffmpeg", "-y", "-rtsp_transport", "tcp", "-i", url, "-frames:v", "1", "-r", "1", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Cancel = func() error {
		return cmd.Process.Signal(os.Interrupt)
	}
	cmd.WaitDelay = 10 * time.Second
	return cmd.Run()
}
