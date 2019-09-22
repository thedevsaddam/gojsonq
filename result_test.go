package gojsonq

import (
	"reflect"
	"testing"
	"time"
)

func TestNewResult(t *testing.T) {
	result := NewResult("gojsonq")
	if reflect.ValueOf(result).Type().String() != "*gojsonq.Result" {
		t.Error("failed to match gojsonq.Result type")
	}
}

func TestNil(t *testing.T) {
	result := NewResult(nil)
	if result.Nil() == false {
		t.Error("failed to check Nil")
	}
}
func TestBool(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect bool
		errExpect bool
	}{
		{tag: "bool value as expected", value: true, valExpect: true, errExpect: false},
		{tag: "invalid bool, error expected", value: 123, valExpect: false, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Bool()
		if err != nil && !tc.errExpect {
			t.Error("bool:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestTime(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	tm, err := time.Parse(layout, str)
	if err != nil {
		t.Error("failed to parse time:", err)
	}
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect time.Time
		errExpect bool
	}{
		{tag: "time value as expected", value: "2014-11-12T11:45:26.371Z", valExpect: tm, errExpect: false},
		{tag: "invalid time, error expected", value: "2014-11-12", valExpect: time.Time{}, errExpect: true},
		{tag: "invalid time, error expected", value: 12322, valExpect: time.Time{}, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Time(layout)
		if err != nil && !tc.errExpect {
			t.Error("time:", err)
		}
		if !reflect.DeepEqual(v, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestDuration(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect time.Duration
		errExpect bool
	}{
		{tag: "duration value as expected", value: "10s", valExpect: time.Duration(10 * time.Second), errExpect: false},
		{tag: "duration value as expected", value: "10m", valExpect: time.Duration(10 * time.Minute), errExpect: false},
		{tag: "duration value as expected", value: float64(10), valExpect: time.Duration(10 * time.Nanosecond), errExpect: false}, // go decode number to float64
		{tag: "invalid duration, error expected", value: "1", valExpect: time.Duration(0), errExpect: true},
		{tag: "invalid duration, error expected", value: 1, valExpect: time.Duration(0), errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Duration()
		if err != nil && !tc.errExpect {
			t.Error("duration:", err)
		}
		if !reflect.DeepEqual(v, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect string
		errExpect bool
	}{
		{tag: "string value as expected", value: "hello", valExpect: "hello", errExpect: false},
		{tag: "invalid string, error expected", value: 123, valExpect: "", errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).String()
		if err != nil && !tc.errExpect {
			t.Error("string: ", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestInt(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect int
		errExpect bool
	}{
		{tag: "int value as expected", value: 123.8, valExpect: 123, errExpect: false},
		{tag: "int value as expected", value: 12.3, valExpect: 12, errExpect: false},
		{tag: "invalid int, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Int()
		if err != nil && !tc.errExpect {
			t.Error("int:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestInt8(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect int8
		errExpect bool
	}{
		{tag: "int8 value as expected", value: 123.8, valExpect: int8(123), errExpect: false},
		{tag: "int8 value as expected", value: 12.3, valExpect: int8(12), errExpect: false},
		{tag: "invalid int8, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Int8()
		if err != nil && !tc.errExpect {
			t.Error("int8:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestInt16(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect int16
		errExpect bool
	}{
		{tag: "int16 value as expected", value: 123.8, valExpect: int16(123), errExpect: false},
		{tag: "int16 value as expected", value: 12.3, valExpect: int16(12), errExpect: false},
		{tag: "invalid int16, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Int16()
		if err != nil && !tc.errExpect {
			t.Error("int16:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestInt32(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect int32
		errExpect bool
	}{
		{tag: "int32 value as expected", value: 123.8, valExpect: int32(123), errExpect: false},
		{tag: "int32 value as expected", value: 12.3, valExpect: int32(12), errExpect: false},
		{tag: "invalid int32, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Int32()
		if err != nil && !tc.errExpect {
			t.Error("int32:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestInt64(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect int64
		errExpect bool
	}{
		{tag: "int64 value as expected", value: 123.8, valExpect: int64(123), errExpect: false},
		{tag: "int64 value as expected", value: 12.3, valExpect: int64(12), errExpect: false},
		{tag: "invalid int64, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Int64()
		if err != nil && !tc.errExpect {
			t.Error("int64:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestUint(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect uint
		errExpect bool
	}{
		{tag: "uint value as expected", value: 123.8, valExpect: uint(123), errExpect: false},
		{tag: "uint value as expected", value: 12.3, valExpect: uint(12), errExpect: false},
		{tag: "invalid uint, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Uint()
		if err != nil && !tc.errExpect {
			t.Error("uint:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestUint8(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect uint8
		errExpect bool
	}{
		{tag: "uint8 value as expected", value: 123.8, valExpect: uint8(123), errExpect: false},
		{tag: "uint8 value as expected", value: 12.3, valExpect: uint8(12), errExpect: false},
		{tag: "invalid uint8, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Uint8()
		if err != nil && !tc.errExpect {
			t.Error("uint8:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestUint16(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect uint16
		errExpect bool
	}{
		{tag: "uint16 value as expected", value: 123.8, valExpect: uint16(123), errExpect: false},
		{tag: "uint16 value as expected", value: 12.3, valExpect: uint16(12), errExpect: false},
		{tag: "invalid uint16, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Uint16()
		if err != nil && !tc.errExpect {
			t.Error("uint16:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestUint32(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect uint32
		errExpect bool
	}{
		{tag: "uint32 value as expected", value: 123.8, valExpect: uint32(123), errExpect: false},
		{tag: "uint32 value as expected", value: 12.3, valExpect: uint32(12), errExpect: false},
		{tag: "invalid uint32, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Uint32()
		if err != nil && !tc.errExpect {
			t.Error("uint32:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestUint64(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect uint64
		errExpect bool
	}{
		{tag: "uint64 value as expected", value: 123.8, valExpect: uint64(123), errExpect: false},
		{tag: "uint64 value as expected", value: 12.3, valExpect: uint64(12), errExpect: false},
		{tag: "invalid uint64, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Uint64()
		if err != nil && !tc.errExpect {
			t.Error("uint64:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestFloat32(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect float32
		errExpect bool
	}{
		{tag: "float32 value as expected", value: 123.8, valExpect: float32(123.8), errExpect: false},
		{tag: "float32 value as expected", value: 12.3, valExpect: float32(12.3), errExpect: false},
		{tag: "invalid float32, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Float32()
		if err != nil && !tc.errExpect {
			t.Error("float32:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestFloat64(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect float64
		errExpect bool
	}{
		{tag: "float64 value as expected", value: 123.8, valExpect: float64(123.8), errExpect: false},
		{tag: "float64 value as expected", value: 12.3, valExpect: float64(12.3), errExpect: true},
		{tag: "invalid float64, error expected", value: "123", valExpect: 0, errExpect: true},
	}

	for _, tc := range testCases {
		v, err := NewResult(tc.value).Float64()
		if err != nil && !tc.errExpect {
			t.Error("float64:", err)
		}
		if v != tc.valExpect && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, v)
		}
	}
}

func TestBoolSlice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []bool
		errExpect bool
	}{
		{tag: "boolSlice value as expected", value: []interface{}{true, false}, valExpect: []bool{true, false}, errExpect: false},
		{tag: "boolSlice value as expected", value: []interface{}{false, true, true}, valExpect: []bool{false, true, true}, errExpect: false},
		{tag: "invalid boolSlice, error expected", value: []interface{}{1, 3}, valExpect: []bool{}, errExpect: false},
		{tag: "invalid boolSlice, error expected", value: []int{1, 3}, valExpect: []bool{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).BoolSlice()
		if err != nil && !tc.errExpect {
			t.Error("boolSlice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestTimelice(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	t1, err1 := time.Parse(layout, "2014-11-12T11:45:26.371Z")
	if err1 != nil {
		t.Error("failed to parse time1:", err1)
	}
	t2, err2 := time.Parse(layout, "2019-11-12T11:45:26.371Z")
	if err2 != nil {
		t.Error("failed to parse time2:", err2)
	}
	testCases := []struct {
		tag        string
		value      interface{}
		timeLayout string
		valExpect  []time.Time
		errExpect  bool
	}{
		{tag: "timeSlice value as expected", value: []interface{}{"2014-11-12T11:45:26.371Z", "2019-11-12T11:45:26.371Z"}, timeLayout: layout, valExpect: []time.Time{t1, t2}, errExpect: false},
		{tag: "invalid timeSlice layout, error expected", value: []interface{}{"2014-11-12T11:45:26.371Z", "2019-11-12T11:45:26.371Z"}, timeLayout: "invalid layout", valExpect: []time.Time{}, errExpect: true},
		{tag: "invalid timeSlice, error expected", value: []int{1, 3}, timeLayout: layout, valExpect: []time.Time{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).TimeSlice(tc.timeLayout)
		if err != nil && !tc.errExpect {
			t.Error("timeSlice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}

}

func TestDurationlice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []time.Duration
		errExpect bool
	}{
		{tag: "durationSlice value as expected", value: []interface{}{"1s", "1m"}, valExpect: []time.Duration{1 * time.Second, 1 * time.Minute}, errExpect: false},
		{tag: "durationSlice value as expected", value: []interface{}{"1", "2"}, valExpect: []time.Duration{1 * time.Nanosecond, 2 * time.Nanosecond}, errExpect: false},
		{tag: "durationSlice value as expected", value: []interface{}{float64(2), float64(3)}, valExpect: []time.Duration{2 * time.Nanosecond, 3 * time.Nanosecond}, errExpect: false},
		{tag: "invalid durationSlice, error expected", value: []interface{}{"invalid duration 1", "invalid duration 2"}, valExpect: []time.Duration{}, errExpect: true},
		{tag: "invalid durationSlice, error expected", value: []float64{3, 5}, valExpect: []time.Duration{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).DurationSlice()
		if err != nil && !tc.errExpect {
			t.Error("durationSlice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestStringSlice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []string
		errExpect bool
	}{
		{tag: "stringSlice value as expected", value: []interface{}{"hello", "world"}, valExpect: []string{"hello", "world"}, errExpect: false},
		{tag: "stringSlice value as expected", value: []interface{}{"tom", "jerry"}, valExpect: []string{"tom", "jerry"}, errExpect: false},
		{tag: "invalid stringSlice, error expected", value: []interface{}{1, 3}, valExpect: []string{}, errExpect: false},
		{tag: "invalid stringSlice, error expected", value: []int{1, 3}, valExpect: []string{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).StringSlice()
		if err != nil && !tc.errExpect {
			t.Error("stringSlice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestIntSlice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []int
		errExpect bool
	}{
		{tag: "intSlice value as expected", value: []interface{}{132.1, 12.99}, valExpect: []int{132, 12}, errExpect: false},
		{tag: "intSlice value as expected", value: []interface{}{float64(131), float64(12)}, valExpect: []int{131, 12}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid intSlice, error expected", value: []interface{}{1, 3}, valExpect: []int{}, errExpect: false},
		{tag: "invalid intSlice, error expected", value: []int{1, 3}, valExpect: []int{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).IntSlice()
		if err != nil && !tc.errExpect {
			t.Error("intSlice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestInt8Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []int8
		errExpect bool
	}{
		{tag: "int8Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []int8{3, 12}, errExpect: false},
		{tag: "int8Slice value as expected", value: []interface{}{float64(11), float64(12)}, valExpect: []int8{11, 12}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid int8Slice, error expected", value: []interface{}{1, 3}, valExpect: []int8{}, errExpect: false},
		{tag: "invalid int8Slice, error expected", value: []int{1, 3}, valExpect: []int8{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Int8Slice()
		if err != nil && !tc.errExpect {
			t.Error("int8Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestInt16Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []int16
		errExpect bool
	}{
		{tag: "int16Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []int16{3, 12}, errExpect: false},
		{tag: "int16Slice value as expected", value: []interface{}{float64(11), float64(12)}, valExpect: []int16{11, 12}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid int16Slice, error expected", value: []interface{}{1, 3}, valExpect: []int16{}, errExpect: false},
		{tag: "invalid int16Slice, error expected", value: []int{1, 3}, valExpect: []int16{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Int16Slice()
		if err != nil && !tc.errExpect {
			t.Error("int16Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestInt32Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []int32
		errExpect bool
	}{
		{tag: "int32Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []int32{3, 12}, errExpect: false},
		{tag: "int32Slice value as expected", value: []interface{}{float64(131), float64(132)}, valExpect: []int32{131, 132}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid int32Slice, error expected", value: []interface{}{1, 3}, valExpect: []int32{}, errExpect: false},
		{tag: "invalid int32Slice, error expected", value: []int{1, 3}, valExpect: []int32{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Int32Slice()
		if err != nil && !tc.errExpect {
			t.Error("int32Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestInt64Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []int64
		errExpect bool
	}{
		{tag: "int64Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []int64{3, 12}, errExpect: false},
		{tag: "int64Slice value as expected", value: []interface{}{float64(131), float64(132)}, valExpect: []int64{131, 132}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid int64Slice, error expected", value: []interface{}{1, 3}, valExpect: []int64{}, errExpect: false},
		{tag: "invalid int64Slice, error expected", value: []int{1, 3}, valExpect: []int64{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Int64Slice()
		if err != nil && !tc.errExpect {
			t.Error("int64Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestUintSlice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []uint
		errExpect bool
	}{
		{tag: "uintSlice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []uint{3, 12}, errExpect: false},
		{tag: "uintSlice value as expected", value: []interface{}{float64(131), float64(132)}, valExpect: []uint{131, 132}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid uintSlice, error expected", value: []interface{}{1, 3}, valExpect: []uint{}, errExpect: false},
		{tag: "invalid uintSlice, error expected", value: []int{1, 3}, valExpect: []uint{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).UintSlice()
		if err != nil && !tc.errExpect {
			t.Error("uintSlice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestUint8Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []uint8
		errExpect bool
	}{
		{tag: "uint8Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []uint8{3, 12}, errExpect: false},
		{tag: "uint8Slice value as expected", value: []interface{}{float64(131), float64(132)}, valExpect: []uint8{131, 132}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid uint8Slice, error expected", value: []interface{}{1, 3}, valExpect: []uint8{}, errExpect: false},
		{tag: "invalid uint8Slice, error expected", value: []int{1, 3}, valExpect: []uint8{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Uint8Slice()
		if err != nil && !tc.errExpect {
			t.Error("uint8Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestUint16Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []uint16
		errExpect bool
	}{
		{tag: "uint16Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []uint16{3, 12}, errExpect: false},
		{tag: "uint16Slice value as expected", value: []interface{}{float64(131), float64(132)}, valExpect: []uint16{131, 132}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid uint16Slice, error expected", value: []interface{}{1, 3}, valExpect: []uint16{}, errExpect: false},
		{tag: "invalid uint16Slice, error expected", value: []int{1, 3}, valExpect: []uint16{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Uint16Slice()
		if err != nil && !tc.errExpect {
			t.Error("uint16Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestUint32Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []uint32
		errExpect bool
	}{
		{tag: "uint32Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []uint32{3, 12}, errExpect: false},
		{tag: "uint32Slice value as expected", value: []interface{}{float64(131), float64(132)}, valExpect: []uint32{131, 132}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid uint32Slice, error expected", value: []interface{}{1, 3}, valExpect: []uint32{}, errExpect: false},
		{tag: "invalid uint32Slice, error expected", value: []int{1, 3}, valExpect: []uint32{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Uint32Slice()
		if err != nil && !tc.errExpect {
			t.Error("uint32Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestUint64Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []uint64
		errExpect bool
	}{
		{tag: "uint64Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []uint64{3, 12}, errExpect: false},
		{tag: "uint64Slice value as expected", value: []interface{}{float64(131), float64(132)}, valExpect: []uint64{131, 132}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid uint64Slice, error expected", value: []interface{}{1, 3}, valExpect: []uint64{}, errExpect: false},
		{tag: "invalid uint64Slice, error expected", value: []int{1, 3}, valExpect: []uint64{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Uint64Slice()
		if err != nil && !tc.errExpect {
			t.Error("uint64Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestFloat32Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []float32
		errExpect bool
	}{
		{tag: "float32Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []float32{3.1, 12.99}, errExpect: false},
		{tag: "float32Slice value as expected", value: []interface{}{float64(131), float64(132)}, valExpect: []float32{131, 132}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid float32Slice, error expected", value: []interface{}{1, 3}, valExpect: []float32{}, errExpect: false},
		{tag: "invalid float32Slice, error expected", value: []int{1, 3}, valExpect: []float32{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Float32Slice()
		if err != nil && !tc.errExpect {
			t.Error("float32Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}

func TestFloat64Slice(t *testing.T) {
	testCases := []struct {
		tag       string
		value     interface{}
		valExpect []float64
		errExpect bool
	}{
		{tag: "float64Slice value as expected", value: []interface{}{3.1, 12.99}, valExpect: []float64{3.1, 12.99}, errExpect: false},
		{tag: "float64Slice value as expected", value: []interface{}{float64(131), float64(132)}, valExpect: []float64{131, 132}, errExpect: false}, // as golang decode number to float64
		{tag: "invalid float64Slice, error expected", value: []interface{}{1, 3}, valExpect: []float64{}, errExpect: false},
		{tag: "invalid float64Slice, error expected", value: []int{1, 3}, valExpect: []float64{}, errExpect: true},
	}

	for _, tc := range testCases {
		vv, err := NewResult(tc.value).Float64Slice()
		if err != nil && !tc.errExpect {
			t.Error("float64Slice:", err)
		}
		if !reflect.DeepEqual(vv, tc.valExpect) && !tc.errExpect {
			t.Errorf("tag: %s\nexpected: %v got %v", tc.tag, tc.valExpect, vv)
		}
	}
}
