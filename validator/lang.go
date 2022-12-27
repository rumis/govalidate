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
