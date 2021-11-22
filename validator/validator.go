package validator

import (
	"strconv"
	"strings"

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
		if executor.IsNil(opts.Value) && len(defaultVal) != 0 {
			opts.Value = defaultVal[0]
			return Break()
		}
		if executor.IsNil(opts.Value) {
			return Break()
		}
		return Succ()
	}
}

// Int 参数为整形
func Int(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		v, ok := utils.GetIntValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		opts.Value = v
		return Succ()
	}
}

// Float 浮点数
func Float(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		v, ok := utils.GetFloatValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		opts.Value = v
		return Succ()
	}
}

// String 类型为字符串
func String(emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		str, ok := utils.GetStringValue(opts.Value)
		if !ok || len(str) == 0 {
			return Fail(emsg)
		}
		opts.Value = str
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
		opts.Value = str
		return Succ()
	}
}

// OmitEmpty 允许空
func OmitEmpty() Validator {
	return func(opts *ValidateOptions) ValidateResult {
		if executor.IsNil(opts.Value) {
			return Break()
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
		v, ok := utils.GetBooleanValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		opts.Value = v
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
		ok = executor.Length(min, max)(val)
		if !ok {
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
		ok = executor.Between(min, max)(val)
		if !ok {
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
		ok = executor.EnumInt(enums)(val)
		if !ok {
			return Fail(emsg)
		}
		opts.Value = val
		return Succ()
	}
}

// EnumString 枚举，值类型为字符串
func EnumString(enums []string, emsg ...string) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, ok := utils.GetStringValue(opts.Value)
		if !ok {
			return Fail(emsg)
		}
		ok = executor.EnumString(enums)(val)
		if !ok {
			return Fail(emsg)
		}
		opts.Value = val
		return Succ()
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

// Dotint64ToSlice 逗号分隔的ID字符串转为数组
func Dotint64ToSlice() Validator {
	return func(opts *ValidateOptions) ValidateResult {
		val, _ := utils.GetStringValue(opts.Value)
		vals := strings.Split(val, ",")
		if len(vals) > 0 {
			ids := make([]int64, len(vals))
			for idx, id := range vals {
				i, _ := strconv.ParseInt(id, 10, 64)
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
// 一个参数：可以是【错误信息】或者是【单个要素的校验条件】，校验条件可为单个或数组
// 两个参数：第一个参数一定为【错误信息】，第二个参数为【单个要素的校验条件】，校验条件可为单个或数组
func IntSlice(msgExecutor ...interface{}) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		errMsgs := make([]string, 0)
		execs := make([]executor.IntExecutor, 0)
		paramLen := len(msgExecutor)

		switch paramLen {
		case 0:
			// nothing to do
		case 1:
			switch p := msgExecutor[0].(type) {
			case string:
				errMsgs = append(errMsgs, p)
			case executor.IntExecutor:
				execs = append(execs, p)
			case []executor.IntExecutor:
				execs = append(execs, p...)
			}
		case 2:
			errMsg, ok := msgExecutor[0].(string)
			if !ok {
				return Fail(errMsgs)
			}
			errMsgs = append(errMsgs, errMsg)
			switch p := msgExecutor[1].(type) {
			case executor.IntExecutor:
				execs = append(execs, p)
			case []executor.IntExecutor:
				execs = append(execs, p...)
			}
		}
		vals, ok := utils.GetIntSlice(opts.Value)
		if !ok {
			return Fail(errMsgs)
		}
		if len(execs) > 0 {
			for _, exe := range execs {
				for _, val := range vals {
					ok = exe(val)
					if !ok {
						return Fail(errMsgs)
					}
				}
			}
		}
		opts.Value = vals
		return Succ()
	}
}

// StringSlice 字符串数组
// 一个参数：可以是【错误信息】或者是【单个要素的校验条件】，校验条件可为单个或数组
// 两个参数：第一个参数一定为【错误信息】，第二个参数为【单个要素的校验条件】，校验条件可为单个或数组
func StringSlice(msgExecutor ...interface{}) Validator {
	return func(opts *ValidateOptions) ValidateResult {
		errMsgs := make([]string, 0)
		execs := make([]executor.StringExecutor, 0)
		paramLen := len(msgExecutor)

		switch paramLen {
		case 0:
			// nothing to do
		case 1:
			switch p := msgExecutor[0].(type) {
			case string:
				errMsgs = append(errMsgs, p)
			case executor.StringExecutor:
				execs = append(execs, p)
			case []executor.StringExecutor:
				execs = append(execs, p...)
			}
		case 2:
			errMsg, ok := msgExecutor[0].(string)
			if !ok {
				return Fail(errMsgs)
			}
			errMsgs = append(errMsgs, errMsg)
			switch p := msgExecutor[1].(type) {
			case executor.StringExecutor:
				execs = append(execs, p)
			case []executor.StringExecutor:
				execs = append(execs, p...)
			}
		}
		vals, ok := utils.GetStringSlice(opts.Value)
		if !ok {
			return Fail(errMsgs)
		}
		if len(execs) > 0 {
			for _, exe := range execs {
				for _, val := range vals {
					ok = exe(val)
					if !ok {
						return Fail(errMsgs)
					}
				}
			}
		}
		opts.Value = vals
		return Succ()
	}
}
