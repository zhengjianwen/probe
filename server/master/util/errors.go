package util

import "fmt"

func ParamNotFoundErr(param string) error {
	return fmt.Errorf("param %s not found", param)
}

func ParamLengthExceedErr(param string) error {
	return fmt.Errorf("param %s length exceeded", param)
}

func ParamInvalid(param string) error {
	return fmt.Errorf("param %s invalid", param)
}

func ParamLengthInvalid(param string, min, max uint16) error {
	return fmt.Errorf("param %s length must between %d and %d", param, min, max)
}

func ParamValueInvalid(param string, min, max int) error {
	return fmt.Errorf("param %s must between %d and %d", param, min, max)
}

func ParamNotInValidList(param string, l ...interface{}) error {
	return fmt.Errorf("param %s not in valid list %v", param, l)
}
