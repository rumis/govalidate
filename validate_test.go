package govalidate

import (
	"testing"
)

func TestValidate(t *testing.T) {

	// base
	params := map[string]interface{}{
		"age": 1,
	}
	rules := map[string]FilterItem{
		"age": Filter([]Validator{Int()}),
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
		"age": "1",
	}
	rules = map[string]FilterItem{
		"age": Filter([]Validator{Int()}, "错误信息", "10086"),
	}
	_, code, err := Validate(params, rules)
	if err == nil || err.Error() != "错误信息" {
		t.Error("error not find")
	}
	if code != 10086 {
		t.Errorf("error code is error: %d", code)
	}

	// required
	rules = map[string]FilterItem{
		"name": Filter([]Validator{Required("参数X必须")}),
	}
	_, _, err = Validate(params, rules)
	if err == nil || err.Error() != "参数X必须" {
		t.Error(err)
	}

	// optional
	rules = map[string]FilterItem{
		"name":  Filter([]Validator{Optional("默认值")}),
		"grade": Filter([]Validator{Optional()}),
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
	rules = map[string]FilterItem{
		"name1": Filter([]Validator{Int()}),
		"name2": Filter([]Validator{OmitEmpty()}),
		"name3": Filter([]Validator{Float()}),
		"name4": Filter([]Validator{Float()}),
		"name5": Filter([]Validator{Boolean()}),
		"name6": Filter([]Validator{Required()}),
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
	rules = map[string]FilterItem{
		"e1": Filter([]Validator{Required(), Email()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	params = map[string]interface{}{
		"e1": "@tal.com",
	}
	rules = map[string]FilterItem{
		"e1": Filter([]Validator{Required(), Email()}),
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
	rules = map[string]FilterItem{
		"u1": Filter([]Validator{Optional(), OmitEmpty(), Url()}),
		"u2": Filter([]Validator{Optional(), OmitEmpty(), Url()}),
		"u3": Filter([]Validator{Required(), Url()}),
		"u4": Filter([]Validator{Required(), Url()}),
		"u5": Filter([]Validator{Required(), Url()}),
		"u6": Filter([]Validator{Required(), Url()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	// 手机号
	params = map[string]interface{}{
		"p1": "15810562936",
	}
	rules = map[string]FilterItem{
		"p1": Filter([]Validator{Optional(), OmitEmpty(), Phone()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	params = map[string]interface{}{
		"p2": "12810562936",
	}
	rules = map[string]FilterItem{
		"p2": Filter([]Validator{Optional(), OmitEmpty(), Phone()}),
	}
	_, _, err = Validate(params, rules)
	if err == nil {
		t.Error(err)
	}

	// ipv4地址
	params = map[string]interface{}{
		"ip1": "127.127.127.127",
	}
	rules = map[string]FilterItem{
		"ip1": Filter([]Validator{Required(), Ipv4()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	params = map[string]interface{}{
		"ip2": "127.333.1.1",
	}
	rules = map[string]FilterItem{
		"ip2": Filter([]Validator{Required(), Ipv4()}),
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
	rules = map[string]FilterItem{
		"d1": Filter([]Validator{Required(), Datetime()}),
		"d2": Filter([]Validator{Required(), Date()}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}

	params = map[string]interface{}{
		"d3": "2021-1-11 15:33:21",
	}
	rules = map[string]FilterItem{
		"d3": Filter([]Validator{Required(), Datetime()}),
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
	rules = map[string]FilterItem{
		"l1": Filter([]Validator{Required(), Length(4, 6)}),
		"r1": Filter([]Validator{Required(), Between(1, 100)}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}

	params = map[string]interface{}{
		"l3": "2021-1-11 15:33:21",
	}
	rules = map[string]FilterItem{
		"dl": Filter([]Validator{Required(), Length(1, 2)}),
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
	rules = map[string]FilterItem{
		"ei1": Filter([]Validator{Required(), EnumInt([]int{1, 2, 3, 4})}),
		"es1": Filter([]Validator{Required(), EnumString([]string{"man", "feman"})}),
	}
	_, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}

	// 逗号分隔的整型IDs
	params = map[string]interface{}{
		"di1": "1,2,3,4",
	}
	rules = map[string]FilterItem{
		"di1": Filter([]Validator{Required(), DotInt(), Maxdot(5), Dotint2Slice()}),
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
	rules = map[string]FilterItem{
		"r1": Filter([]Validator{Required(), Regex("^[0-9]*$")}),
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
	rules = map[string]FilterItem{
		"curpage": Filter([]Validator{Optional(1), Int()}),
		"perpage": Filter([]Validator{Optional(13), Int()}),
		"x":       Filter([]Validator{Paginate()}),
	}
	res, _, err = Validate(params, rules)
	if err != nil {
		t.Error(err)
	}
	cur, ok := getIntValFromMap("curpage", res)
	if !ok || cur != 2 {
		t.Error("curpage error")
	}
	per, ok := getIntValFromMap("perpage", res)
	if !ok || per != 14 {
		t.Error("perpage error")
	}
	offset, ok := getIntValFromMap("offset", res)
	if !ok || offset != 14 {
		t.Error("offset error")
	}

}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		params := map[string]interface{}{
			"name": "1",
		}
		rules := map[string]FilterItem{
			"name": Filter([]Validator{Int()}),
		}
		_, _, _ = Validate(params, rules)
	}
}
