// converts between Parsec and LightYear
package converter

// return the number of parsecs for the given distance in lightyears
func LRtoParsec (lr LightYear) Parsec {
  return Parsec(lr * ParsecRef )
}

// return the number of parsecs for the given distance in lightyears
func PtoLR (p Parsec) LightYear {
  return LightYear(p * LyInP )
}


//!-
