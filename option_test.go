package gojsonq

import (
	"testing"
)

func TestWithDecoder(t *testing.T) {
	jq := New(WithDecoder(&cDecoder{}))
	if jq.option.decoder == nil {
		t.Error("failed to set decoder as option")
	}
}

func TestWithDecoder_with_nil_expecting_an_error(t *testing.T) {
	jq := New(WithDecoder(nil))
	if jq.Error() == nil {
		t.Error("failed to catch nil in WithDecoder")
	}
}

func TestWithSeparator(t *testing.T) {
	jq := New(WithSeparator("->"))
	if jq.option.separator != "->" {
		t.Error("failed to set separator as option")
	}
}

func TestWithSeparator_with_nil_expecting_an_error(t *testing.T) {
	jq := New(WithSeparator(""))
	if jq.Error() == nil {
		t.Error("failed to catch nil in WithSeparator")
	}
}

// to increase the code coverage; will remove in major release
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

func TestSetSeparator(t *testing.T) {
	jq := New(SetSeparator("->"))
	if jq.option.separator != "->" {
		t.Error("failed to set separator as option")
	}
}

func TestSetSeparator_with_nil_expecting_an_error(t *testing.T) {
	jq := New(SetSeparator(""))
	if jq.Error() == nil {
		t.Error("failed to catch nil in SetSeparator")
	}
}

func TestWithDefaults(t *testing.T) {
	tests := []struct {
		name      string
		defaults  map[string]interface{}
		wantError bool
	}{
		{name: "Nil defaults", wantError: true, defaults: nil},
		{name: "Empty defaults", wantError: false, defaults: map[string]interface{}{}},
		{name: "with defaults value", wantError: false, defaults: map[string]interface{}{"1": 1}},
		{name: "with defaults array", wantError: false, defaults: map[string]interface{}{"1": []int{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jq := New(WithDefaults(tt.defaults))
			if err := jq.Error(); err != nil && !tt.wantError {
				t.Errorf("WithDefaults() = Expected Error but got  %v", err)
			}
		})
	}
}
