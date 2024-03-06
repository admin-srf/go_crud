package handler

import (
	"fmt"
)

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", param, typ)
}
