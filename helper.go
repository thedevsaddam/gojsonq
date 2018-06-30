package gojsonq

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		i = -1 * i
	}
	return i
}

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
}

// toFloat64 converts interface{} value to float64 if value is numeric else return false
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

// sortList sorts a list of interfaces
func sortList(list []interface{}, asc bool) []interface{} {
	ss := []string{}
	ff := []float64{}
	result := []interface{}{}
	for _, v := range list {
		// sort elements for string
		if sv, ok := v.(string); ok {
			ss = append(ss, sv)
		}
		// sort elements for float64
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

type sortMap struct {
	data interface{}
	key  string
	desc bool
}

// Sort sorts the slice of maps
func (s *sortMap) Sort(data interface{}) {
	s.data = data
	sort.Sort(s)
}

// Len satisfies the sort.Interface
func (s *sortMap) Len() int {
	return reflect.ValueOf(s.data).Len()
}

// Swap satisfies the sort.Interface
func (s *sortMap) Swap(i, j int) {
	if i > j {
		i, j = j, i
	}
	list := reflect.ValueOf(s.data)
	tmp := list.Index(i).Interface()
	list.Index(i).Set(list.Index(j))
	list.Index(j).Set(reflect.ValueOf(tmp))
}

// TODO: need improvement
// Less satisfies the sort.Interface
// This will work for string/float64 only
func (s *sortMap) Less(i, j int) (res bool) {
	list := reflect.ValueOf(s.data)
	x := list.Index(i).Interface()
	y := list.Index(j).Interface()

	// compare nested values
	if strings.Contains(s.key, ".") {
		xv, _ := getNestedValue(x, s.key)
		yv, _ := getNestedValue(y, s.key)
		res = s.compare(xv, yv)
	}

	xv, okX := x.(map[string]interface{})
	if !okX {
		return
	}
	yv := y.(map[string]interface{})
	if mvx, ok := xv[s.key]; ok {
		mvy := yv[s.key]
		res = s.compare(mvx, mvy)
	}

	return
}

// compare compare two values
func (s *sortMap) compare(x, y interface{}) (res bool) {
	if mfv, ok := x.(float64); ok {
		if mvy, oky := y.(float64); oky {
			if s.desc {
				return mfv > mvy
			}
			res = mfv < mvy
		}
	}

	if mfv, ok := x.(string); ok {
		if mvy, oky := y.(string); oky {
			if s.desc {
				return mfv > mvy
			}
			res = mfv < mvy
		}
	}

	return
}

// getNestedValue fetch nested value from node
func getNestedValue(input interface{}, node string) (interface{}, error) {
	pp := strings.Split(node, ".")
	for _, n := range pp {
		if isIndex(n) {
			// find slice/array
			if arr, ok := input.([]interface{}); ok {
				indx, err := getIndex(n)
				if err != nil {
					return input, err
				}
				arrLen := len(arr)
				if arrLen == 0 ||
					indx > arrLen-1 {
					return empty, errors.New("empty array")
				}
				input = arr[indx]
			}
		} else {
			// find in map
			validNode := false
			if mp, ok := input.(map[string]interface{}); ok {
				input, ok = mp[n]
				validNode = ok
			}

			// find in group data
			if mp, ok := input.(map[string][]interface{}); ok {
				input, ok = mp[n]
				validNode = ok
			}

			if !validNode {
				return empty, fmt.Errorf("invalid node name %s", n)
			}
		}
	}

	return input, nil
}
