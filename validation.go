package libminio

import (
	"errors"
)

// StringValidate is used to validate empty string
func StringValidate(param ...string) error {
	var isValid string
	for _, val := range param {
		if val == isValid {
			return errors.New("param is required")
		}
	}
	return nil
}
