package govalidate

import (
	"testing"

	"github.com/rumis/govalidate/executor"
	"github.com/rumis/govalidate/utils"
	"github.com/rumis/govalidate/validator"
)

func TestValidate(t *testing.T) {

	// base
	params := map[string]interface{}{
		"age": 1,
	}
	rules := map[string]validator.FilterItem{
		"age": Filter([]validator.Validator{validator.Int()}),
	}
	res, _, err := Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	if _, ok := res["age"]; !ok {
		t.Error("param age not found")
	}

	// error code
	params = map[string]interface{}{
		"age": "s",
	}
	rules = map[string]validator.FilterItem{
		"age": Filter([]validator.Validator{validator.Int()}, "错误信息", "10086"),
	}
	_, code, err := Validate(params, rules)
	if err == nil || err.Error() != "错误信息" {
		t.Error("error not find")
	}
	if code != 10086 {
		t.Errorf("error code is error: %d", code)
	}

	// required
	rules = map[string]validator.FilterItem{
		"name": Filter([]validator.Validator{validator.Required("参数X必须")}),
	}
	_, _, err = Validate(params, rules)
	if err == nil || err.Error() != "参数X必须" {
		t.Error(err)
	}

	// optional
	rules = map[string]validator.FilterItem{
		"name":  Filter([]validator.Validator{validator.Optional("默认值")}),
		"grade": Filter([]validator.Validator{validator.Optional()}),
	}
	res, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	if dname, ok := res["name"]; !ok || dname != "默认值" {
		t.Error("optional default value error")
	}
	if _, ok := res["grade"]; ok {
		t.Error("optional error")
	}

	// 数据类型
	params = map[string]interface{}{
		"name1": 1,
		"name2": "1",
		"name3": 2,
		"name4": 2.3,
		"name5": "true",
		"name6": 2,
	}
	rules = map[string]validator.FilterItem{
		"name1": Filter([]validator.Validator{validator.Int()}),
		"name2": Filter([]validator.Validator{validator.OmitEmpty()}),
		"name3": Filter([]validator.Validator{validator.Float()}),
		"name4": Filter([]validator.Validator{validator.Float()}),
		"name5": Filter([]validator.Validator{validator.Boolean()}),
		"name6": Filter([]validator.Validator{validator.Required()}),
	}
	res, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	fi1, ok := res["name3"]
	if !ok {
		t.Error("float params not found")
	}
	f1, ok := fi1.(int)
	if !ok {
		t.Error("int can't as float")
	}
	if f1 != 2 {
		t.Error("float error")
	}

	fi2, ok := res["name4"]
	if !ok {
		t.Error("float params not found")
	}
	f2, ok := fi2.(float64)
	if !ok {
		t.Error("int can't as float")
	}
	if f2 != 2.3 {
		t.Error("float error")
	}

	// ------- 格式
	// 邮件
	params = map[string]interface{}{
		"e1": "liumurong1@tal.com",
	}
	rules = map[string]validator.FilterItem{
		"e1": Filter([]validator.Validator{validator.Required(), validator.Email()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	params = map[string]interface{}{
		"e1": "@tal.com",
	}
	rules = map[string]validator.FilterItem{
		"e1": Filter([]validator.Validator{validator.Required(), validator.Email()}),
	}
	_, _, err = Validate(params, rules)
	if err == nil {
		t.Error(err)
	}

	// URL
	params = map[string]interface{}{
		"u3": "https://baidu.com",
		"u4": "http://www.baidu.com",
		"u5": "https://www.baidu.com?x=3",
		"u6": "https://www.baidu.com#de",
	}
	rules = map[string]validator.FilterItem{
		"u1": Filter([]validator.Validator{validator.Optional(), validator.OmitEmpty(), validator.Url()}),
		"u2": Filter([]validator.Validator{validator.Optional(), validator.OmitEmpty(), validator.Url()}),
		"u3": Filter([]validator.Validator{validator.Required(), validator.Url()}),
		"u4": Filter([]validator.Validator{validator.Required(), validator.Url()}),
		"u5": Filter([]validator.Validator{validator.Required(), validator.Url()}),
		"u6": Filter([]validator.Validator{validator.Required(), validator.Url()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	// 手机号
	params = map[string]interface{}{
		"p1": "15810562936",
	}
	rules = map[string]validator.FilterItem{
		"p1": Filter([]validator.Validator{validator.Optional(), validator.OmitEmpty(), validator.Phone()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	params = map[string]interface{}{
		"p2": "12810562936",
	}
	rules = map[string]validator.FilterItem{
		"p2": Filter([]validator.Validator{validator.Optional(), validator.OmitEmpty(), validator.Phone()}),
	}
	_, _, err = Validate(params, rules)
	if err == nil {
		t.Error(err)
	}

	// ipv4地址
	params = map[string]interface{}{
		"ip1": "127.127.127.127",
	}
	rules = map[string]validator.FilterItem{
		"ip1": Filter([]validator.Validator{validator.Required(), validator.Ipv4()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	params = map[string]interface{}{
		"ip2": "127.333.1.1",
	}
	rules = map[string]validator.FilterItem{
		"ip2": Filter([]validator.Validator{validator.Required(), validator.Ipv4()}),
	}
	_, _, err = Validate(params, rules)
	if err == nil {
		t.Error(err)
	}

	// 时间，日期
	params = map[string]interface{}{
		"d1": "2021-10-11 15:33:21",
		"d2": "2021-10-11",
	}
	rules = map[string]validator.FilterItem{
		"d1": Filter([]validator.Validator{validator.Required(), validator.Datetime()}),
		"d2": Filter([]validator.Validator{validator.Required(), validator.Date()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}

	params = map[string]interface{}{
		"d3": "2021-1-11 15:33:21",
	}
	rules = map[string]validator.FilterItem{
		"d3": Filter([]validator.Validator{validator.Required(), validator.Datetime()}),
	}
	_, _, err = Validate(params, rules)
	if err == nil {
		t.Error(err)
	}

	// 范围
	params = map[string]interface{}{
		"l1": "字符长度5",
		"r1": 99,
	}
	rules = map[string]validator.FilterItem{
		"l1": Filter([]validator.Validator{validator.Required(), validator.Length(4, 6)}),
		"r1": Filter([]validator.Validator{validator.Required(), validator.Between(1, 100)}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}

	params = map[string]interface{}{
		"l3": "2021-1-11 15:33:21",
	}
	rules = map[string]validator.FilterItem{
		"dl": Filter([]validator.Validator{validator.Required(), validator.Length(1, 2)}),
	}
	_, _, err = Validate(params, rules)
	if err == nil {
		t.Error(err)
	}

	// 枚举
	params = map[string]interface{}{
		"ei1": 3,
		"es1": "man",
	}
	rules = map[string]validator.FilterItem{
		"ei1": Filter([]validator.Validator{validator.Required(), validator.EnumInt([]int{1, 2, 3, 4})}),
		"es1": Filter([]validator.Validator{validator.Required(), validator.EnumString([]string{"man", "feman"})}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}

	// 逗号分隔的整型IDs
	params = map[string]interface{}{
		"di1": "1,2,3,4",
	}
	rules = map[string]validator.FilterItem{
		"di1": Filter([]validator.Validator{validator.Required(), validator.DotInt(), validator.Maxdot(5), validator.Dotint2Slice()}),
	}
	res, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	idSlice, ok := res["di1"]
	if !ok {
		t.Error("return value error")
	}
	if ids, ok := idSlice.([]int); !ok || len(ids) != 4 {
		t.Error("dotint to slice error")
	}

	// 正则表达式
	params = map[string]interface{}{
		"r1": "034433332",
	}
	rules = map[string]validator.FilterItem{
		"r1": Filter([]validator.Validator{validator.Required(), validator.Regex("^[0-9]*$")}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}

	// 分页
	params = map[string]interface{}{
		"curpage": 2,
		"perpage": 14,
	}
	rules = map[string]validator.FilterItem{
		"curpage": Filter([]validator.Validator{validator.Optional(1), validator.Int()}),
		"perpage": Filter([]validator.Validator{validator.Optional(13), validator.Int()}),
		"x":       Filter([]validator.Validator{validator.Paginate()}),
	}
	res, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	cur, ok := utils.GetIntValFromMap("curpage", res)
	if !ok || cur != 2 {
		t.Error("curpage error")
	}
	per, ok := utils.GetIntValFromMap("perpage", res)
	if !ok || per != 14 {
		t.Error("perpage error")
	}
	offset, ok := utils.GetIntValFromMap("offset", res)
	if !ok || offset != 14 {
		t.Error("offset error")
	}

	// 切片
	params = map[string]interface{}{
		"s1": []string{"1", "2"},
	}
	rules = map[string]validator.FilterItem{
		"s1": Filter([]validator.Validator{validator.Required(), validator.IntSlice("切片异常", executor.EnumInt([]int{1, 2, 3, 4}))}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		params := map[string]interface{}{
			"name": "1",
		}
		rules := map[string]validator.FilterItem{
			"name": Filter([]validator.Validator{validator.Int()}),
		}
		_, _, _ = Validate(params, rules)
	}
}
