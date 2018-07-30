package gojsonq

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	signEq           = "="
	signEqEng        = "eq"
	signNotEq        = "!="
	signNotEqEng     = "neq"
	signNotEqAnother = "<>"
	signGt           = ">"
	signGtEng        = "gt"
	signLt           = "<"
	signLtEng        = "lt"
	signGtE          = ">="
	signGtEEng       = "gte"
	signLtE          = "<="
	signLtEEng       = "lte"

	signStrictContains = "strictContains"
	signContains       = "contains"
	signEndsWith       = "endsWith"
	signStartsWith     = "startsWith"
	signIn             = "in"
	signNotIn          = "notIn"

	signLenEq    = "leneq"
	signLenNotEq = "lenneq"
	signLenGt    = "lengt"
	signLenGte   = "lengte"
	signLenLt    = "lenlt"
	signLenLte   = "lenlte"
)

// QueryFunc describes a conditional function which perform comparison
type QueryFunc func(x, y interface{}) (bool, error)

// eq checks whether x, y are deeply eq
func eq(x, y interface{}) (bool, error) {
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		y = fv
	}
	return reflect.DeepEqual(x, y), nil
}

// neq checks whether x, y are deeply not equal
func neq(x, y interface{}) (bool, error) {
	b, err := eq(x, y)
	return !b, err
}

// gt checks whether x is greather than y
func gt(x, y interface{}) (bool, error) {
	xv, ok := x.(float64)
	if !ok {
		return false, fmt.Errorf("%v must be numeric", x)
	}
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		return xv > fv, nil
	}
	return false, nil
}

// lt checks whether x is less than y
func lt(x, y interface{}) (bool, error) {
	xv, ok := x.(float64)
	if !ok {
		return false, fmt.Errorf("%v must be numeric", x)
	}
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		return xv < fv, nil
	}
	return false, nil
}

// gte checks whether x is greater than or equal to y
func gte(x, y interface{}) (bool, error) {
	xv, ok := x.(float64)
	if !ok {
		return false, fmt.Errorf("%v must be numeric", x)
	}
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		return xv >= fv, nil
	}
	return false, nil
}

// lte checks whether x is less than or equal to y
func lte(x, y interface{}) (bool, error) {
	xv, ok := x.(float64)
	if !ok {
		return false, fmt.Errorf("%v must be numeric", x)
	}
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		return xv <= fv, nil
	}
	return false, nil
}

// strStrictContains checks if x contains y
// This is case sensitive search
func strStrictContains(x, y interface{}) (bool, error) {
	xv, okX := x.(string)
	if !okX {
		return false, fmt.Errorf("%v must be string", x)
	}
	yv, okY := y.(string)
	if !okY {
		return false, fmt.Errorf("%v must be string", y)
	}
	return strings.Contains(xv, yv), nil
}

// strContains checks if x contains y
// This is case insensitive search
func strContains(x, y interface{}) (bool, error) {
	xv, okX := x.(string)
	if !okX {
		return false, fmt.Errorf("%v must be string", x)
	}
	yv, okY := y.(string)
	if !okY {
		return false, fmt.Errorf("%v must be string", y)
	}
	return strings.Contains(strings.ToLower(xv), strings.ToLower(yv)), nil
}

// strStartsWith checks if x starts with y
func strStartsWith(x, y interface{}) (bool, error) {
	xv, okX := x.(string)
	if !okX {
		return false, fmt.Errorf("%v must be string", x)
	}
	yv, okY := y.(string)
	if !okY {
		return false, fmt.Errorf("%v must be string", y)
	}
	return strings.HasPrefix(xv, yv), nil
}

// strEndsWith checks if x ends with y
func strEndsWith(x, y interface{}) (bool, error) {
	xv, okX := x.(string)
	if !okX {
		return false, fmt.Errorf("%v must be string", x)
	}
	yv, okY := y.(string)
	if !okY {
		return false, fmt.Errorf("%v must be string", y)
	}
	return strings.HasSuffix(xv, yv), nil
}

