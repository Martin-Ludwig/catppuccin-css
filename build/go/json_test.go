package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestPaletteJsonShouldUnmarshal(t *testing.T) {
	data, err := os.ReadFile("palette.json")
	if err != nil {
		t.Errorf("read error: %v", err)
		return
	}

	var pal Palette
	if err := json.Unmarshal(data, &pal); err != nil {
		t.Errorf("unmarshal error: %v", err)
		return
	}
}
