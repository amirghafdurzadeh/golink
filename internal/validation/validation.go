package validation

import (
	"fmt"
)

func Required(value, name string) error {
	if value == "" {
		return fmt.Errorf("%s is required", name)
	}
	return nil
}
