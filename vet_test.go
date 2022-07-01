package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestPrintType(t *testing.T) {
	// str := "hello world"
	// fmt.Printf("%d\n", str)
}

func TestBoolean(t *testing.T) {
	// var i int
	// fmt.Println(i != 0 || i != 1)
}

func TestRangeRoutine(t *testing.T) {
	words := []string{"foo", "bar", "baz"}
	for _, word := range words {
		go func() {
			fmt.Println(word)
		}()
	}
}

func TestWhat(t *testing.T) {
	res, err := http.Get("https://www.spreadsheetdb.io/")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
