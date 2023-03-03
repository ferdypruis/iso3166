// Package iso3166 provides the ISO 3166-1 codes for the representation of countries.
package iso3166

// Country is a representation of a country.
type Country uint16

// Alpha2 returns the ISO 3166-1 two-letter alphabetic code.
func (c Country) Alpha2() string { return countries[c].alpha2 }

// Alpha3 returns the ISO 3166-1 three-letter alphabetic code.
func (c Country) Alpha3() string { return countries[c].alpha3 }

// Numeric returns the ISO 3166-1 three-digit numeric code.
func (c Country) Numeric() string { return countries[c].numeric }

// Name returns the English name.
func (c Country) Name() string { return countries[c].name }

// FromAlpha2 returns Country for the two-letter alpha2 code.
// Or an error if it does not exist.
func FromAlpha2(alpha2 string) (Country, error) {
	if c, ok := fromAlpha2[alpha2]; ok {
		return c, nil
	}
	return Country(0), Error("no country exists with alpha2-code " + alpha2)
}

// FromAlpha3 returns Country for the three-letter alpha3 code.
// Or an error if it does not exist.
func FromAlpha3(alpha3 string) (Country, error) {
	if c, ok := fromAlpha3[alpha3]; ok {
		return c, nil
	}
	return Country(0), Error("no country exists with alpha3-code " + alpha3)
}

// FromNumeric returns Country for the three-digit numeric code.
// Or an error if it does not exist.
func FromNumeric(numeric string) (Country, error) {
	if c, ok := fromNumeric[numeric]; ok {
		return c, nil
	}
	return Country(0), Error("no country exists with numeric-code " + numeric)
}

// FromName returns Country for the name.
// Or an error if it does not exist.
func FromName(name string) (Country, error) {
	if c, ok := fromName[name]; ok {
		return c, nil
	}
	return Country(0), Error("no country exists with name " + name)
}

// Must panics if err is non-nil and otherwise returns c.
// Could be used to return a single value from FromAlpha2/FromAlpha3/FromNumeric.
func Must(c Country, err error) Country {
	if err != nil {
		panic(err)
	}
	return c
}

// Error is the type of error returned by this package
type Error string

func (e Error) Error() string { return "iso3166: " + string(e) }
