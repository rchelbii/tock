package pow

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func intToHex[T constraints.Signed](n T) string {
	return fmt.Sprintf("%x", n)
}
