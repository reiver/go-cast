package cast

import (
	"fmt"
)

// typeof returns a string containing the naame of the type of `v`.
func typeof(v any) string {
	return fmt.Sprintf("%T", v)
}
