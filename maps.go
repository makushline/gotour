package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) (wordCount map[string]int) {
    wordCount = make(map[string]int)
    for _, v := range strings.Fields(s)[:] {
        wordCount[v]++
    }
    return wordCount
}

func main() {
    wc.Test(WordCount)
}
