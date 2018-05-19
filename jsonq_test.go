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
	expected := fmt.Sprintf("\nContent: %s\nQuries:%v\n", string(jq.raw), jq.queries)
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

func TestJSONQ_File(t *testing.T) {
	filename := "./data.json"
	fc := createTestFile(t, filename)
	defer fc()

	t.Run("valid_file", func(t *testing.T) {
		if err := New().File(filename).Error(); err != nil {
			t.Error(err)
		}
	})

	t.Run("file_not_exist", func(t *testing.T) {
		if err := New().File("./invalid_file_name").Error(); err == nil {
			t.Error(err)
		}
	})
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
	t.Run("custom_macro", func(t *testing.T) {
		jq.Macro("mac1", func(x, y interface{}) bool {
			return true
		})

		if _, ok := jq.queryMap["mac1"]; !ok {
			t.Error("failed to register macro")
		}
	})

	t.Run("already registered macro", func(t *testing.T) {
		jq.Macro("mac1", func(x, y interface{}) bool {
			return true
		})

		if jq.Error() == nil {
			t.Error("failed to throw error for already registered macro")
		}
	})
}

func TestJSONQ_From(t *testing.T) {
	node := "root.items.[0].name"
	jq := New().From(node)
	if jq.node != node {
		t.Error("failed to set node name")
	}
}

func TestJSONQ_findNode(t *testing.T) {
	t.Run("accessing node", func(t *testing.T) {
		jq := New().JSONString(jsonStr)
		expected := "Star Trek"
		if out := jq.From("vendor.name").Get(); out != expected {
			t.Errorf("Expected: %v\n Got: %v", expected, out)
		}
	})

	t.Run("accessing index", func(t *testing.T) {
		jq := New().JSONString(jsonStr)
		expJSON := `{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}`
		out := jq.From("vendor.items.[0]").Get()
		assertJSON(t, out, expJSON)
	})

	t.Run("accessing not existed index", func(t *testing.T) {
		jq := New().JSONString(jsonStr)
		expJSON := `null`
		out := jq.From("vendor.items.[10]").Get()
		assertJSON(t, out, expJSON)
	})

	t.Run("accessing invalid index error", func(t *testing.T) {
		jq := New().JSONString(jsonStr)
		jq.From("vendor.items.[x]").Get()
		if jq.Error() == nil {
			t.Error("expecting an error")
		}
	})

	t.Run("accessing group by data", func(t *testing.T) {
		jq := New().JSONString(jsonStr)
		expJSON := `[{"id":3,"name":"Sony VAIO","price":1200}]`
		out := jq.From("vendor.items").GroupBy("price").From("1200").Get()
		assertJSON(t, out, expJSON)
	})
}

