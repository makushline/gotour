package main

import (
        "io"
        "os"
        "strings"
)

type rot13Reader struct {
        r io.Reader
}

func (reader rot13Reader) Read(p []byte) (n int, err error) {
        n, err = reader.r.Read(p)
        for i := 0; i < n; i++ {
            if ((p[i] >= 'a' && p[i] <= 'm') || (p[i] >= 'A' && p[i] <= 'M')) {
                p[i] = p[i] + 13
            } else if ((p[i] > 'm' && p[i] <= 'z') || (p[i] > 'M' && p[i] <= 'Z')) {
                p[i] = p[i] - 13
            }
        }
        return
}

func main() {
        s := strings.NewReader("Lbh penpxrq gur pbqr!")
        r := rot13Reader{s}
        io.Copy(os.Stdout, &r)
}