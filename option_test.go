package gojsonq

import (
	"testing"
)

func TestSetDecoder(t *testing.T) {
	jq := New(SetDecoder(&cDecoder{}))
	if jq.option.decoder == nil {
		t.Error("failed to set decoder as option")
	}
}

func TestSetDecoder_with_nil_expecting_an_error(t *testing.T) {
	jq := New(SetDecoder(nil))
	if jq.Error() == nil {
		t.Error("failed to catch nil in SetDecoder")
	}
}

func TestSetSeperator(t *testing.T) {
	jq := New(SetSeperator("->"))
	if jq.option.seperator != "->" {
		t.Error("failed to set seperator as option")
	}
}

func TestSetSeperator_with_nil_expecting_an_error(t *testing.T) {
	jq := New(SetSeperator(""))
	if jq.Error() == nil {
		t.Error("failed to catch nil in SetSeperator")
	}
}
