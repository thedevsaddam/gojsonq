package gojsonq

import (
	"testing"
)

func Test_DefaultDecoder(t *testing.T) {
	dd := DefaultDecoder{}
	var user = struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	if err := dd.Decode([]byte(`{"name": "tom", "age": 27}`), &user); err != nil {
		t.Errorf("failed to decode using default decoder: %v", err)
	}

	if user.Name != "tom" || user.Age != 27 {
		t.Error("failed to decode properly by default decoder")
	}

}
