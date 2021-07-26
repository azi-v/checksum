package checksum

import "testing"

func TestEncode(t *testing.T) {
	str, err := Encode("appSecret", "nonce", "time")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(str)
}
