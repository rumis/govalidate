package validator

import (
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/rumis/govalidate/executor"
	"github.com/rumis/govalidate/utils"
)

// Required 参数必须
func Required(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		if opts.Value != nil {
			return Succ()
		}
		return Fail(emsg)
	}
}

// Optional 参数可选，可设置默认值
func Optional(defaultVal ...interface{}) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		if opts.Value == nil && len(defaultVal) != 0 {
			opts.Value = defaultVal[0]
		}
		return Succ()
	}
}

// Int 参数为整形
func Int(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		_, ok := utils.GetIntValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Float 浮点数
func Float(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		_, ok := utils.GetFloatValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// EmptyString 空字符串，跳过后续校验规则
func EmptyString() Validator {
	return func(opts *ValidateOptions) ValidateResult {
		str, ok := utils.GetStringValue(opts.Value)
		if ok && len(str) == 0 {
			return Break()
		}
		return Succ()
	}
}

// OmitEmpty 允许空
func OmitEmpty() Validator {
	return func(opts *ValidateOptions) ValidateResult {
		if opts.Value == nil {
			return Break()
		}
		switch val := opts.Value.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			if val == 0 {
				return Break()
			}
		case float32, float64:
			if val == 0 {
				return Break()
			}
		case string:
			if len(val) == 0 {
				return Break()
			}
		default:
			vo := reflect.ValueOf(val)
			if vo.IsNil() || vo.IsZero() || !vo.IsValid() {
				return Break()
			}
		}
		return Succ()
	}
}

// ResetKey 重置参数key值
func ResetKey(newKey string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		if newKey != "" {
			opts.Key = newKey
		}
		return Succ()
	}
}

// Boolean 布尔值
func Boolean(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		_, ok := utils.GetBooleanValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Email 邮件
func Email(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		ok = executor.Email(val)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Url URL链接
func Url(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		ok = executor.Url(val)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Phone 手机号码
func Phone(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		ok = executor.Phone(val)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Ipv4 ip地址，v4格式
func Ipv4(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		ok = executor.Ipv4(val)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Date 日期，格式： 2006-01-02
func Date(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		ok = executor.Date(val)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Datetime 时间，格式：2006-01-02 15:04:05
func Datetime(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		ok = executor.Datetime(val)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Length 字符串字符长度限制 [min,max]
func Length(min int, max int, emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		l := utf8.RuneCountInString(val)
		if l < min || l > max {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Between 数字值范围限制 [min,max]
func Between(min int, max int, emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetIntValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		if val < min || val > max {
			return Fail(emsg)
		}
		return Succ()
	}
}

// EnumInt 枚举，值类型为整形
func EnumInt(enums []int, emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetIntValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		for _, v := range enums {
			if val == v {
				return Succ()
			}
		}
		return Fail(emsg)
	}
}

// EnumString 枚举，值类型为字符串
func EnumString(enums []string, emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		for _, v := range enums {
			if val == v {
				return Succ()
			}
		}
		return Fail(emsg)
	}
}

// DotInt 英文逗号分隔的整数
func DotInt(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		ok = executor.DotInt(val)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Maxdot 逗号分隔的ID支持的最多ID个数
func Maxdot(max int, emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		dotCnt := strings.Count(val, ",")
		if dotCnt+1 > max {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Dotint2Slice 逗号分隔的ID字符串转为数组
// 忽略了错误处理，需在DotInt规则后使用
func Dotint2Slice() Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, _ := utils.GetStringValue(opts.Value)
		vals := strings.Split(val, ",")
		if len(vals) > 0 {
			ids := make([]int, len(vals))
			for idx, id := range vals {
				i, _ := strconv.Atoi(id)
				ids[idx] = i
			}
			opts.Value = ids
		}
		return Succ()
	}
}

// Regex 正则表达式
func Regex(reg string, emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		ok = executor.Regex(reg)(val)
		if !ok {
			return Fail(emsg)
		}
		return Succ()
	}
}

// Paginate 处理分页信息
// 页码会被计算为偏移量
// fields【curpage，perpage】
func Paginate(fields ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		perpageKey := "perpage"
		curpageKey := "curpage"
		if len(fields) > 0 {
			curpageKey = fields[0]
		}
		if len(fields) > 1 {
			perpageKey = fields[1]
		}
		curpage := 1  // 默认第一页
		perpage := 10 // 默认每页10数据
		// 优先在处理结果中解析数据
		if cur, ok := utils.GetIntValFromMap(curpageKey, opts.Params); ok {
			curpage = cur
		}
		if per, ok := utils.GetIntValFromMap(perpageKey, opts.Params); ok {
			perpage = per
		}
		if opts.Extend == nil {
			opts.Extend = make(map[string]interface{})
		}
		opts.Extend["offset"] = (curpage - 1) * perpage
		opts.Key = "-"
		return Succ()
	}
}

// IntSlice 整形数组
func IntSlice() Validator {

	return nil
}

func StringSlice() Validator {

	return nil
}
