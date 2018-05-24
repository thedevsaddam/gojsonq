package gojsonq

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	jq := New()
	if reflect.ValueOf(jq).Type().String() != "*gojsonq.JSONQ" {
		t.Error("failed to match JSONQ type")
	}
}

func TestJSONQ_String(t *testing.T) {
	jq := New()
	expected := fmt.Sprintf("\nContent: %s\nQueries:%v\n", string(jq.raw), jq.queries)
	if out := jq.String(); out != expected {
		t.Errorf("Expected: %v\n Got: %v", expected, out)
	}
}

func TestJSONQ_decode(t *testing.T) {
	testCases := []struct {
		tag       string
		jsonStr   string
		errExpect bool
	}{
		{
			tag:       "valid json",
			jsonStr:   `{"name": "John Doe", "age": 30}`,
			errExpect: false,
		},
		{
			tag:       "invalid json should return error",
			jsonStr:   `{"name": "John Doe", "age": 30, "only_key"}`,
			errExpect: true,
		},
	}

	for _, tc := range testCases {
		jq := New()
		jq.raw = json.RawMessage(tc.jsonStr)
		jq.decode()
		if err := jq.Error(); err != nil && !tc.errExpect {
			t.Errorf("failed %s", tc.tag)
		}
	}
}

func TestJSONQ_Copy(t *testing.T) {
	jq := New()
	mp := map[string]int{}
	for i := 0; i < 100; i++ {
		adr := fmt.Sprintf("%p", jq.Copy())
		if _, ok := mp[adr]; ok {
			t.Error("failed to copy JSONQ")
		} else {
			mp[adr] = i
		}
	}
}

func TestJSONQ_File(t *testing.T) {
	filename := "./data.json"
	fc := createTestFile(t, filename)
	defer fc()

	testCases := []struct {
		tag         string
		filename    string
		expectedErr bool
	}{
		{
			tag:         "valid file name does not expect error",
			filename:    filename,
			expectedErr: false,
		},
		{
			tag:         "invalid valid file name expecting error",
			filename:    "invalid_file.xjson",
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		err := New().File(tc.filename).Error()
		if tc.expectedErr && err == nil {
			t.Errorf("%s", tc.tag)
		}
	}
}

func TestJSONQ_JSONString(t *testing.T) {
	testCases := []struct {
		tag       string
		jsonStr   string
		errExpect bool
	}{
		{
			tag:       "valid json",
			jsonStr:   `{"name": "John Doe", "age": 30}`,
			errExpect: false,
		},
		{
			tag:       "invalid json should return error",
			jsonStr:   `{"name": "John Doe", "age": 30, "only_key"}`,
			errExpect: true,
		},
	}

	for _, tc := range testCases {
		if err := New().JSONString(tc.jsonStr).Error(); err != nil && !tc.errExpect {
			t.Errorf("failed %s", tc.tag)
		}
	}
}

func TestJSONQ_Reader(t *testing.T) {
	testCases := []struct {
		tag       string
		jsonStr   string
		errExpect bool
	}{
		{
			tag:       "valid json",
			jsonStr:   `{"name": "John Doe", "age": 30}`,
			errExpect: false,
		},
		{
			tag:       "invalid json should return error",
			jsonStr:   `{"name": "John Doe", "age": 30, "only_key"}`,
			errExpect: true,
		},
	}

	for _, tc := range testCases {
		rdr := strings.NewReader(tc.jsonStr)
		if err := New().Reader(rdr).Error(); err != nil && !tc.errExpect {
			t.Errorf("failed %s", tc.tag)
		}
	}
}

func TestJSONQ_Errors(t *testing.T) {
	testCases := []struct {
		tag     string
		jsonStr string
	}{
		{
			tag:     "invalid json 1",
			jsonStr: `{"name": "John Doe", "age": 30, :""}`,
		},
		{
			tag:     "invalid json 2",
			jsonStr: `{"name": "John Doe", "age": 30, "only_key"}`,
		},
	}

	for _, tc := range testCases {
		if errs := New().JSONString(tc.jsonStr).Errors(); len(errs) == 0 {
			t.Errorf("failed %s", tc.tag)
		}
	}
}

func TestJSONQ_Macro(t *testing.T) {
	jq := New()
	jq.Macro("mac1", func(x, y interface{}) (bool, error) {
		return true, nil
	})

	if _, ok := jq.queryMap["mac1"]; !ok {
		t.Error("failed to register macro")
	}

	jq.Macro("mac1", func(x, y interface{}) (bool, error) {
		return true, nil
	})
	if jq.Error() == nil {
		t.Error("failed to throw error for already registered macro")
	}
}

func TestJSONQ_From(t *testing.T) {
	node := "root.items.[0].name"
	jq := New().From(node)
	if jq.node != node {
		t.Error("failed to set node name")
	}
}

func TestJSONQ_reset(t *testing.T) {
	node := "root.items"
	jq := New().From(node).WhereEqual("price", "1900").WhereEqual("id", 1)
	jq.reset()
	if len(jq.queries) != 0 || jq.queryIndex != 0 {
		t.Error("reset failed")
	}
}

func TestJSONQ_Reset(t *testing.T) {
	node := "root.items"
	jq := New().From(node).WhereEqual("price", "1900").WhereEqual("id", 1)
	jq.Reset()
	if len(jq.queries) != 0 || jq.queryIndex != 0 {
		t.Error("reset failed")
	}
}

func TestJSONQ_findNode(t *testing.T) {
	testCases := []struct {
		tag         string
		query       string
		expected    string
		expectError bool
	}{
		{
			tag:         "accessing node",
			query:       "vendor.name",
			expected:    `"Star Trek"`,
			expectError: false,
		},
		{
			tag:         "accessing not existed index",
			query:       "vendor.items.[0]",
			expected:    `{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}`,
			expectError: false,
		},
		{
			tag:         "accessing not existed index",
			query:       "vendor.items.[10]",
			expected:    `null`,
			expectError: false,
		},
		{
			tag:         "accessing invalid index error",
			query:       "vendor.items.[x]",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		jq := New().JSONString(jsonStr)
		out := jq.From(tc.query).Get()
		if tc.expectError && jq.Error() == nil {
			t.Error("failed to catch error")
		}
		if !tc.expectError {
			assertJSON(t, out, tc.expected, tc.tag)
		}
	}

	jq := New().JSONString(jsonStr)
	expJSON := `[{"id":3,"name":"Sony VAIO","price":1200}]`
	out := jq.From("vendor.items").GroupBy("price").From("1200").Get()
	assertJSON(t, out, expJSON, "accessing group by data")
}

func TestJSONQ_Where_single_where(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("price", "=", 1700)
	expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "single Where")
}

func TestJSONQ_Where_multiple_where_expecting_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("price", "=", 1700).
		Where("id", "=", 2)
	expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "multiple Where expecting data")
}

