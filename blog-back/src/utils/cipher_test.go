package utils

import (
	"flag"
	"testing"
)

func TestEnDecryptString(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}

	argList := flag.Args()
	if len(argList) != 1 {
		t.Log("Usage: go test -run TestEnDecryptString -v -args stringEncrypted")
		t.Log("Example: go test -run TestEnDecryptString -v -args hello")
	}

	os := argList[0]

	es := EncryptString(os)
	t.Log(es)

	ds := DecryptString(es)
	t.Log(ds)
}
