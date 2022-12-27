package validator

import (
	"context"
)

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
type ValidateResult interface {
	Stat(ctx context.Context) ValidateStatus
	ErrMsg(ctx context.Context) string
}

// ValidateResult 规则校验结果
type NormalValidateResult struct {
	Status ValidateStatus
	Emsg   string
}

// Stat 校验结果
func (r NormalValidateResult) Stat(ctx context.Context) ValidateStatus {
	return r.Status
}

// ErrMsg 校验错误
func (r NormalValidateResult) ErrMsg(ctx context.Context) string {
	return r.Emsg
}

// MultiLangValidateResult 多语言
type MultiLangValidateResult struct {
	Status ValidateStatus
	Emsg   string
}

// Stat 校验结果
func (r MultiLangValidateResult) Stat(ctx context.Context) ValidateStatus {
	return r.Status
}

// ErrMsg 校验错误
func (r MultiLangValidateResult) ErrMsg(ctx context.Context) string {
	if r.Emsg == "" {
		return ""
	}
	ilocalizer := ctx.Value(localizerKey)
	localizer, ok := ilocalizer.(Localizer)
	if !ok {
		return r.Emsg
	}
	msg := localizer.Localize(r.Emsg)
	return msg
}

// Filter 参数校验规则器
type Filter interface {
	Key(ctx context.Context) string
	ErrMsg(ctx context.Context) string
	ErrCode(ctx context.Context) int32
	Rules(ctx context.Context) []Validator
}

// NormalFilter 校验规则结构
type NormalFilter struct {
	key     string
	rules   []Validator
	errMsg  string
	errCode int32
}

// NewNormalFilter
func NewNormalFilter(k string, rules []Validator, err string, code int32) NormalFilter {
	return NormalFilter{
		key:     k,
		rules:   rules,
		errMsg:  err,
		errCode: code,
	}
}

// Key 获取参数KEY
func (f NormalFilter) Key(ctx context.Context) string {
	return f.key
}

// ErrMsg 获取错误信息
func (f NormalFilter) ErrMsg(ctx context.Context) string {
	return f.errMsg
}

// ErrCode 获取错误码
func (f NormalFilter) ErrCode(ctx context.Context) int32 {
	return f.errCode
}

// GetRules 获取所有规则
func (f NormalFilter) Rules(ctx context.Context) []Validator {
	return f.rules
}

// MultiLangFilter 支持多语言
type MultiLangFilter struct {
	key     string
	rules   []Validator
	errMsg  string
	errCode int32
}

// NewMultiLangFilter
func NewMultiLangFilter(k string, rules []Validator, err string, code int32) MultiLangFilter {
	return MultiLangFilter{
		key:     k,
		rules:   rules,
		errMsg:  err,
		errCode: code,
	}
}

// Key 获取参数KEY
func (f MultiLangFilter) Key(ctx context.Context) string {
	return f.key
}

// ErrMsg 获取错误信息
func (f MultiLangFilter) ErrMsg(ctx context.Context) string {
	if f.errMsg == "" {
		return ""
	}
	ilocalizer := ctx.Value(localizerKey)
	localizer, ok := ilocalizer.(Localizer)
	if !ok {
		return f.errMsg
	}
	msg := localizer.Localize(f.errMsg)
	return msg
}

// ErrCode 获取错误码
func (f MultiLangFilter) ErrCode(ctx context.Context) int32 {
	return f.errCode
}

// GetRules 获取所有规则
func (f MultiLangFilter) Rules(ctx context.Context) []Validator {
	return f.rules
}

// Validator 规则
type Validator func(opts *ValidateOptions) ValidateResult

// Succ 规则校验通过
func Succ() ValidateResult {
	return NormalValidateResult{
		Status: VS_SUCCESS,
		Emsg:   "",
	}
}

// Fail 规则校验失败
func Fail(emsg []string) ValidateResult {
	result := NormalValidateResult{
		Status: VS_FAILUE,
	}
	if len(emsg) > 0 {
		result.Emsg = emsg[0]
	}
	return result
}

// Fail 规则校验失败
func FailMultiLang(emsg []string) ValidateResult {
	result := MultiLangValidateResult{
		Status: VS_FAILUE,
	}
	if len(emsg) > 0 {
		result.Emsg = emsg[0]
	}
	return result
}

// Break 中断后续校验流程
func Break() ValidateResult {
	return NormalValidateResult{
		Status: VS_BREAK,
		Emsg:   "",
	}
}
