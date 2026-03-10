package malgo

import "testing"

func TestPlayFromFile(t *testing.T) {
	err := PlayFromFile("test.wav")
	if err != nil {
		t.Fatal(err)
	}
	err = PlayFromFile("test.wav")
	if err != nil {
		t.Fatal(err)
	}
}
