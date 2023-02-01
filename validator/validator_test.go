package validator

import (
	"fmt"
	"testing"
)

func TestXSS(t *testing.T) {

	opt := ValidateOptions{
		Value: "<Iframe srcdoc=<svg/o&#x6Eload&equals;alert&lpar;l)&gt;>",
	}
	fn := XSS()
	fn(&opt)
	fmt.Println(opt)
}
