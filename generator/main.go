// Generate data structures based on the CSV available at https://github.com/datasets/country-codes
// +build ignore

package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

const url = "https://raw.githubusercontent.com/datasets/country-codes/master/data/country-codes.csv"

const tpl = `generator/data.go.tpl`
const outfile = `data.go`

type tplCountry struct {
	Alpha2  string
	Alpha3  string
	Numeric string
	Name    string
}

func main() {
	// Parse template file
	tmpl, err := template.ParseFiles(tpl)
	if err != nil {
		log.Fatalln(`iso3166: error parsing template %s:`, tpl, err)
	}

	// Open output file
	w, err := os.Create(outfile)
	if err != nil {
		log.Fatalln(`iso3166: error opening output file %s:`, outfile, err)
	}
	defer w.Close()

	countries := countryCodes()

	data := map[string]interface{}{
		"generator": "from CSV",
		"countries": countries,
	}

	// Render template into outfile
	if err := tmpl.Execute(w, data); err != nil {
		log.Fatalln(`iso3166: error rendering template:`, err)
	}

	log.Printf("iso3166: generated %d countries", len(countries))
}

func countryCodes() map[string]tplCountry {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(`iso3166: error downloading csv:`, err)
	}
	defer resp.Body.Close()

	var countries = make(map[string]tplCountry)

	decoder := csv.NewReader(resp.Body)

	headers, err := decoder.Read()
	if err != nil {
		log.Fatalln(`iso3166: error reading csv:`, err)
	}

	for {
		row, err := decoder.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(`iso3166: error reading csv:`, err)
			log.Fatalln(`iso3166: error reading csv:`, err)
		}

		// Map values onto header names
		record := make(map[string]string, len(headers))
		for i, value := range row {
			record[headers[i]] = value
		}

		if record["ISO3166-1-Alpha-3"] == "" {
			// Skip countries without code
			continue
		}

		var c = tplCountry{
			Name:    record["CLDR display name"],
			Alpha2:  record["ISO3166-1-Alpha-2"],
			Alpha3:  record["ISO3166-1-Alpha-3"],
			Numeric: record["ISO3166-1-numeric"],
		}

		countries[c.Alpha2] = c
	}
	return countries
}
