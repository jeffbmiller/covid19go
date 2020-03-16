package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/countries", allCountriesHandler)
	router.HandleFunc("/countries/{name}", countryHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func countryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	countryName := params["name"]

	countries := parseCountries()

	var country Country

	for _, c := range countries {
		if strings.ToLower(c.Name) == strings.ToLower(countryName) {
			country = c
			break
		}
	}

	foo, _ := json.Marshal(country)
	fmt.Fprintf(w, string(foo))
}

func allCountriesHandler(w http.ResponseWriter, r *http.Request) {

	countries := parseCountries()

	foo, _ := json.Marshal(countries)
	fmt.Fprintf(w, string(foo))
}

func parseCountries() []Country {
	response, err := http.Get("https://www.worldometers.info/coronavirus/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	countries := []Country{}
	// Find and print image URLs
	document.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			if indextr != 0 {
				mapCountry := make(map[string]interface{})

				rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
					if indexth == 0 {
						mapCountry["name"] = strings.Trim(tablecell.Text(), "\t \n")
					}
					if indexth == 1 {
						mapCountry["totalCases"] = strings.Trim(tablecell.Text(), "\t \n")
					}
					if indexth == 2 {
						mapCountry["newCases"] = strings.Trim(tablecell.Text(), "\t \n")
					}
					if indexth == 3 {
						mapCountry["totalDeaths"] = strings.Trim(tablecell.Text(), "\t \n")
					}
					if indexth == 4 {
						mapCountry["newDeaths"] = strings.Trim(tablecell.Text(), "\t \n")
					}
					if indexth == 5 {
						mapCountry["totalRecovered"] = strings.Trim(tablecell.Text(), "\t \n")
					}
					if indexth == 6 {
						mapCountry["activeCases"] = strings.Trim(tablecell.Text(), "\t \n")
					}
					if indexth == 7 {
						mapCountry["seriousCritical"] = strings.Trim(tablecell.Text(), "\t \n")
					}
				})

				var country Country
				err := mapstructure.Decode(mapCountry, &country)
				if err != nil {
					// error
				}

				if country.Name != "Total:" {
					countries = append(countries, country)
				}
			}
		})
	})
	return countries
}

type Country struct {
	Name            string
	TotalCases      string
	NewCases        string
	TotalDeaths     string
	NewDeaths       string
	TotalRecovered  string
	ActiveCases     string
	SeriousCritical string
}
