package main

import(
    "golang.org/x/tour/reader"
    "fmt"
)

type MyReader struct{}

func (reader MyReader) Read(b []byte) (int, error) {
    if len(b) == 0 {
        return 0, fmt.Errorf("Buffer is empty")
    }
    for i, _ := range b {
        b[i] = 'A'
    }
    return 1, nil
}

func main() {
    reader.Validate(MyReader{})
}