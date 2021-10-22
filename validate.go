package govalidate

import (
	"errors"
	"fmt"
	"strconv"
)

// ValidateStatus 规则校验结果
type ValidateStatus int8

const (
	VS_BREAK   ValidateStatus = 1
	VS_FAILUE  ValidateStatus = 2
	VS_SUCCESS ValidateStatus = 4
)

type ValidateOptions struct {
	Key    string
	Value  interface{}
	Params map[string]interface{}
	Extend map[string]interface{}
}

// Validator 规则
type Validator func(opts *ValidateOptions) ValidateResult

// ValidateResult 规则校验结果
type ValidateResult struct {
	Stat ValidateStatus
	Emsg error
}

// FilterItem 校验规则结构
type FilterItem struct {
	Rules   []Validator
	ErrMsg  error
	ErrCode int32
}

// Filter 构建新的FilterItem对象
func Filter(rules []Validator, errMsgCode ...string) FilterItem {
	item := FilterItem{
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
func Validate(params map[string]interface{}, rules map[string]FilterItem) (map[string]interface{}, int32, error) {
	if len(rules) == 0 {
		return nil, 0, nil
	}
	vRes := make(map[string]interface{})
	for key, filter := range rules {
		paramVal, ok := params[key]
		if !ok {
			paramVal = nil
		}
		opts := &ValidateOptions{
			Key:    key,
			Value:  paramVal,
			Params: params,
		}
		for _, fn := range filter.Rules {
			res := fn(opts)
			if res.Stat == VS_BREAK {
				break
			}
			if res.Stat == VS_FAILUE {
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

// Succ 规则校验通过
func Succ() ValidateResult {
	return ValidateResult{
		Stat: VS_SUCCESS,
		Emsg: nil,
	}
}

// Fail 规则校验失败
func Fail(emsg []string) ValidateResult {
	result := ValidateResult{
		Stat: VS_FAILUE,
	}
	if len(emsg) > 0 {
		result.Emsg = errors.New(emsg[0])
	}
	return result
}

// Break 中断后续校验流程
func Break() ValidateResult {
	return ValidateResult{
		Stat: VS_BREAK,
		Emsg: nil,
	}
}
