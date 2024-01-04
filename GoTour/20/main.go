package main

import (
	"fmt"
	"math"
	"time"
)

type ErrNegativeSqrt struct {
	When time.Time
	What float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt negative number: %v", e.What)
}

func ValidNumber(x float64) error {
	if x < 0 {
		return &ErrNegativeSqrt{time.Now(), x}
	}
	return nil
}

func Sqrt(x float64) (float64, error) {
	if err := ValidNumber(x); err != nil {
		return 0, err
	}
	z := 1.0
	for delta := 1.0; math.Abs(delta) > 1e-15; {
		y := z - (z*z-x)/(2*z)
		fmt.Println(delta)
		delta = y - z
		z = y
	}
	return z, nil
}

func main() {
	if result, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	if result, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
