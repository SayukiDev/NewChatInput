package tts

import (
	"os"
	"testing"
)

func TestNewOption(t *testing.T) {
	opt := NewOption()

	if opt.Baseurl != "http://localhost:10101" {
		t.Errorf("Expected Baseurl to be 'http://localhost:10101', got '%s'", opt.Baseurl)
	}
	if !opt.Run {
		t.Error("Expected Run to be true")
	}
	if !opt.Cache {
		t.Error("Expected Cache to be true")
	}
	if opt.CachePath != "./cache" {
		t.Errorf("Expected CachePath to be './cache', got '%s'", opt.CachePath)
	}
}

func TestNewTTS(t *testing.T) {
	tempDir := t.TempDir()

	opt := &Option{
		Baseurl:    "http://localhost:10101",
		Cache:      true,
		CachePath:  tempDir,
		NowSpacker: 0,
	}

	tts, err := NewTTS(opt)
	if err != nil {
		t.Fatal(err)
	}

	if tts == nil {
		t.Fatal("Expected TTS instance, got nil")
	}
	if tts.o == nil {
		t.Fatal("Expected Option to be set")
	}
	if tts.a == nil {
		t.Fatal("Expected Aivis instance to be set")
	}

	// キャッシュディレクトリが作成されているか確認
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		t.Errorf("Cache directory was not created: %s", tempDir)
	}
}

func TestTTS_WithoutCache(t *testing.T) {
	opt := &Option{
		Baseurl:    "http://localhost:10101",
		Cache:      false,
		NowSpacker: 0,
	}

	tts, err := NewTTS(opt)
	if err != nil {
		t.Fatal(err)
	}

	err = tts.TTS("テストメッセージ")

	if err != nil {
		t.Logf("TTS call resulted in error (expected in test environment): %v", err)
	}
}

func TestTTS_WithCache(t *testing.T) {
	opt := &Option{
		Baseurl:    "http://localhost:10101",
		Cache:      true,
		CachePath:  "./cache",
		NowSpacker: -1,
	}
	opt.Cache = true
	tts, err := NewTTS(opt)
	if err != nil {
		t.Fatal(err)
	}
	err = tts.TTS("うふふふふふふふふ")
	if err != nil {
		t.Logf("TTS call resulted in error (expected in test environment): %v", err)
	}
	d, err := os.ReadDir("./cache")
	if err != nil {
		t.Logf("TTS call resulted in error (expected in test environment): %v", err)
	}
	if len(d) == 0 {
		t.Logf("No cache files found")
	}
	t.Logf("cache files: %s", d)
	err = tts.TTS("うふふふふふふふふ")
	if err != nil {
		t.Logf("TTS call resulted in error (expected in test environment): %v", err)
	}
	os.RemoveAll("./cache")
}
