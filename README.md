# iso3166
[![GoDoc](https://godoc.org/github.com/ferdypruis/iso3166?status.svg)](https://godoc.org/github.com/ferdypruis/iso3166)
[![GolangCI](https://golangci.com/badges/github.com/ferdypruis/iso3166.svg)](https://golangci.com/r/github.com/ferdypruis/iso3166)

A Go package providing all ISO 3166-1 country codes as constants of type `iso3166.Country`.

For each country the two-letter and three-letter alphabetic codes, the three-digit numeric code
and the English name are available.

Countries can either be hardcoded using the available constants or loaded from a string using `FromAlpha2()`, 
`FromAlpha3()` and `FromNumeric()`.

## Examples
Use the constants to directly reference countries.
```go
fmt.Println("The numeric code for Antarctica is", iso3166.AQ.Numeric())

// Output:
// The numeric code for Antarctica is 010
```

Use `FromAlpha2()`, `FromAlpha3()` and `FromNumeric()` to load a country from a string.
```go
country, _ := iso3166.FromAlpha2("NL") // Ignoring error for simplicity
fmt.Println("The three-letter code for", country.Name(), "is", country.Alpha3())

// Output:
// The three-letter code for Netherlands is NLD
```

Wrap in `Must()` to return a single value or panic on error;
```go
fmt.Println("The two-letter code for the United States is", iso3166.Must(iso3166.FromAlpha3("USA")).Alpha2())

// Output:
// The two-letter code for the United States is US
```

## Source
- [DataHub Core Data](https://datahub.io/core/country-codes)
