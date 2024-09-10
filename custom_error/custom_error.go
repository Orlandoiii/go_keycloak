package custom_error

import "fmt"

func WrapError(message string, err error) error {
	return fmt.Errorf("%s %w", message, err)
}
