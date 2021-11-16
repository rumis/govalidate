package validator

import "errors"

// ValidateStatus 规则校验结果
type ValidateStatus int8

const (
	VS_BREAK   ValidateStatus = 1
	VS_FAILUE  ValidateStatus = 2
	VS_SUCCESS ValidateStatus = 4
)

// ValidateOptions 校验规则
type ValidateOptions struct {
	Key    string
	Value  interface{}
	Params map[string]interface{}
	Extend map[string]interface{}
}

// ValidateResult 规则校验结果
type ValidateResult struct {
	Stat ValidateStatus
	Emsg error
}

// FilterItem 校验规则结构
type FilterItem struct {
	Key     string
	Rules   []Validator
	ErrMsg  error
	ErrCode int32
}

// Validator 规则
type Validator func(opts *ValidateOptions) ValidateResult

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
