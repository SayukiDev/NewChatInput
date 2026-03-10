package aivis

import (
	"os"
	"testing"
)

var aivis *Aivis

func init() {
	aivis = New("http://127.0.0.1:10101")
}

func TestAivis_Synthesis(t *testing.T) {
	ss, err := aivis.Speakers()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
	if len(ss) == 0 {
		t.Fatal("expected non-empty speaker list")
	}
	if len(ss[0].Styles) == 0 {
		t.Fatal("expected non-empty style list")
	}
	id := ss[0].Styles[0].ID
	r, err := aivis.AudioQuery("テストテストテストテストテスト", id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	data, err := aivis.Synthesis("テストテストテストテストテストテスト", id, r)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty audio data")
	}
	t.Logf("audio data size: %d bytes", len(data))
	os.WriteFile("test.wav", data, 0644)
}

func TestAivis_Speakers(t *testing.T) {
	r, err := aivis.Speakers()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