func TestJSONQ_Where_multiple_where_expecting_empty_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("price", "=", 1700).
		Where("id", "=", "1700")
	expected := `[]`
	out := jq.Get()
	assertJSON(t, out, expected, "multiple Where expecting empty result")
}

func TestJSONQ_Where_multiple_where_with_invalid_operator_expecting_error(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("price", "invalid_op", 1700)
	jq.Get()

	if jq.Error() == nil {
		t.Error("expecting: invalid operator invalid_op")
	}
}

func TestJSONQ_single_WhereEqual(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereEqual("price", 1700)
	expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "single WhereEqual")
}

func TestJSONQ_multiple_WhereEqual_expecting_data(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereEqual("price", 1700).
		WhereEqual("id", 2)
	expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "multiple WhereEqual expecting data")
}

func TestJSONQ_multiple_WhereEqual_expecting_empty_data(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereEqual("price", 1700).
		WhereEqual("id", "1700")
	expected := `[]`
	out := jq.Get()
	assertJSON(t, out, expected, "multiple WhereEqual expecting empty result")
}

func TestJSONQ_single_WhereNotEqual(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereNotEqual("price", 850)
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":3,"name":"Sony VAIO","price":1200},{"id":6,"name":"HP core i7","price":950}]`
	out := jq.Get()
	assertJSON(t, out, expected, "single WhereNotEqual")
}

func TestJSONQ_multiple_WhereNotEqual(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereNotEqual("price", 850).
		WhereNotEqual("id", 2)
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":3,"name":"Sony VAIO","price":1200},{"id":6,"name":"HP core i7","price":950}]`
	out := jq.Get()
	assertJSON(t, out, expected, "multiple WhereNotEqual expecting result")
}

func TestJSONQ_WhereNil(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereNil("id")
	expected := `[{"id":null,"name":"HP core i3 SSD","price":850}]`
	out := jq.Get()
	assertJSON(t, out, expected)
}

