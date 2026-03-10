package aivis

import "testing"

func Test_installModel(t *testing.T) {
	err := installModel("test.aivmx")
	if err != nil {
		t.Fatal(err)
	}
}
