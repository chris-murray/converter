// converts between Parsec and LightYear
package converter

// return the number of parsecs for the given distance in lightyears
func LYtoParsec (ly LightYear) Parsec {
  return Parsec(ly) * PinLY
}

// return the number of parsecs for the given distance in lightyears
func PtoLY (p Parsec) LightYear {
  return LightYear(p) * LYInP
}


//!-