func TestJSONQ_WhereNotNil(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereNotNil("id")
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":3,"name":"Sony VAIO","price":1200},{"id":4,"name":"Fujitsu","price":850},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":6,"name":"HP core i7","price":950}]`
	out := jq.Get()
	assertJSON(t, out, expected)
}

func TestJSONQ_WhereIn_expecting_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereIn("id", []int{1, 3, 5})
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":3,"name":"Sony VAIO","price":1200},{"id":5,"key":2300,"name":"HP core i5","price":850}]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereIn expecting result")
}

func TestJSONQ_WhereIn_expecting_empty_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereIn("id", []int{18, 39, 85})
	expected := `[]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereIn expecting empty result")
}

func TestJSONQ_WhereNotIn_expecting_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereNotIn("id", []int{1, 3, 5, 6})
	expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":4,"name":"Fujitsu","price":850},{"id":null,"name":"HP core i3 SSD","price":850}]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereIn expecting result")
}

func TestJSONQ_WhereNotIn_expecting_empty_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereNotIn("price", []float64{850, 950, 1200, 1700, 1350})
	expected := `[]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereIn expecting empty result")
}

func TestJSONQ_OrWhere(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		OrWhere("price", ">", 1200)
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "OrWhere expecting result")
}

func TestJSONQ_WhereStartsWith_expecting_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereStartsWith("name", "Mac")
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereStartsWith expecting result")
}

func TestJSONQ_WhereStartsWith_expecting_empty_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereStartsWith("name", "xyz")
	expected := `[]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereStartsWith expecting empty result")
}

func TestJSONQ_WhereEndsWith(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereEndsWith("name", "retina")
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereStartsWith expecting result")
}

func TestJSONQ_WhereEndsWith_empty_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereEndsWith("name", "xyz")
	expected := `[]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereStartsWith expecting empty result")
}

func TestJSONQ_WhereContains_expecting_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereContains("name", "RetinA")
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereContains expecting result")
}

func TestJSONQ_WhereContains_expecting_empty_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereContains("name", "xyz")
	expected := `[]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereContains expecting empty result")
}

func TestJSONQ_WhereStrictContains_expecting_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereStrictContains("name", "retina")
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereContains expecting result")
}

func TestJSONQ_WhereStrictContains_expecting_empty_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		WhereStrictContains("name", "RetinA")
	expected := `[]`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereContains expecting empty result")
}

func TestJSONQ_GroupBy(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		GroupBy("price")
	expected := `{"1200":[{"id":3,"name":"Sony VAIO","price":1200}],"1350":[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}],"1700":[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}],"850":[{"id":4,"name":"Fujitsu","price":850},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":null,"name":"HP core i3 SSD","price":850}],"950":[{"id":6,"name":"HP core i7","price":950}]}`
	out := jq.Get()
	assertJSON(t, out, expected, "WhereContains expecting result")
}

func TestJSONQ_Sort_string_ascending_order(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.names").
		Sort()
	expected := `["Abby","Jane Doe","Jerry","John Doe","Nicolas","Tom"]`
	out := jq.Get()
	assertJSON(t, out, expected, "sorting array of string in ascending desc")
}

func TestJSONQ_Sort_float64_descending_order(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.prices").
		Sort("desc")
	expected := `[2400,2100,1200,400.87,150.1,89.9]`
	out := jq.Get()
	assertJSON(t, out, expected, "sorting array of float in descending order")
}

func TestJSONQ_Sort_with_two_args_expecting_error(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.prices").
		Sort("asc", "desc")
	jq.Get()
	if jq.Error() == nil {
		t.Error("expecting an error")
	}
}

func TestJSONQ_SortBy_float_ascending_order(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		SortBy("price")
	expected := `[{"id":null,"name":"HP core i3 SSD","price":850},{"id":4,"name":"Fujitsu","price":850},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":6,"name":"HP core i7","price":950},{"id":3,"name":"Sony VAIO","price":1200},{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
	out := jq.Get()
	assertJSON(t, out, expected, "sorting array of object by its key (price-float64) in ascending desc")
}

func TestJSONQ_SortBy_float_descending_order(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		SortBy("price", "desc")
	expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":3,"name":"Sony VAIO","price":1200},{"id":6,"name":"HP core i7","price":950},{"id":4,"name":"Fujitsu","price":850},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":null,"name":"HP core i3 SSD","price":850}]`
	out := jq.Get()
	assertJSON(t, out, expected, "sorting array of object by its key (price-float64) in descending desc")
}

