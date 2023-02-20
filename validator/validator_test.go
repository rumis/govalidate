package validator

import (
	"fmt"
	"testing"
)

func TestXSS(t *testing.T) {

	opt := ValidateOptions{
		Value: "<Iframe srcdoc=<svg/o&#x6Eload&equals;alert&lpar;l)&gt;>",
	}
	opt1 := ValidateOptions{
		Value: "<script>小小</script>",
	}
	fn := XSS()

	fn(&opt)
	fn(&opt1)
	fmt.Println(opt)
	fmt.Println(opt1)
}
