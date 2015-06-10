package linodego

import (
	"fmt"
)

// CustomBool is a type to handle Linode's insistance of using ints as boolean values.
type CustomBool struct {
	bool
}

func (cb *CustomBool) UnmarshalJSON(b []byte) error {
	if len(b) != 1 {
		return fmt.Errorf("Unable to marshal value with length %d into a CustomBool.", len(b))
	}
	if int(b[0]) == 0 {
		cb.bool = false
	} else if int(b[0]) == 1 {
		cb.bool = true
	}
	return nil
}

func (cb *CustomBool) MarshalJSON() ([]byte, error) {
	if cb == nil {
		return []byte{}, fmt.Errorf("Unable to marshal nil value for CustomBool")
	}
	if cb.bool {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}