func TestJSONQ_Where(t *testing.T) {
	t.Run("single Where", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("price", "=", 1700)
		expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("multiple Where expecting data", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("price", "=", 1700).
			Where("id", "=", 2)
		expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("multiple Where expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("price", "=", 1700).
			Where("id", "=", "1700")
		expected := `[]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("Where with invalid operator expecting error", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("price", "invalid_op", 1700)
		jq.Get()

		if jq.Error() == nil {
			t.Error("expecting: invalid operator invalid_op")
		}
	})
}

func TestJSONQ_WhereEqual(t *testing.T) {
	t.Run("single WhereEqual", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereEqual("price", 1700)
		expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("multiple WhereEqual expecting data", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereEqual("price", 1700).
			WhereEqual("id", 2)
		expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("multiple WhereEqual expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereEqual("price", 1700).
			WhereEqual("id", "1700")
		expected := `[]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_WhereNotEqual(t *testing.T) {
	t.Run("single WhereNotEqual", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereNotEqual("price", 850)
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":3,"name":"Sony VAIO","price":1200},{"id":6,"name":"HP core i7","price":950}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("multiple WhereNotEqual expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereNotEqual("price", 850).
			WhereNotEqual("id", 2)
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":3,"name":"Sony VAIO","price":1200},{"id":6,"name":"HP core i7","price":950}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
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

func TestJSONQ_WhereIn(t *testing.T) {
	t.Run("WhereIn expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereIn("id", []int{1, 3, 5})
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":3,"name":"Sony VAIO","price":1200},{"id":5,"key":2300,"name":"HP core i5","price":850}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("WhereIn expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereIn("id", []int{18, 39, 85})
		expected := `[]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_WhereNotIn(t *testing.T) {
	t.Run("WhereIn expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereNotIn("id", []int{1, 3, 5, 6})
		expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":4,"name":"Fujitsu","price":850},{"id":null,"name":"HP core i3 SSD","price":850}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("WhereIn expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereNotIn("price", []float64{850, 950, 1200, 1700, 1350})
		expected := `[]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_OrWhere(t *testing.T) {
	t.Run("OrWhere expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			OrWhere("price", ">", 1200)
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_WhereStartsWith(t *testing.T) {
	t.Run("WhereStartsWith expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereStartsWith("name", "Mac")
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("WhereStartsWith expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereStartsWith("name", "xyz")
		expected := `[]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_WhereEndsWith(t *testing.T) {
	t.Run("WhereStartsWith expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereEndsWith("name", "retina")
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("WhereStartsWith expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereEndsWith("name", "xyz")
		expected := `[]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_WhereContains(t *testing.T) {
	t.Run("WhereContains expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereContains("name", "RetinA")
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("WhereContains expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereContains("name", "xyz")
		expected := `[]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_WhereStrictContains(t *testing.T) {
	t.Run("WhereContains expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereStrictContains("name", "retina")
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("WhereContains expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			WhereStrictContains("name", "RetinA")
		expected := `[]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_GroupBy(t *testing.T) {
	t.Run("WhereContains expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			GroupBy("price")
		expected := `{"1200":[{"id":3,"name":"Sony VAIO","price":1200}],"1350":[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}],"1700":[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}],"850":[{"id":4,"name":"Fujitsu","price":850},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":null,"name":"HP core i3 SSD","price":850}],"950":[{"id":6,"name":"HP core i7","price":950}]}`
		out := jq.Get()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_Sort(t *testing.T) {
	t.Run("sorring array of string in ascending desc", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.names").
			Sort()
		expected := `["Abby","Jane Doe","Jerry","John Doe","Nicolas","Tom"]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("sorring array of float in descending order", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.prices").
			Sort("desc")
		expected := `[2400,2100,1200,400.87,150.1,89.9]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("passing two args in Sort expecting an error", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.prices").
			Sort("asc", "desc")
		jq.Get()
		if jq.Error() == nil {
			t.Error("expecting an error")
		}
	})
}

func TestJSONQ_SortBy(t *testing.T) {
	t.Run("sorring array of object by its key (price-float64) in ascending desc", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			SortBy("price")
		expected := `[{"id":null,"name":"HP core i3 SSD","price":850},{"id":4,"name":"Fujitsu","price":850},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":6,"name":"HP core i7","price":950},{"id":3,"name":"Sony VAIO","price":1200},{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("sorring array of object by its key (price-float64) in descending desc", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			SortBy("price", "desc")
		expected := `[{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":3,"name":"Sony VAIO","price":1200},{"id":6,"name":"HP core i7","price":950},{"id":4,"name":"Fujitsu","price":850},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":null,"name":"HP core i3 SSD","price":850}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("sorring array of object by its key (name-string) in ascending desc", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			SortBy("name")
		expected := `[{"id":4,"name":"Fujitsu","price":850},{"id":null,"name":"HP core i3 SSD","price":850},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":6,"name":"HP core i7","price":950},{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":3,"name":"Sony VAIO","price":1200}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("sorring array of object by its key (name-string) in descending desc", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			SortBy("name", "desc")
		expected := `[{"id":3,"name":"Sony VAIO","price":1200},{"id":2,"name":"MacBook Pro 15 inch retina","price":1700},{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":6,"name":"HP core i7","price":950},{"id":5,"key":2300,"name":"HP core i5","price":850},{"id":null,"name":"HP core i3 SSD","price":850},{"id":4,"name":"Fujitsu","price":850}]`
		out := jq.Get()
		assertJSON(t, out, expected)
	})

	t.Run("passing no argument in SortBy expecting an error", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			SortBy()
		jq.Get()
		if jq.Error() == nil {
			t.Error("expecting an error")
		}
	})

	t.Run("passing more than 2 arguments in SortBy expecting an error", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			SortBy("name", "desc", "asc")
		jq.Get()
		if jq.Error() == nil {
			t.Error("expecting an error")
		}
	})
}

func TestJSONQ_Only(t *testing.T) {
	jq := New().JSONString(jsonStr).
		From("vendor.items").
		Only("id", "price")
	expected := `[{"id":1,"price":1350},{"id":2,"price":1700},{"id":3,"price":1200},{"id":4,"price":850},{"id":5,"price":850},{"id":6,"price":950},{"id":null,"price":850}]`
	out := jq.Get()
	assertJSON(t, out, expected)
}

func TestJSONQ_First(t *testing.T) {
	t.Run("First expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		expected := `{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}`
		out := jq.First()
		assertJSON(t, out, expected)
	})
	t.Run("First expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("price", ">", 1800)
		expected := `null`
		out := jq.First()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_Last(t *testing.T) {
	t.Run("Last expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		expected := `{"id":null,"name":"HP core i3 SSD","price":850}`
		out := jq.Last()
		assertJSON(t, out, expected)
	})

	t.Run("Last expecting empty result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("price", ">", 1800)
		expected := `null`
		out := jq.Last()
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_Nth(t *testing.T) {
	t.Run("Nth expecting result", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		expected := `{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}`
		out := jq.Nth(1)
		assertJSON(t, out, expected)
	})

	t.Run("Nth expecting empty result with an error", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("price", ">", 1800)
		expected := `null`
		out := jq.Nth(1)
		assertJSON(t, out, expected)

		if jq.Error() == nil {
			t.Error("expecting an error for empty result nth value")
		}
	})

	t.Run("Nth expecting empty result with an error of index out of range", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		expected := `null`
		out := jq.Nth(100)
		assertJSON(t, out, expected)

		if jq.Error() == nil {
			t.Error("expecting an error for empty result nth value")
		}
	})

	t.Run("Nth expecting result form last when providing -1", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		expected := `{"id":null,"name":"HP core i3 SSD","price":850}`
		out := jq.Nth(-1)
		assertJSON(t, out, expected)
	})

	t.Run("Nth expecting error is provide 0", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("price", ">", 1800)
		jq.Nth(0)
		if jq.Error() == nil {
			t.Error("expecting error")
		}
	})

	t.Run("Nth expecting empty result if the node is a map", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items.[0]")
		out := jq.Nth(0)
		expected := `null`
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_Find(t *testing.T) {
	t.Run("Find expecting name computers", func(t *testing.T) {
		jq := New().JSONString(jsonStr)
		out := jq.Find("name")
		expected := `"computers"`
		assertJSON(t, out, expected)
	})

	t.Run("Find expecting a nested object", func(t *testing.T) {
		jq := New().JSONString(jsonStr)
		out := jq.Find("vendor.items.[0]")
		expected := `{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}`
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_Pluck(t *testing.T) {
	t.Run("Pluck expecting prices from list of objects", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		out := jq.Pluck("price")
		expected := `[1350,1700,1200,850,850,950,850]`
		assertJSON(t, out, expected)
	})

	t.Run("Pluck expecting empty list from list of objects, because of invalid property name", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		out := jq.Pluck("invalid_prop")
		expected := `[]`
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_Count(t *testing.T) {
	t.Run("Count expecting a int number of total item of an arry", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		out := jq.Count()
		expected := `7`
		assertJSON(t, out, expected)
	})

	t.Run("Count expecting a int number of total item of an arry of objects", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items.[0]")
		out := jq.Count()
		expected := `3`
		assertJSON(t, out, expected)
	})

	t.Run("Count expecting a int number of total item of an arry of groupped objects", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			GroupBy("price")
		out := jq.Count()
		expected := `5`
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_Sum(t *testing.T) {
	t.Run("Sum expecting sum an arry", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.prices")
		out := jq.Sum()
		expected := `6340.87`
		assertJSON(t, out, expected)
	})

	t.Run("Sum expecting sum an arry of objects property", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		out := jq.Sum("price")
		expected := `7750`
		assertJSON(t, out, expected)
	})

	t.Run("Sum expecting an error for providing property for arry", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.prices")
		jq.Sum("key")
		if jq.Error() == nil {
			t.Error("expecting: unnecessary property name for array")
		}
	})

	t.Run("Sum expecting an error for not providing property for arry of objects", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		jq.Sum()
		if jq.Error() == nil {
			t.Error("expecting: property name can not be empty for object")
		}
	})

	t.Run("Sum expecting an error for not providing property for object", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items.[0]")
		jq.Sum()
		if jq.Error() == nil {
			t.Error("expecting: property name can not be empty for object")
		}
	})

	t.Run("Sum expecting an error for not providing property for object", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items.[0]")
		out := jq.Sum("price")
		expected := `1350`
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_Avg(t *testing.T) {
	t.Run("Avg expecting average an arry", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.prices")
		out := jq.Avg()
		expected := `1056.8116666666667`
		assertJSON(t, out, expected)
	})

	t.Run("Avg expecting average an arry of objects property", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		out := jq.Avg("price")
		expected := `1107.142857142857`
		assertJSON(t, out, expected)
	})

}

func TestJSONQ_Min(t *testing.T) {
	t.Run("Min expecting min an arry", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.prices")
		out := jq.Min()
		expected := `89.9`
		assertJSON(t, out, expected)
	})

	t.Run("Min expecting min an arry of objects property", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		out := jq.Min("price")
		expected := `850`
		assertJSON(t, out, expected)
	})
}

func TestJSONQ_Max(t *testing.T) {
	t.Run("Max expecting max an arry", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.prices")
		out := jq.Max()
		expected := `2400`
		assertJSON(t, out, expected)
	})

	t.Run("Max expecting max an arry of objects property", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items")
		out := jq.Max("price")
		expected := `1700`
		assertJSON(t, out, expected)
	})
}

// TODO: Need to write some more combined query test
func TestJSONQ_CombinedWhereOrWhere(t *testing.T) {
	t.Run("combined Where with orWhere", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("id", "=", 1).
			OrWhere("name", "=", "Sony VAIO").
			Where("price", "=", 1200)
		out := jq.Get()
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350},{"id":3,"name":"Sony VAIO","price":1200}]`
		assertJSON(t, out, expected)
	})

	t.Run("combined Where with orWhere containing invalid key", func(t *testing.T) {
		jq := New().JSONString(jsonStr).
			From("vendor.items").
			Where("id", "=", 1).
			OrWhere("invalid_key", "=", "Sony VAIO")
		out := jq.Get()
		expected := `[{"id":1,"name":"MacBook Pro 13 inch retina","price":1350}]`
		assertJSON(t, out, expected)
	})
}
