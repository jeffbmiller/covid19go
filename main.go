package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func covid(w http.ResponseWriter, r *http.Request) {
	// Make HTTP GET request
	response, err := http.Get("http://www.google.ca/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find and print image URLs
	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			fmt.Println(imgSrc)
		}
	})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/covid", covid)
	log.Fatal(http.ListenAndServe(":8080", router))
}
