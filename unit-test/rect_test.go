package unittest

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

var (
  rect      Rect    = Rect{4}
  realVol   float64 = 64
  realSurf  float64 = 96
  realCirc  float64 = 48
)

func TestVol(t *testing.T) {
  assert.Equal(t, rect.Volume(), realVol, "wrong bitch")
} 

func TestSurf(t *testing.T) {
  assert.Equal(t, rect.SurfaceArea(), realSurf, "wrong bitch")
} 

func TestCirc(t *testing.T) {
  assert.Equal(t, rect.Circ(), realCirc, "wrong bitch")
} 
