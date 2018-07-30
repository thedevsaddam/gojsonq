package gojsonq

import (
	"bytes"
	"encoding/json"
	"testing"
)

func Test_abs(t *testing.T) {
	testCases := []struct {
		data     int
		expected int
	}{
		{
			data:     15,
			expected: 15,
		},
		{
			data:     -25,
			expected: 25,
		},
		{
			data:     0,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		if o := abs(tc.data); o != tc.expected {
			t.Errorf("expected: %v got: %v", tc.expected, o)
		}
	}
}

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
			expected: 99.01000213623047, // The nearest IEEE754 float32 value of 99.01 is 99.01000213623047; which are not equal (while using ==). Need suggestions for precision float value.
			// one way to solve the comparison using convertFloat(string with float precision)==float64
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
			tag:    "list of string, result should be in ascending order",
			asc:    true,
			inArr:  []interface{}{"x", "b", "a", "c", "z"},
			outArr: []interface{}{"a", "b", "c", "x", "z"},
		},
		{
			tag:    "list of string, result should be in descending order",
			asc:    false,
			inArr:  []interface{}{"x", "b", "a", "c", "z"},
			outArr: []interface{}{"z", "x", "c", "b", "a"},
		},
		{
			tag:    "list of float64, result should be in ascending order",
			asc:    true,
			inArr:  []interface{}{8.0, 7.0, 1.0, 3.0, 5.0, 8.0},
			outArr: []interface{}{1.0, 3.0, 5.0, 7.0, 8.0, 8.0},
		},
		{
			tag:    "list of float64, result should be in descending order",
			asc:    false,
			inArr:  []interface{}{8.0, 7.0, 1.0, 3.0, 5.0, 8.0},
			outArr: []interface{}{8.0, 8.0, 7.0, 5.0, 3.0, 1.0},
		},
	}

	for _, tc := range testCases {
		obb, _ := json.Marshal(sortList(tc.inArr, tc.asc))
		ebb, _ := json.Marshal(tc.outArr)
		if !bytes.Equal(obb, ebb) {
			t.Errorf("expected: %v got: %v", string(obb), string(ebb))
		}
	}
}

func Test_sortMap(t *testing.T) {
	testCases := []struct {
		tag     string
		inObjs  interface{}
		outObjs interface{}
		key     string
		asc     bool
	}{
		{
			tag: "should return in ascending order of string value name",
			key: "name",
			asc: true,
			inObjs: []map[string]interface{}{
				{"name": "Z", "height": 5.8},
				{"name": "A", "height": 5.5},
				{"name": "D", "height": 4.9},
				{"name": "X", "height": 5.9},
			},
			outObjs: []map[string]interface{}{
				{"name": "A", "height": 5.5},
				{"name": "D", "height": 4.9},
				{"name": "X", "height": 5.9},
				{"name": "Z", "height": 5.8},
			},
		},
		{
			tag: "should return in descending order of string value name",
			key: "name",
			asc: false,
			inObjs: []map[string]interface{}{
				{"name": "Z", "height": 5.8},
				{"name": "A", "height": 5.5},
				{"name": "D", "height": 4.9},
				{"name": "X", "height": 5.9},
			},
			outObjs: []map[string]interface{}{
				{"name": "Z", "height": 5.8},
				{"name": "X", "height": 5.9},
				{"name": "D", "height": 4.9},
				{"name": "A", "height": 5.5},
			},
		},
		{
			tag: "should return in ascending order of float value height",
			key: "height",
			asc: true,
			inObjs: []map[string]interface{}{
				{"name": "Z", "height": 5.8},
				{"name": "A", "height": 5.5},
				{"name": "D", "height": 4.9},
				{"name": "X", "height": 5.9},
			},
			outObjs: []map[string]interface{}{
				{"name": "D", "height": 4.9},
				{"name": "A", "height": 5.5},
				{"name": "Z", "height": 5.8},
				{"name": "X", "height": 5.9},
			},
		},
		{
			tag: "should return in descending order of float value height",
			key: "height",
			asc: false,
			inObjs: []map[string]interface{}{
				{"name": "Z", "height": 5.8},
				{"name": "A", "height": 5.5},
				{"name": "D", "height": 4.9},
				{"name": "X", "height": 5.9},
			},
			outObjs: []map[string]interface{}{
				{"name": "X", "height": 5.9},
				{"name": "Z", "height": 5.8},
				{"name": "A", "height": 5.5},
				{"name": "D", "height": 4.9},
			},
		},
		{
			key:     "height",
			asc:     false,
			inObjs:  []string{"a", "z", "x"},
			outObjs: []string{"a", "z", "x"},
		},
		{
			key:     "invalid_key",
			asc:     false,
			inObjs:  []string{"x", "z", "a"},
			outObjs: []string{"x", "z", "a"},
		},
	}

	for _, tc := range testCases {
		inObjs := tc.inObjs
		sm := &sortMap{}
		sm.key = tc.key
		sm.desc = !tc.asc
		sm.Sort(inObjs)
		assertInterface(t, inObjs, tc.outObjs, tc.tag)
	}
}

