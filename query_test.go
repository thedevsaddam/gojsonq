package gojsonq

import (
	"testing"
)

func Test_eq(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        9.0, // our expectation for json unmarshalar is little bit different. here 9 provided by user will be equal to float64 9
			y:        9,
			expected: true,
		},
		{
			x:        110,
			y:        120,
			expected: false,
		},
		{
			x:        10.09,
			y:        10.09,
			expected: true,
		},
		{
			x:        10.09,
			y:        10.89,
			expected: false,
		},
		{
			x:        "john",
			y:        "john",
			expected: true,
		},
		{
			x:        "tom",
			y:        "jane",
			expected: false,
		},
		{
			x:        "",
			y:        "",
			expected: true,
		},
		{
			x:        nil,
			y:        nil,
			expected: true,
		},
	}

	for _, tc := range testCases {
		if o, _ := eq(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_neq(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        9.0, // as x is out json unmarshal value which is float64
			y:        9,   // our expectation for json unmarshalar is little bit different. here 9 provided by user will be equal to float64 9
			expected: false,
		},
		{
			x:        110,
			y:        120,
			expected: true,
		},
		{
			x:        10.09,
			y:        10.09,
			expected: false,
		},
		{
			x:        10.09,
			y:        10.89,
			expected: true,
		},
		{
			x:        "john",
			y:        "john",
			expected: false,
		},
		{
			x:        "tom",
			y:        "jane",
			expected: true,
		},
		{
			x:        "",
			y:        "",
			expected: false,
		},
		{
			x:        nil,
			y:        nil,
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := neq(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_gt(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        float64(10),
			y:        5,
			expected: true,
		},
		{
			x:        float64(10),
			y:        15,
			expected: false,
		},
		{
			x:        float64(10),
			y:        "10",
			expected: false,
		},
		{
			x:        "101",
			y:        "101",
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := gt(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_lt(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        float64(10),
			y:        5,
			expected: false,
		},
		{
			x:        float64(10),
			y:        15,
			expected: true,
		},
		{
			x:        float64(10),
			y:        "10",
			expected: false,
		},
		{
			x:        "101",
			y:        "101",
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := lt(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_gte(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        float64(10),
			y:        5,
			expected: true,
		},
		{
			x:        float64(10),
			y:        15,
			expected: false,
		},
		{
			x:        float64(18),
			y:        18,
			expected: true,
		},
		{
			x:        float64(30.9),
			y:        30.9,
			expected: true,
		},
		{
			x:        float64(10),
			y:        "10",
			expected: false,
		},
		{
			x:        "101",
			y:        "101",
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := gte(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_lte(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        float64(10),
			y:        5,
			expected: false,
		},
		{
			x:        float64(10),
			y:        15,
			expected: true,
		},
		{
			x:        float64(18),
			y:        18,
			expected: true,
		},
		{
			x:        float64(30.9),
			y:        30.9,
			expected: true,
		},
		{
			x:        float64(40.9),
			y:        30.9,
			expected: false,
		},
		{
			x:        float64(10),
			y:        "10",
			expected: false,
		},
		{
			x:        "101",
			y:        "101",
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := lte(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_strStrictContains(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        float64(10),
			y:        "10",
			expected: false,
		},
		{
			x:        float64(11),
			y:        float64(11),
			expected: false,
		},
		{
			x:        "131",
			y:        float64(131),
			expected: false,
		},
		{
			x:        "458",
			y:        "458",
			expected: true,
		},
		{
			x:        "arch",
			y:        "arch",
			expected: true,
		},
		{
			x:        "Arch",
			y:        "arch",
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := strStrictContains(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_strContains(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        float64(10),
			y:        "10",
			expected: false,
		},
		{
			x:        float64(11),
			y:        float64(11),
			expected: false,
		},
		{
			x:        "131",
			y:        float64(131),
			expected: false,
		},
		{
			x:        "458",
			y:        "458",
			expected: true,
		},
		{
			x:        "arch",
			y:        "arch",
			expected: true,
		},
		{
			x:        "Arch",
			y:        "arcH",
			expected: true,
		},
	}

	for _, tc := range testCases {
		if o, _ := strContains(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_strStartsWith(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        float64(10),
			y:        "10",
			expected: false,
		},
		{
			x:        float64(11),
			y:        float64(11),
			expected: false,
		},
		{
			x:        "131",
			y:        float64(131),
			expected: false,
		},
		{
			x:        "458",
			y:        "458",
			expected: true,
		},
		{
			x:        "arch",
			y:        "arch",
			expected: true,
		},
		{
			x:        "erik",
			y:        "er",
			expected: true,
		},
	}

	for _, tc := range testCases {
		if o, _ := strStartsWith(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_strEndsWith(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        float64(10),
			y:        "10",
			expected: false,
		},
		{
			x:        float64(11),
			y:        float64(11),
			expected: false,
		},
		{
			x:        "131",
			y:        float64(131),
			expected: false,
		},
		{
			x:        "458",
			y:        "458",
			expected: true,
		},
		{
			x:        "arch",
			y:        "arch",
			expected: true,
		},
		{
			x:        "sky",
			y:        "ky",
			expected: true,
		},
	}

	for _, tc := range testCases {
		if o, _ := strEndsWith(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_in(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        "sky",
			y:        []string{"river", "sun", "moon"},
			expected: false,
		},
		{
			x:        "sun",
			y:        []string{"river", "sun", "moon"},
			expected: true,
		},
		{
			x:        float64(10), // ass json numeric value will be treat as float64
			y:        []int{11, 12, 14, 18},
			expected: false,
		},
		{
			x:        float64(14), // ass json numeric value will be treat as float64
			y:        []int{11, 12, 14, 18},
			expected: true,
		},
		{
			x:        float64(10),
			y:        []float64{11, 12, 14, 18},
			expected: false,
		},
		{
			x:        float64(14),
			y:        []float64{11, 12, 14, 18},
			expected: true,
		},
	}

	for _, tc := range testCases {
		if o, _ := in(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_notIn(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        "sky",
			y:        []string{"river", "sun", "moon"},
			expected: true,
		},
		{
			x:        "sun",
			y:        []string{"river", "sun", "moon"},
			expected: false,
		},
		{
			x:        float64(10), // ass json numeric value will be treat as float64
			y:        []int{11, 12, 14, 18},
			expected: true,
		},
		{
			x:        float64(14), // ass json numeric value will be treat as float64
			y:        []int{11, 12, 14, 18},
			expected: false,
		},
		{
			x:        float64(10),
			y:        []float64{11, 12, 14, 18},
			expected: true,
		},
		{
			x:        float64(14),
			y:        []float64{11, 12, 14, 18},
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := notIn(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_lenEq(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        []interface{}{"river", "sun", "moon"},
			y:        3,
			expected: true,
		},
		{
			x:        "sun",
			y:        3,
			expected: true,
		},
		{
			x:        100,
			y:        3,
			expected: false,
		},
		{
			x:        100,
			y:        3,
			expected: false,
		},
		{
			x:        "moon",
			y:        4.0,
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := lenEq(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_lenNotEq(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        []interface{}{"river", "sun", "moon"},
			y:        6,
			expected: true,
		},
		{
			x:        "sun",
			y:        8,
			expected: true,
		},
		{
			x:        "moon",
			y:        4,
			expected: false,
		},
		{
			x:        100.6,
			y:        4,
			expected: false,
		},
		{
			x:        "moon",
			y:        4.0,
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := lenNotEq(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_lenGt(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        []interface{}{"river", "sun", "moon"},
			y:        6,
			expected: false,
		},
		{
			x:        "sun",
			y:        2,
			expected: true,
		},
		{
			x:        "moon",
			y:        2,
			expected: true,
		},
		{
			x:        100.6,
			y:        4,
			expected: false,
		},
		{
			x:        "moon",
			y:        4.0,
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := lenGt(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_lenLt(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        []interface{}{"river", "sun", "moon"},
			y:        6,
			expected: true,
		},
		{
			x:        "sun",
			y:        20,
			expected: true,
		},
		{
			x:        "moon",
			y:        3,
			expected: false,
		},
		{
			x:        100.6,
			y:        4,
			expected: false,
		},
		{
			x:        "john",
			y:        4.0,
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := lenLt(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_lenGte(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        []interface{}{"river", "sun", "moon"},
			y:        3,
			expected: true,
		},
		{
			x:        "sun",
			y:        2,
			expected: true,
		},
		{
			x:        "jane",
			y:        5,
			expected: false,
		},
		{
			x:        100.6,
			y:        4,
			expected: false,
		},
		{
			x:        "moon",
			y:        4.0,
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := lenGte(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_lenLte(t *testing.T) {
	testCases := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{
			x:        []interface{}{"river", "sun", "moon"},
			y:        3,
			expected: true,
		},
		{
			x:        "sun",
			y:        12,
			expected: true,
		},
		{
			x:        "jane",
			y:        2,
			expected: false,
		},
		{
			x:        100.6,
			y:        4,
			expected: false,
		},
		{
			x:        "moon",
			y:        4.0,
			expected: false,
		},
	}

	for _, tc := range testCases {
		if o, _ := lenLte(tc.x, tc.y); o != tc.expected {
			t.Errorf("for %v expected: %v got: %v", tc.x, tc.expected, o)
		}
	}
}

func Test_loadDefaultQueryMap(t *testing.T) {
	if len(loadDefaultQueryMap()) != 25 {
		t.Error("mismatched default query map size")
	}
}
