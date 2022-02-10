package utils

import (
	"encoding/json"
	"strconv"
)

// GetIntValFromMap map中读取int值
func GetIntValFromMap(key string, vals map[string]interface{}) (int, bool) {
	iv, ok := vals[key]
	if !ok {
		return 0, false
	}
	return GetIntValue(iv)
}

// GetStringValFromMap map中读取string值
func GetStringValFromMap(key string, vals map[string]interface{}) (string, bool) {
	sv, ok := vals[key]
	if !ok {
		return "", false
	}
	return GetStringValue(sv)
}

// GetIntValue 转为整数
func GetIntValue(val interface{}) (int, bool) {
	switch v := val.(type) {
	case uint8:
		return int(v), true
	case uint16:
		return int(v), true
	case uint32:
		return int(v), true
	case uint64:
		return int(v), true
	case uint:
		return int(v), true
	case int8:
		return int(v), true
	case int16:
		return int(v), true
	case int32:
		return int(v), true
	case int64:
		return int(v), true
	case int:
		return int(v), true
	case string:
		if v == "" {
			return 0, true
		}
		vint, err := strconv.Atoi(v)
		if err != nil {
			return 0, false
		}
		return vint, true
	case json.Number:
		v64, err := v.Int64()
		if err != nil {
			return 0, false
		}
		return int(v64), true
	}
	return 0, false
}

// GetFloatValue 转为浮点
func GetFloatValue(val interface{}) (float64, bool) {
	switch v := val.(type) {
	case uint8:
		return float64(v), true
	case uint16:
		return float64(v), true
	case uint32:
		return float64(v), true
	case uint64:
		return float64(v), true
	case uint:
		return float64(v), true
	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	case int:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return v, true
	case string:
		vf, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, false
		}
		return vf, true
	case json.Number:
		v64, err := v.Float64()
		if err != nil {
			return 0, false
		}
		return v64, true
	}
	return 0, false
}

// GetStringValue 获取字符串
func GetStringValue(val interface{}) (string, bool) {
	switch v := val.(type) {
	case uint8:
		return strconv.FormatUint(uint64(v), 10), true
	case uint16:
		return strconv.FormatUint(uint64(v), 10), true
	case uint32:
		return strconv.FormatUint(uint64(v), 10), true
	case uint64:
		return strconv.FormatUint(uint64(v), 10), true
	case uint:
		return strconv.FormatUint(uint64(v), 10), true
	case int8:
		return strconv.FormatInt(int64(v), 10), true
	case int16:
		return strconv.FormatInt(int64(v), 10), true
	case int32:
		return strconv.FormatInt(int64(v), 10), true
	case int64:
		return strconv.FormatInt(v, 10), true
	case int:
		return strconv.FormatInt(int64(v), 10), true
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), true
	case float64:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), true
	case string:
		return v, true
	case json.Number:
		return v.String(), true
	}
	return "", false
}

// GetBooleanValue 获取布尔值
func GetBooleanValue(val interface{}) (bool, bool) {
	switch v := val.(type) {
	case bool:
		return v, true
	case string:
		vbool, err := strconv.ParseBool(v)
		if err != nil {
			return false, false
		}
		return vbool, true
	case json.Number:
		vbool, err := strconv.ParseBool(v.String())
		if err != nil {
			return false, false
		}
		return vbool, true
	}
	return false, false
}

// GetIntSlice 获取整形数组
func GetIntSlice(val interface{}) ([]int, bool) {
	switch vSlice := val.(type) {
	case []int:
		return vSlice, true
	case []string:
		res := make([]int, len(vSlice))
		for k, v := range vSlice {
			vi, err := strconv.Atoi(v)
			if err != nil {
				return []int{}, false
			}
			res[k] = int(vi)
		}
		return res, true
	case []json.Number:
		res := make([]int, len(vSlice))
		for k, v := range vSlice {
			vi, err := v.Int64()
			if err != nil {
				return []int{}, false
			}
			res[k] = int(vi)
		}
		return res, true
	case []interface{}:
		res := make([]int, len(vSlice))
		for k, v := range vSlice {
			vi, ok := GetIntValue(v)
			if !ok {
				return []int{}, false
			}
			res[k] = int(vi)
		}
		return res, true
	}
	return []int{}, false
}

// GetStringSlice 获取字符串数组
func GetStringSlice(val interface{}) ([]string, bool) {
	switch vSlice := val.(type) {
	case []int:
		res := make([]string, len(vSlice))
		for k, v := range vSlice {
			vs := strconv.FormatInt(int64(v), 10)
			res[k] = vs
		}
		return res, true
	case []string:
		return vSlice, true
	case []json.Number:
		res := make([]string, len(vSlice))
		for k, v := range vSlice {
			res[k] = v.String()
		}
		return res, true
	case []interface{}:
		res := make([]string, len(vSlice))
		for k, v := range vSlice {
			vi, ok := GetStringValue(v)
			if !ok {
				return []string{}, false
			}
			res[k] = vi
		}
		return res, true
	}
	return []string{}, false
}
