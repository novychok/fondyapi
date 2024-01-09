package main

import (
	"testing"
)

func TestSetSignature(t *testing.T) {
	var testCase = struct {
		signature string
	}{
		signature: "94b68dcc010bbf99a5c9eb03f9add35b5c70e003",
	}

	requestObj := NewRequestObj("123", "test order", "USD",
		"1250", "1396424")

	requestObj.SetSignature("test")

	if testCase.signature != requestObj.Request.Signature {
		t.Errorf("expected signature %v, got: %v", testCase.signature, requestObj.Request.Signature)
	}

	if testCase.signature == requestObj.Request.Signature {
		t.Skipped()
	}
}
