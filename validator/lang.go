package validator

// Localizer ..
type Localizer interface {
	Localize(id string) string
}

var localizerKey string = "i18n-localizer-key"

// InitLocalizerKey 初始化Localizer Key
func InitLocalizerKey(lkey string) {
	localizerKey = lkey
}

// GetLocalizerKey 获取Localizer Key
func GetLocalizerKey() string {
	return localizerKey
}
