package mtbfile

import (
	_ "embed"
	"testing"
)

//go:embed tests/fake_MTBFile.json
var fakeMtbfile []byte

func TestShouldDeserializeJson(t *testing.T) {
	_, err := UnmarshalMtbFile(fakeMtbfile)

	if err != nil {
		t.Errorf("Cannot deserialize MTB file")
	}
}