func Test_getNestedValue(t *testing.T) {
	var content interface{}
	if err := json.Unmarshal([]byte(jsonStr), &content); err != nil {
		t.Error("failed to decode json:", err)
	}

	testCases := []struct {
		tag         string
		query       string
		expected    interface{}
		expectError bool
	}{
		{
			tag:         "accessing node",
			query:       "vendor.name",
			expected:    `Star Trek`,
			expectError: false,
		},
		{
			tag:         "should return nil",
			query:       "vendor.xox",
			expected:    nil,
			expectError: true,
		},
		{
			tag:         "should return a map",
			query:       "vendor.items.[0]",
			expected:    map[string]interface{}{"id": 1, "name": "MacBook Pro 13 inch retina", "price": 1350},
			expectError: false,
		},
		{
			tag:         "accessing not existed index",
			query:       "vendor.items.[10]",
			expected:    nil,
			expectError: true,
		},
		{
			tag:         "accessing invalid index error",
			query:       "vendor.items.[x]",
			expected:    nil,
			expectError: true,
		},
		{
			tag:         "should receive valid float value",
			query:       "vendor.items.[0].price",
			expected:    1350,
			expectError: false,
		},
	}

	for _, tc := range testCases {
		out, err := getNestedValue(content, tc.query)
		if tc.expectError && err == nil {
			t.Error("failed to catch error")
		}
		if !tc.expectError {
			assertInterface(t, tc.expected, out, tc.tag)
		}
	}
}

func Test_makeAlias(t *testing.T) {
	testCases := []struct {
		tag   string
		input string
		node  string
		alias string
	}{
		{
			tag:   "scenario 1",
			input: "user.name as uname",
			node:  "user.name",
			alias: "uname",
		},
		{
			tag:   "scenario 2",
			input: "post.title",
			node:  "post.title",
			alias: "title",
		},
		{
			tag:   "scenario 3",
			input: "name",
			node:  "name",
			alias: "name",
		},
	}

	for _, tc := range testCases {
		n, a := makeAlias(tc.input)
		if tc.node != n || tc.alias != a {
			t.Errorf("Tag: %v\nExpected: %v %v \nGot: %v %v\n", tc.tag, tc.node, tc.alias, n, a)
		}
	}
}

func Test_length(t *testing.T) {
	testCases := []struct {
		tag         string
		input       interface{}
		output      int
		errExpected bool
	}{
		{
			tag:         "scenario 1: should return 5 with no error",
			input:       "Hello",
			output:      5,
			errExpected: false,
		},
		{
			tag:         "scenario 2: must return error with -1",
			input:       45,
			output:      -1,
			errExpected: true,
		},
		{
			tag:         "scenario 3: must return length of array",
			input:       []interface{}{"john", "31", false},
			output:      3,
			errExpected: false,
		},
		{
			tag:         "scenario 4: must return length of map",
			input:       map[string]interface{}{"name": "john", "age": 31, "is_designer": false},
			output:      3,
			errExpected: false,
		},
	}

	for _, tc := range testCases {
		out, outErr := length(tc.input)
		if out != tc.output {
			if tc.errExpected && outErr == nil {
				t.Errorf("tag: %s\nExpected: %v\nGot: %v", tc.tag, tc.output, out)
			}
		}
	}
}
