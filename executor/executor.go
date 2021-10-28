package executor

import (
	"encoding/json"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

// IntExecutor 检查整形值是否符合要求
type IntExecutor func(val int) bool

// defaultIntExecutor 异常时使用，一定返回false
func defaultIntExecutor(val int) bool {
	return false
}

// StringExecutor 检查字符串值是否符合要求
type StringExecutor func(val string) bool

// defaultStringExecutor 异常时使用，一定返回false
func defaultStringExecutor(val string) bool {
	return false
}

// 检查字符串是否为邮箱地址
func Email(val string) bool {
	ok := emailReg.MatchString(val)
	return ok
}

// Url 检查字符串是否为有效的URL
func Url(val string) bool {
	if i := strings.Index(val, "#"); i > -1 {
		val = val[:i]
	}
	if len(val) == 0 {
		return false
	}
	url, err := url.ParseRequestURI(val)
	if err != nil || url.Scheme == "" {
		return false
	}
	return true
}

// Phone 是否为手机号
func Phone(val string) bool {
	return phoneReg.MatchString(val)
}

// Ipv4
func Ipv4(val string) bool {
	ip := net.ParseIP(val)
	if ip == nil || ip.To4() == nil {
		return false
	}
	return true
}

// Date 日期
func Date(val string) bool {
	_, err := time.Parse("2006-01-02", val)
	return err == nil
}

// Datetime 时间
func Datetime(val string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", val)
	return err == nil
}

// Length 字符串长度 [min,max]
func Length(min int, max int) StringExecutor {
	return func(val string) bool {
		l := utf8.RuneCountInString(val)
		if l < min || l > max {
			return false
		}
		return true
	}
}

// Between 数字值范围限制 [min,max]
func Between(min int, max int) IntExecutor {
	return func(val int) bool {
		if val < min || val > max {
			return false
		}
		return true
	}
}

// EnumInt 整数枚举
func EnumInt(enums []int) IntExecutor {
	return func(val int) bool {
		for _, v := range enums {
			if val == v {
				return true
			}
		}
		return false
	}
}

// EnumString 字符串枚举
func EnumString(enums []string) StringExecutor {
	return func(val string) bool {
		for _, v := range enums {
			if val == v {
				return true
			}
		}
		return false
	}
}

// DotInt 逗号分隔的字符串
func DotInt(val string) bool {
	return dotIntReg.MatchString(val)
}

// Regex 正则表达式
func Regex(reg string) StringExecutor {
	regExpr, err := regexp.Compile(reg)
	if err != nil {
		return defaultStringExecutor
	}
	return func(val string) bool {
		return regExpr.MatchString(val)
	}
}

func IsNil(val interface{}) bool {
	if val == nil {
		return true
	}
	switch v := val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		if v == 0 {
			return true
		}
	case float32, float64:
		if v == 0 {
			return true
		}
	case string:
		if len(v) == 0 {
			return true
		}
	case json.Number:
		vi, err := v.Int64()
		if err != nil {
			return true
		}
		if vi == 0 {
			return true
		}
	default:
		vo := reflect.ValueOf(val)
		if vo.IsNil() || vo.IsZero() || !vo.IsValid() {
			return true
		}
	}
	return false
}
