package govalidate

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/rumis/govalidate/validator"
)

// NewFilter 构建新的FilterItem对象
func NewFilter(key string, rules []validator.Validator, errMsgCode ...string) validator.FilterItem {
	item := validator.FilterItem{
		Key:   key,
		Rules: rules,
	}
	if len(errMsgCode) > 0 {
		item.ErrMsg = errors.New(errMsgCode[0])
	}
	if len(errMsgCode) > 1 {
		eCode, _ := strconv.Atoi(errMsgCode[1])
		item.ErrCode = int32(eCode)
	}
	return item
}

// Validate 校验
func Validate(params map[string]interface{}, rules []validator.FilterItem) (map[string]interface{}, int32, error) {
	if len(rules) == 0 {
		return nil, 0, nil
	}
	vRes := make(map[string]interface{})
	for _, filter := range rules {
		key := filter.Key
		paramVal, ok := params[key]
		if !ok {
			paramVal = nil
		}
		opts := &validator.ValidateOptions{
			Key:    key,
			Value:  paramVal,
			Params: params,
		}
		for _, fn := range filter.Rules {
			res := fn(opts)
			if res.Stat == validator.VS_BREAK {
				break
			}
			if res.Stat == validator.VS_FAILUE {
				if res.Emsg != nil {
					return vRes, filter.ErrCode, res.Emsg
				}
				if filter.ErrMsg != nil {
					return vRes, filter.ErrCode, filter.ErrMsg
				}
				return vRes, filter.ErrCode, fmt.Errorf("field %s error", key)
			}
		}
		// 记录校验结果
		if opts.Value != nil && opts.Key != "-" {
			vRes[opts.Key] = opts.Value
		}
		// 记录扩展数据
		if opts.Extend != nil {
			for ek, ev := range opts.Extend {
				vRes[ek] = ev
			}
		}
	}
	return vRes, 1, nil
}