// in checks if x exists in y e.g: in("id", []int{1,3,5,8})
func in(x, y interface{}) (bool, error) {
	if yv, ok := y.([]string); ok {
		for _, v := range yv {
			if ok, _ := eq(x, v); ok {
				return true, nil
			}
		}
	}
	if yv, ok := y.([]int); ok {
		for _, v := range yv {
			if ok, _ := eq(x, v); ok {
				return true, nil
			}
		}
	}
	if yv, ok := y.([]float64); ok {
		for _, v := range yv {
			if ok, _ := eq(x, v); ok {
				return true, nil
			}
		}
	}
	return false, nil
}

// notIn checks if x doesn't exists in y e.g: in("id", []int{1,3,5,8})
func notIn(x, y interface{}) (bool, error) {
	b, err := in(x, y)
	return !b, err
}

// lenEq checks if the string/array/list value is equal
func lenEq(x, y interface{}) (bool, error) {
	yv, ok := y.(int)
	if !ok {
		return false, fmt.Errorf("%v must be integer", y)
	}
	xv, err := length(x)
	if err != nil {
		return false, err
	}

	return xv == yv, nil
}

// lenNotEq checks if the string/array/list value is not equal
func lenNotEq(x, y interface{}) (bool, error) {
	yv, ok := y.(int)
	if !ok {
		return false, fmt.Errorf("%v must be integer", y)
	}
	xv, err := length(x)
	if err != nil {
		return false, err
	}

	return xv != yv, nil
}

// lenGt checks if the string/array/list value is greater
func lenGt(x, y interface{}) (bool, error) {
	yv, ok := y.(int)
	if !ok {
		return false, fmt.Errorf("%v must be integer", y)
	}
	xv, err := length(x)
	if err != nil {
		return false, err
	}

	return xv > yv, nil
}

// lenLt checks if the string/array/list value is less
func lenLt(x, y interface{}) (bool, error) {
	yv, ok := y.(int)
	if !ok {
		return false, fmt.Errorf("%v must be integer", y)
	}
	xv, err := length(x)
	if err != nil {
		return false, err
	}

	return xv < yv, nil
}

// lenGte checks if the string/array/list value is greater than equal
func lenGte(x, y interface{}) (bool, error) {
	yv, ok := y.(int)
	if !ok {
		return false, fmt.Errorf("%v must be integer", y)
	}
	xv, err := length(x)
	if err != nil {
		return false, err
	}

	return xv >= yv, nil
}

// lenLte checks if the string/array/list value is less than equal
func lenLte(x, y interface{}) (bool, error) {
	yv, ok := y.(int)
	if !ok {
		return false, fmt.Errorf("%v must be integer", y)
	}
	xv, err := length(x)
	if err != nil {
		return false, err
	}

	return xv <= yv, nil
}

func loadDefaultQueryMap() map[string]QueryFunc {
	// queryMap contains the registered conditional functions
	var queryMap = make(map[string]QueryFunc)

	queryMap[signEq] = eq
	queryMap[signEqEng] = eq

	queryMap[signNotEq] = neq
	queryMap[signNotEqEng] = neq
	queryMap[signNotEqAnother] = neq // also an alias of not equal

	queryMap[signGt] = gt
	queryMap[signGtEng] = gt

	queryMap[signLt] = lt
	queryMap[signLtEng] = lt

	queryMap[signGtE] = gte
	queryMap[signGtEEng] = gte

	queryMap[signLtE] = lte
	queryMap[signLtEEng] = lte

	queryMap[signStrictContains] = strStrictContains
	queryMap[signContains] = strContains
	queryMap[signStartsWith] = strStartsWith
	queryMap[signEndsWith] = strEndsWith

	queryMap[signIn] = in
	queryMap[signNotIn] = notIn

	queryMap[signLenEq] = lenEq
	queryMap[signLenNotEq] = lenNotEq

	queryMap[signLenGt] = lenGt
	queryMap[signLenGte] = lenGte

	queryMap[signLenLt] = lenLt
	queryMap[signLenLte] = lenLte

	return queryMap
}
