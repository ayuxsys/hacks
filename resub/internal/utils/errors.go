package utils

import "fmt"

func WrapError(context string, err error) error {
	return fmt.Errorf("%q: %w", context, err)
}