func TestJSONQ_SortBy_string_ascending_order(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		SortBy("name")
	expected := `[{"id":4,"name":"Fujitsu","price":850},{"id":null,"name":"HP core i3 SSD","price":850},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":6,"name":"HP core i7","price":950},{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":3,"name":"Sony VAIO","price":1200}]`
	out := jq.Get()
	assertJSON(t, out, expected, "sorting array of object by its key (name-string) in ascending desc")
}

func TestJSONQ_SortBy_string_descending_order(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		SortBy("name", "desc")
	expected := `[{"id":3,"name":"Sony VAIO","price":1200},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":6,"name":"HP core i7","price":950},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":null,"name":"HP core i3 SSD","price":850},{"id":4,"name":"Fujitsu","price":850}]`
	out := jq.Get()
	assertJSON(t, out, expected, "sorting array of object by its key (name-string) in descending desc")
}

func TestJSONQ_SortBy_no_argument_expecting_error(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		SortBy()
	jq.Get()
	if jq.Error() == nil {
		t.Error("expecting an error")
	}
}

func TestJSONQ_SortBy_more_than_two_argument_expecting_error(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		SortBy("name", "desc", "asc")
	jq.Get()
	if jq.Error() == nil {
		t.Error("expecting an error")
	}
}

func TestJSONQ_SortBy_expecting_as_provided_node_is_not_list(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("name").
		SortBy("name", "desc")
	out := jq.Get()
	expJSON := `"computers"`
	assertJSON(t, out, expJSON)
}

func TestJSONQ_SortBy_expecting_empty_as_provided_node_is_not_list(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").Where("price", ">", 2500).
		SortBy("name", "desc")
	out := jq.Get()
	expJSON := `[]`
	assertJSON(t, out, expJSON)
}

func TestJSONQ_Only(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Only("id", "price")
	expected := `[{"id":1,"price":1350},{"id":2,"price":1700},{"id":3,"price":1200},{"id":4,"price":850},{"id":5,"price":850},{"id":6,"price":950},{"id":null,"price":850}]`
	out := jq.Get()
	assertJSON(t, out, expected)
}

func TestJSONQ_First_expecting_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	expected := `{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}`
	out := jq.First()
	assertJSON(t, out, expected, "First expecting result")
}

func TestJSONQ_First_expecting_empty_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("price", ">", 1800)
	expected := `null`
	out := jq.First()
	assertJSON(t, out, expected, "First expecting empty result")
}

func TestJSONQ_Last_expecting_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	expected := `{"id":null,"name":"HP core i3 SSD","price":850}`
	out := jq.Last()
	assertJSON(t, out, expected, "Last expecting result")
}

func TestJSONQ_Last_expecting_empty_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("price", ">", 1800)
	expected := `null`
	out := jq.Last()
	assertJSON(t, out, expected, "Last expecting empty result")
}

func TestJSONQ_Nth_expecting_result(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	expected := `{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}`
	out := jq.Nth(1)
	assertJSON(t, out, expected, "Nth expecting result")
}

func TestJSONQ_Nth_expecting_empty_result_with_error(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("price", ">", 1800)
	expected := `null`
	out := jq.Nth(1)
	assertJSON(t, out, expected, "Nth expecting empty result with an error")

	if jq.Error() == nil {
		t.Error("expecting an error for empty result nth value")
	}
}

func TestJSONQ_Nth_expecting_empty_result_with_error_index_out_of_range(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	expected := `null`
	out := jq.Nth(100)
	assertJSON(t, out, expected, "Nth expecting empty result with an error of index out of range")

	if jq.Error() == nil {
		t.Error("expecting an error for empty result nth value")
	}
}

func TestJSONQ_Nth_expecting_result_from_last_using_negative_index(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	expected := `{"id":null,"name":"HP core i3 SSD","price":850}`
	out := jq.Nth(-1)
	assertJSON(t, out, expected, "Nth expecting result form last when providing -1")
}

func TestJSONQ_Nth_expecting_error_providing_zero_as_index(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("price", ">", 1800)
	jq.Nth(0)
	if jq.Error() == nil {
		t.Error("expecting error")
	}
}

func TestJSONQ_Nth_expecting_empty_result_as_node_is_map(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items.[0]")
	out := jq.Nth(0)
	expected := `null`
	assertJSON(t, out, expected, "Nth expecting empty result if the node is a map")
}
func TestJSONQ_Find_simple_property(t *testing.T) {
	jq := New().JSONString(jsonStr)
	out := jq.Find("name")
	expected := `"computers"`
	assertJSON(t, out, expected, "Find expecting name computers")
}

