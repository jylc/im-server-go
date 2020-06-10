package telephone

import (
	"regexp"
)

//校验手机号
func IsValid(telephone string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	compile, err := regexp.Compile(regular)
	if err != nil {
		return false
	}
	if match := compile.MatchString(telephone); !match {
		return false
	}
	return true
}
