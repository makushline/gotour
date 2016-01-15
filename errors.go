package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

const delta = 1e-5

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %.0f", e)
}

func Sqrt(x float64) (float64, error) {
    z      := 1.0
    if x < 0.0 {
        return -z, ErrNegativeSqrt(x)
    }
    oldz  := 2.0
    for math.Abs(oldz - z) > delta {
        oldz = z
        z = z - ((z*z - x) / (2*z))
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
