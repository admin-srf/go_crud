package services

import "fmt"

func ErrParamIsRequired(param string, typ string) error {

	return fmt.Errorf("param: %s (type %s) is required", param, typ)
}

func MissingValue(param string) error {

	return fmt.Errorf("%s not found", param)
}

func CustomError(msg string) error {

	return fmt.Errorf(msg)
}
