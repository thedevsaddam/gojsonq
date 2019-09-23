package gojsonq

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const errMessage = "gojsonq: wrong method call for %v"

// NewResult return an instance of Result
func NewResult(v interface{}) *Result {
	return &Result{value: v}
}

// Result represent custom type
type Result struct {
	value interface{}
}

// Nil check the query has result or not
func (r *Result) Nil() bool {
	return r.value == nil
}

// Bool assert the result to boolean value
func (r *Result) Bool() (bool, error) {
	switch v := r.value.(type) {
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Time assert the result to time.Time
func (r *Result) Time(layout string) (time.Time, error) {
	switch v := r.value.(type) {
	case string:
		return time.Parse(layout, v)
	default:
		return time.Time{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Duration assert the result to time.Duration
func (r *Result) Duration() (time.Duration, error) {
	switch v := r.value.(type) {
	case float64:
		return time.Duration(v), nil
	case string:
		if strings.ContainsAny(v, "nsuµmh") {
			return time.ParseDuration(v)
		}
		return time.ParseDuration(v + "ns")
	default:
		return time.Duration(0), fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// String assert the result to String
func (r *Result) String() (string, error) {
	switch v := r.value.(type) {
	case string:
		return v, nil
	default:
		return "", fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Int assert the result to int
func (r *Result) Int() (int, error) {
	switch v := r.value.(type) {
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Int8 assert the result to int8
func (r *Result) Int8() (int8, error) {
	switch v := r.value.(type) {
	case float64:
		return int8(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Int16 assert the result to int16
func (r *Result) Int16() (int16, error) {
	switch v := r.value.(type) {
	case float64:
		return int16(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Int32 assert the result to int32
func (r *Result) Int32() (int32, error) {
	switch v := r.value.(type) {
	case float64:
		return int32(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Int64 assert the result to int64
func (r *Result) Int64() (int64, error) {
	switch v := r.value.(type) {
	case float64:
		return int64(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Uint assert the result to uint
func (r *Result) Uint() (uint, error) {
	switch v := r.value.(type) {
	case float64:
		return uint(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Uint8 assert the result to uint8
func (r *Result) Uint8() (uint8, error) {
	switch v := r.value.(type) {
	case float64:
		return uint8(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Uint16 assert the result to uint16
func (r *Result) Uint16() (uint16, error) {
	switch v := r.value.(type) {
	case float64:
		return uint16(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Uint32 assert the result to uint32
func (r *Result) Uint32() (uint32, error) {
	switch v := r.value.(type) {
	case float64:
		return uint32(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Uint64 assert the result to uint64
func (r *Result) Uint64() (uint64, error) {
	switch v := r.value.(type) {
	case float64:
		return uint64(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Float32 assert the result to float32
func (r *Result) Float32() (float32, error) {
	switch v := r.value.(type) {
	case float64:
		return float32(v), nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Float64 assert the result to 64
func (r *Result) Float64() (float64, error) {
	switch v := r.value.(type) {
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// === slcie ===

// BoolSlice assert the result to []bool
func (r *Result) BoolSlice() ([]bool, error) {
	switch v := r.value.(type) {
	case []interface{}:
		bb := []bool{}
		for _, si := range v {
			if s, ok := si.(bool); ok {
				bb = append(bb, s)
			}
		}
		return bb, nil
	default:
		return []bool{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// TimeSlice assert the result to []time.Time
func (r *Result) TimeSlice(layout string) ([]time.Time, error) {
	switch v := r.value.(type) {
	case []interface{}:
		tt := []time.Time{}
		for _, si := range v {
			if s, ok := si.(string); ok {
				ts, err := time.Parse(layout, s)
				if err != nil {
					return tt, err
				}
				tt = append(tt, ts)
			}
		}
		return tt, nil
	default:
		return []time.Time{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// DurationSlice assert the result to []time.Duration
func (r *Result) DurationSlice() ([]time.Duration, error) {
	switch v := r.value.(type) {
	case []interface{}:
		dd := []time.Duration{}
		for _, si := range v {
			if s, ok := si.(string); ok {
				var d time.Duration
				var err error
				if strings.ContainsAny(s, "nsuµmh") {
					d, err = time.ParseDuration(s)
				} else {
					d, err = time.ParseDuration(s + "ns")
				}
				if err != nil {
					return dd, err
				}
				dd = append(dd, d)
			}

			if v, ok := si.(float64); ok {
				dd = append(dd, time.Duration(v))
			}
		}
		return dd, nil
	default:
		return []time.Duration{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// StringSlice assert the result to []string
func (r *Result) StringSlice() ([]string, error) {
	switch v := r.value.(type) {
	case []interface{}:
		ss := []string{}
		for _, si := range v {
			if s, ok := si.(string); ok {
				ss = append(ss, s)
			}
		}
		return ss, nil
	default:
		return []string{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// IntSlice assert the result to []int
func (r *Result) IntSlice() ([]int, error) {
	switch v := r.value.(type) {
	case []interface{}:
		ii := []int{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				ii = append(ii, int(s))
			}
		}
		return ii, nil
	default:
		return []int{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Int8Slice assert the result to []int8
func (r *Result) Int8Slice() ([]int8, error) {
	switch v := r.value.(type) {
	case []interface{}:
		ii := []int8{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				ii = append(ii, int8(s))
			}
		}
		return ii, nil
	default:
		return []int8{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Int16Slice assert the result to []int16
func (r *Result) Int16Slice() ([]int16, error) {
	switch v := r.value.(type) {
	case []interface{}:
		ii := []int16{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				ii = append(ii, int16(s))
			}
		}
		return ii, nil
	default:
		return []int16{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Int32Slice assert the result to []int32
func (r *Result) Int32Slice() ([]int32, error) {
	switch v := r.value.(type) {
	case []interface{}:
		ii := []int32{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				ii = append(ii, int32(s))
			}
		}
		return ii, nil
	default:
		return []int32{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Int64Slice assert the result to []int64
func (r *Result) Int64Slice() ([]int64, error) {
	switch v := r.value.(type) {
	case []interface{}:
		ii := []int64{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				ii = append(ii, int64(s))
			}
		}
		return ii, nil
	default:
		return []int64{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// UintSlice assert the result to []uint
func (r *Result) UintSlice() ([]uint, error) {
	switch v := r.value.(type) {
	case []interface{}:
		uu := []uint{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				uu = append(uu, uint(s))
			}
		}
		return uu, nil
	default:
		return []uint{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Uint8Slice assert the result to []uint8
func (r *Result) Uint8Slice() ([]uint8, error) {
	switch v := r.value.(type) {
	case []interface{}:
		uu := []uint8{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				uu = append(uu, uint8(s))
			}
		}
		return uu, nil
	default:
		return []uint8{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Uint16Slice assert the result to []uint16
func (r *Result) Uint16Slice() ([]uint16, error) {
	switch v := r.value.(type) {
	case []interface{}:
		uu := []uint16{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				uu = append(uu, uint16(s))
			}
		}
		return uu, nil
	default:
		return []uint16{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Uint32Slice assert the result to []uint32
func (r *Result) Uint32Slice() ([]uint32, error) {
	switch v := r.value.(type) {
	case []interface{}:
		uu := []uint32{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				uu = append(uu, uint32(s))
			}
		}
		return uu, nil
	default:
		return []uint32{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Uint64Slice assert the result to []uint64
func (r *Result) Uint64Slice() ([]uint64, error) {
	switch v := r.value.(type) {
	case []interface{}:
		uu := []uint64{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				uu = append(uu, uint64(s))
			}
		}
		return uu, nil
	default:
		return []uint64{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Float32Slice assert the result to []float32
func (r *Result) Float32Slice() ([]float32, error) {
	switch v := r.value.(type) {
	case []interface{}:
		ff := []float32{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				ff = append(ff, float32(s))
			}
		}
		return ff, nil
	default:
		return []float32{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}

// Float64Slice assert the result to []float64
func (r *Result) Float64Slice() ([]float64, error) {
	switch v := r.value.(type) {
	case []interface{}:
		ff := []float64{}
		for _, si := range v {
			if s, ok := si.(float64); ok {
				ff = append(ff, s)
			}
		}
		return ff, nil
	default:
		return []float64{}, fmt.Errorf(errMessage, reflect.ValueOf(r.value).Kind())
	}
}
