package gojsonq

import (
	"bytes"
	"encoding/json"
	"testing"
)

func Test_isIndex(t *testing.T) {
	testCases := []struct {
		node     string
		expected bool
	}{
		{
			node:     "items",
			expected: false,
		},
		{
			node:     "[0]",
			expected: true,
		},
		{
			node:     "[101]",
			expected: true,
		},
		{
			node:     "101",
			expected: false,
		},
	}
	for _, tc := range testCases {
		if o := isIndex(tc.node); o != tc.expected {
			t.Errorf("expected: %v got: %v", tc.expected, o)
		}
	}
}

func Test_getIndex(t *testing.T) {
	testCases := []struct {
		node     string
		expected int
	}{
		{
			node:     "Invalid integer",
			expected: -1,
		},
		{
			node:     "item",
			expected: -1,
		},
		{
			node:     "[0]",
			expected: 0,
		},
		{
			node:     "101",
			expected: -1,
		},
		{
			node:     "[101]",
			expected: 101,
		},
	}
	for _, tc := range testCases {
		if o, _ := getIndex(tc.node); o != tc.expected {
			t.Errorf("expected: %v got: %v", tc.expected, o)
		}
	}
}

func Test_toString(t *testing.T) {
	testCases := []struct {
		val      interface{}
		expected string
	}{
		{
			val:      10,
			expected: "10",
		},
		{
			val:      -10,
			expected: "-10",
		},
		{
			val:      10.99,
			expected: "10.99",
		},
		{
			val:      -10.99,
			expected: "-10.99",
		},
		{
			val:      true,
			expected: "true",
		},
	}

	for _, tc := range testCases {
		if o := toString(tc.val); o != tc.expected {
			t.Errorf("expected: %v got: %v", tc.expected, o)
		}
	}
}

func Test_toFloat64(t *testing.T) {
	testCases := []struct {
		val      interface{}
		expected float64
	}{
		{
			val:      10,
			expected: 10,
		},
		{
			val:      int8(1),
			expected: 1,
		},
		{
			val:      int16(91),
			expected: 91,
		},
		{
			val:      int32(88),
			expected: 88,
		},
		{
			val:      int64(898),
			expected: 898,
		},
		{
			val:      float32(99.01),
			expected: 99.01000213623047, // The nearest IEEE754 float32 value of 99.01 is 99.01000213623047; which are not equal (while using ==). Need suggestions for precesion float value.
			// one way to solve the comparison using convertFloat(string with float precesion)==float64
		},
		{
			val:      float32(-99),
			expected: -99,
		},
		{
			val:      float64(-99.91),
			expected: -99.91,
		},
		{
			val:      "",
			expected: 0,
		},
		{
			val:      []int{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		if o, _ := toFloat64(tc.val); o != tc.expected {
			t.Errorf("expected: %v got: %v", tc.expected, o)
		}
	}
}

func Test_sorter(t *testing.T) {
	testCases := []struct {
		tag    string
		asc    bool
		inArr  []interface{}
		outArr []interface{}
	}{
		{
			tag:    "list of string, result hould be in ascending order",
			asc:    true,
			inArr:  []interface{}{"x", "b", "a", "c", "z"},
			outArr: []interface{}{"a", "b", "c", "x", "z"},
		},
		{
			tag:    "list of string, result hould be in descending order",
			asc:    false,
			inArr:  []interface{}{"x", "b", "a", "c", "z"},
			outArr: []interface{}{"z", "x", "c", "b", "a"},
		},
		{
			tag:    "list of float64, result hould be in ascending order",
			asc:    true,
			inArr:  []interface{}{8.0, 7.0, 1.0, 3.0, 5.0, 8.0},
			outArr: []interface{}{1.0, 3.0, 5.0, 7.0, 8.0, 8.0},
		},
		{
			tag:    "list of float64, result hould be in descending order",
			asc:    false,
			inArr:  []interface{}{8.0, 7.0, 1.0, 3.0, 5.0, 8.0},
			outArr: []interface{}{8.0, 8.0, 7.0, 5.0, 3.0, 1.0},
		},
	}

	for _, tc := range testCases {
		obb, _ := json.Marshal(sorter(tc.inArr, tc.asc))
		ebb, _ := json.Marshal(tc.outArr)
		if !bytes.Equal(obb, ebb) {
			t.Errorf("expected: %v got: %v", string(obb), string(ebb))
		}
	}
}
