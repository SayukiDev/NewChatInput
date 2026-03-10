package fa

import (
	"os"
	"testing"
)

var fs *FishAudio

func init() {
	fs = NewFishAudio("3ef5a6065ded41a88fb88715d6876946", "s1")
}

func TestFishAudio_TTS(t *testing.T) {
	b, err := fs.TTS(
		"(happy)国境の長いトンネルを抜けると、(break)雪国であった。(break)夜の底が白くなった。(panting)",
		"5161d41404314212af1254556477c17d",
	)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("1.mp3", b, 0644)
}
