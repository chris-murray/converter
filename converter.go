// Package converter performs misc conversions.
package converter

import "fmt"

type LightYear float64
type Parsec float64

const (
	PinLY Parsec = 0.306601
  LYInP LightYear = 3.26
)

func (p Parsec) String() string    { return fmt.Sprintf("%gp", p) }
func (ly LightYear) String() string { return fmt.Sprintf("%gly", ly) }
