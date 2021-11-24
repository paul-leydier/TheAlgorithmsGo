package pi

import (
	"fmt"
	"testing"
)

func TestChudnovsky(t *testing.T) {
	pi, prec := Chudnovsky(1000)
	fmt.Println(pi.Text('f', prec))
	fmt.Printf("Precision: %d decimals", prec)
}
