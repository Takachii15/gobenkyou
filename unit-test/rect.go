package unittest

import "math"

type Rect struct{
  side float64
}

func (k Rect) Volume() float64 {
  return math.Pow(k.side, 3)
}

func (k Rect) SurfaceArea() float64 {
  return math.Pow(k.side, 2) * 6
}

func (k Rect) Circ() float64 {
  return k.side * 12
}

type Dipt interface{

}
