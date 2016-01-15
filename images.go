package main

import(
  "golang.org/x/tour/pic"
  "image"
  "image/color"
)

type Image [][]uint8

func main() {
  i := NewImage(255, 255, func(a, b int) uint8 { return uint8((a ^ b) / 1 + (a + b)) }) // Need the one so you don't divide by zero
  pic.ShowImage(i)
}

func (i Image) ColorModel() color.Model {
  return color.RGBAModel 
}

func (i Image) Bounds() image.Rectangle {
  dx, dy := len(i), 0
  if dx > 0 {
    dy = len(i[0])
  }
  return image.Rect(0, 0, dx, dy)
}

func (i Image) At(x, y int) color.Color {
  c := i[x][y]
  return color.RGBA{c, c,255, 255}
}

func NewImage(x int, y int, f func(int, int) uint8 ) (m Image) {
  m = make([][]uint8, x)
  for i := range m {
    m[i] = make([]uint8, y)
    for j := range m[i] {
      m[i][j] = f(i,j)
    }
  }
  return
}