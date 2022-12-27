package govalidate

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/rumis/govalidate/validator"
)

// NewFilter 构建新的FilterItem对象
func NewFilter(key string, rules []validator.Validator, errMsgCode ...string) validator.Filter {
	var err string
	if len(errMsgCode) > 0 {
		err = errMsgCode[0]
	}
	var code int32
	if len(errMsgCode) > 1 {
		eCode, _ := strconv.Atoi(errMsgCode[1])
		code = int32(eCode)
	}
	return validator.NewNormalFilter(key, rules, err, code)
}

// NewMultiLangFilter 多语言Filter对象
func NewMultiLangFilter(key string, rules []validator.Validator, errMsgCode ...string) validator.Filter {
	var err string
	if len(errMsgCode) > 0 {
		err = errMsgCode[0]
	}
	var code int32
	if len(errMsgCode) > 1 {
		eCode, _ := strconv.Atoi(errMsgCode[1])
		code = int32(eCode)
	}
	return validator.NewMultiLangFilter(key, rules, err, code)
}

// Validate 校验
func Validate(params map[string]interface{}, rules []validator.Filter) (map[string]interface{}, int32, error) {
	ctx := context.Background()
	return Validate1(ctx, params, rules)
}

// Validate1 校验
func Validate1(ctx context.Context, params map[string]interface{}, rules []validator.Filter) (map[string]interface{}, int32, error) {
	if len(rules) == 0 {
		return nil, 0, nil
	}
	vRes := make(map[string]interface{})
	for _, filter := range rules {
		key := filter.Key(ctx)
		paramVal, ok := params[key]
		if !ok {
			paramVal = nil
		}
		opts := &validator.ValidateOptions{
			Key:    key,
			Value:  paramVal,
			Params: params,
		}
		for _, fn := range filter.Rules(ctx) {
			res := fn(opts)
			if res.Stat(ctx) == validator.VS_BREAK {
				break
			}
			if res.Stat(ctx) == validator.VS_FAILUE {
				if res.ErrMsg(ctx) != "" {
					return vRes, filter.ErrCode(ctx), errors.New(res.ErrMsg(ctx))
				}
				if filter.ErrMsg(ctx) != "" {
					return vRes, filter.ErrCode(ctx), errors.New(filter.ErrMsg(ctx))
				}
				return vRes, filter.ErrCode(ctx), fmt.Errorf("field %s error", key)
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