func TestJSONQ_Find_nested_property(t *testing.T) {
	jq := New().JSONString(jsonStr)
	out := jq.Find("vendor.items.[0]")
	expected := `{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}`
	assertJSON(t, out, expected, "Find expecting a nested object")
}

func TestJSONQ_Pluck_expecting_list_of_float64(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	out := jq.Pluck("price").Get()
	expected := `[1350,1700,1200,850,850,950,850]`
	assertJSON(t, out, expected, "Pluck expecting prices from list of objects")
}

func TestJSONQ_Pluck_expecting_empty_list_of_float64(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	out := jq.Pluck("invalid_prop").Get()
	expected := `[]`
	assertJSON(t, out, expected, "Pluck expecting empty list from list of objects, because of invalid property name")
}

func TestJSONQ_Count_expecting_int_from_list(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	out := jq.Count()
	expected := `7`
	assertJSON(t, out, expected, "Count expecting a int number of total item of an array")
}

func TestJSONQ_Count_expecting_int_from_list_of_objects(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items.[0]")
	out := jq.Count()
	expected := `3`
	assertJSON(t, out, expected, "Count expecting a int number of total item of an array of objects")
}

func TestJSONQ_Count_expecting_int_from_objects(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		GroupBy("price")
	out := jq.Count()
	expected := `5`
	assertJSON(t, out, expected, "Count expecting a int number of total item of an array of grouped objects")
}

func TestJSONQ_Sum_of_array_numeric_values(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.prices")
	out := jq.Sum()
	expected := `6340.87`
	assertJSON(t, out, expected, "Sum expecting sum an array")
}

func TestJSONQ_Sum_of_array_objects_property_numeric_values(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	out := jq.Sum("price")
	expected := `7750`
	assertJSON(t, out, expected, "Sum expecting sum an array of objects property")
}

func TestJSONQ_Sum_expecting_error_for_providing_property_of_array(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.prices")
	jq.Sum("key")
	if jq.Error() == nil {
		t.Error("expecting: unnecessary property name for array")
	}
}

func TestJSONQ_Sum_expecting_error_for_not_providing_property_of_array_of_objects(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	jq.Sum()
	if jq.Error() == nil {
		t.Error("expecting: property name can not be empty for object")
	}
}

func TestJSONQ_Sum_expecting_error_for_not_providing_property_of_object(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items.[0]")
	jq.Sum()
	if jq.Error() == nil {
		t.Error("expecting: property name can not be empty for object")
	}
}

func TestJSONQ_Sum_expecting_result_from_nested_object(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items.[0]")
	out := jq.Sum("price")
	expected := `1350`
	assertJSON(t, out, expected, "Sum expecting result from nested object")
}

func TestJSONQ_Avg_array(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.prices")
	out := jq.Avg()
	expected := `1056.8116666666667`
	assertJSON(t, out, expected, "Avg expecting average an array")
}

func TestJSONQ_Avg_array_of_objects(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	out := jq.Avg("price")
	expected := `1107.142857142857`
	assertJSON(t, out, expected, "Avg expecting average an array of objects property")
}

func TestJSONQ_Min_array(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.prices")
	out := jq.Min()
	expected := `89.9`
	assertJSON(t, out, expected, "Min expecting min an array")
}

func TestJSONQ_Min_array_of_objects(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	out := jq.Min("price")
	expected := `850`
	assertJSON(t, out, expected, "Min expecting min an array of objects property")
}

func TestJSONQ_Max_array(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.prices")
	out := jq.Max()
	expected := `2400`
	assertJSON(t, out, expected, "Max expecting max an array")
}

func TestJSONQ_Max_array_of_objects(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items")
	out := jq.Max("price")
	expected := `1700`
	assertJSON(t, out, expected, "Max expecting max an array of objects property")
}

// TODO: Need to write some more combined query test
func TestJSONQ_CombinedWhereOrWhere(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("id", "=", 1).
		OrWhere("name", "=", "Sony VAIO").
		Where("price", "=", 1200)
	out := jq.Get()
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":3,"name":"Sony VAIO","price":1200}]`
	assertJSON(t, out, expected, "combined Where with orWhere")
}

func TestJSONQ_CombinedWhereOrWhere_invalid_key(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Where("id", "=", 1).
		OrWhere("invalid_key", "=", "Sony VAIO")
	out := jq.Get()
	expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}]`
	assertJSON(t, out, expected, "combined Where with orWhere containing invalid key")
}
