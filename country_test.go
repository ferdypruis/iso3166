package iso3166_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/ferdypruis/iso3166"
)

func TestCountry_Name(t *testing.T) {
	tests := []struct {
		name string
		c    iso3166.Country
		want string
	}{
		{"AD", iso3166.AD, "Andorra"},
		{"AD", iso3166.NL, "Netherlands"},
		{"ZW", iso3166.ZW, "Zimbabwe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromAlpha2(t *testing.T) {
	want := iso3166.NL
	got, err := iso3166.FromAlpha2("NL")
	if err != nil {
		t.Errorf("FromAlpha2() error = %v, wantErr nil", err)
	}
	if got != want {
		t.Errorf("FromAlpha2() got = %v, want %v", got, want)
	}
}

func TestFromAlpha2Error(t *testing.T) {
	_, err := iso3166.FromAlpha2("00")
	if _, ok := err.(iso3166.Error); !ok {
		t.Fatalf("FromAlpha2() error %T, want iso3166.Error", err)
	}
}

func TestFromAlpha3(t *testing.T) {
	want := iso3166.US
	got, err := iso3166.FromAlpha3("USA")
	if err != nil {
		t.Errorf("FromAlpha3() error = %v, wantErr nil", err)
	}
	if got != want {
		t.Errorf("FromAlpha3() got = %v, want %v", got, want)
	}
}

func TestFromAlpha3Error(t *testing.T) {
	_, err := iso3166.FromAlpha3("000")
	if _, ok := err.(iso3166.Error); !ok {
		t.Fatalf("FromAlpha3() error %T, want iso3166.Error", err)
	}
}

func TestFromNumeric(t *testing.T) {
	want := iso3166.PM
	got, err := iso3166.FromNumeric("666")
	if err != nil {
		t.Errorf("FromNumeric() error = %v, wantErr nil", err)
	}
	if got != want {
		t.Errorf("FromNumeric() got = %v, want %v", got, want)
	}
}

func TestFromNumericError(t *testing.T) {
	_, err := iso3166.FromNumeric("AAA")
	if _, ok := err.(iso3166.Error); !ok {
		t.Fatalf("FromNumeric() error %T, want iso4217.Error", err)
	}
}

func TestMust(t *testing.T) {
	want := iso3166.FM
	got := iso3166.Must(want, nil)

	if got != want {
		t.Errorf("Must() got = %v, want %v", got, want)
	}
}

func TestMustPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Must() did not panic")
		}
	}()

	iso3166.Must(iso3166.Country(0), errors.New("this should cause panic"))
}

func TestErrorString(t *testing.T) {
	err := iso3166.Error("test error")
	if !strings.HasPrefix(err.Error(), "iso3166:") {
		t.Fatalf("Error.String() %q, want prefix 'iso3166:'", err)
	}
}

func ExampleCountry() {
	fmt.Println("The numeric code for Antarctica is", iso3166.AQ.Numeric())

	// Output:
	// The numeric code for Antarctica is 010
}

func ExampleFromAlpha2() {
	alpha2 := "NL"
	country, _ := iso3166.FromAlpha2(alpha2) // Ignoring error for simplicity
	fmt.Println("The three-letter code for", country.Name(), "is", country.Alpha3())

	// Output:
	// The three-letter code for Netherlands is NLD
}

func ExampleFromAlpha3() {
	alpha3 := "USA"
	country, _ := iso3166.FromAlpha3(alpha3) // Ignoring error for simplicity
	fmt.Println("The three-digit code for", country.Name(), "is", country.Numeric())

	// Output:
	// The three-digit code for US is 840
}

func ExampleFromNumeric() {
	numeric := "666"
	country, _ := iso3166.FromNumeric(numeric) // Ignoring error for simplicity
	fmt.Println("The three-letter code for", country.Name(), "is", country.Alpha3())

	// Output:
	// The three-letter code for St. Pierre & Miquelon is SPM
}

func ExampleMust() {
	fmt.Println("The two-letter code for the United States is", iso3166.Must(iso3166.FromAlpha3("USA")).Alpha2())

	// Output:
	// The two-letter code for the United States is US
}
