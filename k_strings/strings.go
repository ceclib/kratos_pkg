package k_strings

import (
	"errors"
	"strconv"
)

func Interface2String(inter interface{}) (string, error) {

	switch inter.(type) {
	case string:
		return inter.(string), nil
	case int:
		return strconv.Itoa(inter.(int)), nil
	case float64:
		return strconv.FormatFloat(inter.(float64), 'f', -1, 64), nil
	}
	return "", errors.New("参数错误")
}
