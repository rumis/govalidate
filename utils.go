package govalidate

import (
	"encoding/json"
	"strconv"
)

// map中读取int值
func getIntValFromMap(key string, vals map[string]interface{}) (int, bool) {
	iv, ok := vals[key]
	if !ok {
		return 0, false
	}
	return getIntValue(iv)
}

// getIntValue 转为整数
func getIntValue(val interface{}) (int, bool) {
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

// getFloatValue 转为浮点
func getFloatValue(val interface{}) (float64, bool) {
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

// getStringValue 获取字符串
func getStringValue(val interface{}) (string, bool) {
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

// getBooleanValue 获取布尔值
func getBooleanValue(val interface{}) (bool, bool) {
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
