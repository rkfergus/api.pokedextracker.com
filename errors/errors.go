package errors

import "fmt"

func NotFound(resource string) error {
	return fmt.Errorf("%s not found", resource)
}
