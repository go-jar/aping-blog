package utils

import "testing"

func TestEnDecryptString(t *testing.T) {
	os := "blog441862?"

	es := EncryptString(os)
	t.Log(es)

	ds := DecryptString(es)
	t.Log(ds)

	if ds != es {
		t.Error("decode error")
	}
}
