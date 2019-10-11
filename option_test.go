package gojsonq

import (
	"testing"
)

func TestSetDecoder(t *testing.T) {
	jq := New(WithDecoder(&cDecoder{}))
	if jq.option.decoder == nil {
		t.Error("failed to set decoder as option")
	}
}

func TestSetDecoder_with_nil_expecting_an_error(t *testing.T) {
	jq := New(WithDecoder(nil))
	if jq.Error() == nil {
		t.Error("failed to catch nil in SetDecoder")
	}
}

func TestSetSeparator(t *testing.T) {
	jq := New(WithSeparator("->"))
	if jq.option.separator != "->" {
		t.Error("failed to set separator as option")
	}
}

func TestSetSeparator_with_nil_expecting_an_error(t *testing.T) {
	jq := New(WithSeparator(""))
	if jq.Error() == nil {
		t.Error("failed to catch nil in SetSeparator")
	}
}
