// converts between Parsec and LightYear
package converter

// return the number of parsecs for the given distance in lightyears
func LYtoParsec (lr LightYear) Parsec {
  return Parsec(lr * ParsecRef )
}

// return the number of parsecs for the given distance in lightyears
func PtoLY (p Parsec) LightYear {
  return LightYear(p * LyInP )
}


//!-
