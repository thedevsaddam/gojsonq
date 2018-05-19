package gojsonq

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func isIndex(in string) bool {
	return strings.HasPrefix(in, "[") && strings.HasSuffix(in, "]")
}

func getIndex(in string) (int, error) {
	if !isIndex(in) {
		return -1, fmt.Errorf("invalid index")
	}
	is := strings.TrimLeft(in, "[")
	is = strings.TrimRight(is, "]")
	oint, err := strconv.Atoi(is)
	if err != nil {
		return -1, err
	}
	return oint, nil
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
	// switch v.(type) {
	// default:
	//  return fmt.Sprintf("%v", v)
	// case string:
	// 	return v.(string)
	// case float64:
	//
	// }
}

// toFloat64 convert interface{} value to float64 if value is numeric else return false
func toFloat64(v interface{}) (float64, bool) {
	var f float64
	flag := true
	// as Go convert the json Numeric value to float64
	switch v.(type) {
	case int:
		f = float64(v.(int))
	case int8:
		f = float64(v.(int8))
	case int16:
		f = float64(v.(int16))
	case int32:
		f = float64(v.(int32))
	case int64:
		f = float64(v.(int64))
	case float32:
		f = float64(v.(float32))
	case float64:
		f = v.(float64)
	default:
		flag = false
	}
	return f, flag
}

// sorter sor a list of intertace
func sorter(list []interface{}, asc bool) []interface{} {
	ss := []string{}
	ff := []float64{}
	result := []interface{}{}
	for _, v := range list {
		//sort elements for string
		if sv, ok := v.(string); ok {
			ss = append(ss, sv)
		}
		//sort elements for float64
		if fv, ok := v.(float64); ok {
			ff = append(ff, fv)
		}
	}

	if len(ss) > 0 {
		if asc {
			sort.Strings(ss)
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(ss)))
		}
		for _, v := range ss {
			result = append(result, v)
		}
	}
	if len(ff) > 0 {
		if asc {
			sort.Float64s(ff)
		} else {
			sort.Sort(sort.Reverse(sort.Float64Slice(ff)))
		}
		for _, v := range ff {
			result = append(result, v)
		}
	}
	return result
}

// func sortMap(mlist []map[string]interface{}, asc bool, stype string) []interface{} {
//
// 	return nil
// }

// sorter18 should use this func for go 1.8 build in future
// func sorter18(slice []interface{}) []interface{} {
// 	sort.SliceStable(slice, func(i int, j int) bool {
// 		if x, ok := slice[i].(string); ok {
// 			y, _ := slice[j].(string)
// 			return x < y
// 		}
// 		if x, ok := slice[i].(float64); ok {
// 			y, _ := slice[j].(float64)
// 			return x < y
// 		}
// 		return false
// 	})
// 	return slice
// }
