package cast

import (
	"fmt"
)

// typeof returns a string containing the naame of the type of `v`.
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
