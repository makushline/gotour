package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) (z float64) {
    z      = 1.0
    oldz  := 2.0
    delta := 1e-5
    for math.Abs(oldz - z) > delta {
        oldz = z
        z = z - ((z*z - x) / (2*z))
    }
    return
}

func main() {
    fmt.Println(Sqrt(3.14))
    fmt.Println(math.Sqrt(3.14))
}
