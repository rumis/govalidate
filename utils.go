package govalidate

// map中读取int值
func getIntValFromMap(key string, vals map[string]interface{}) (int, bool) {
	iv, ok := vals[key]
	if !ok {
		return 0, false
	}
	v, ok := iv.(int)
	if !ok {
		return 0, false
	}
	return v, true
}
