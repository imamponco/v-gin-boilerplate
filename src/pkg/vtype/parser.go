package vtype

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseBoolean(val interface{}) (result bool, ok bool) {
	ok = true
	switch t := val.(type) {
	case bool:
		result = t
	case string:
		result = strings.ToLower(t) == "true"
	case int:
		result = t > 0
	case int8:
		result = t > 0
	case int16:
		result = t > 0
	case int32:
		result = t > 0
	case int64:
		result = t > 0
	case uint:
		result = t > 0
	case uint8:
		result = t > 0
	case uint16:
		result = t > 0
	case uint32:
		result = t > 0
	case uint64:
		result = t > 0
	default:
		// Failed, set ok to false
		ok = false
	}

	return result, ok
}

func ParseBooleanFallback(v interface{}, fallback bool) bool {
	result, ok := ParseBoolean(v)
	if !ok {
		return fallback
	}
	return result
}

func ParseString(val interface{}) (result string, ok bool) {
	// Init success
	ok = true
	switch t := val.(type) {
	case string:
		result = t
	case nil:
		ok = false
	default:
		result = fmt.Sprintf("%v", val)
	}

	return result, ok
}

func ParseStringFallback(val interface{}, fallback string) string {
	result, ok := ParseString(val)
	if !ok || result == "" {
		return fallback
	}

	return result
}

func ParseInt(input interface{}) (result int, ok bool) {
	ok = true
	switch v := input.(type) {
	case string:
		var err error
		result, err = strconv.Atoi(v)
		if err != nil {
			ok = false
		}
	case int:
		result = v
	default:
		ok = false
	}

	return result, ok
}

func ParseIntFallback(input interface{}, fallbackValue int) int {
	i, ok := ParseInt(input)
	if !ok {
		return fallbackValue
	}
	return i
}

func ParseInt64(input interface{}) (result int64, ok bool) {
	// Assume ok to true
	ok = true

	switch v := input.(type) {
	case string:
		var err error
		result, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			ok = false
		}
	case int64:
		result = v
	case float64:
		result = int64(v)
	default:
		// Failed, set ok to false
		ok = false
	}

	return result, ok
}

func ParseInt64Fallback(input interface{}, fallbackValue int64) int64 {
	i, ok := ParseInt64(input)
	if !ok {
		return fallbackValue
	}
	return i
}
