package gssh

import (
	"testing"
)

func TestGetPublicKey(t *testing.T) {
	file := "../../test/id_rsa"
	_, err := GetPublicKey(file)
	if err != nil {
		t.Error("Not able to read public key file")
	}
}
