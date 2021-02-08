package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	http.HandleFunc("/", matchHandler)

	// Serve http handler on port 8080
	go func() {
		log.Println("Serving on 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Wait here until CTRL-C or other term signal is received.
	<-sc
	log.Print("Server Stopped")
}

// Structure that holds the two sample strings
type testStrings struct {
	Samples []string
}

// Finds the length of the submitted samples
func (ts *testStrings) length() []int {
	return []int{len(ts.splitStrings()[0]), len(ts.splitStrings()[1])}
}

// Splits each sample into an array of words. Returns a 2D array.
func (ts *testStrings) splitStrings() [][]string {
	var s [][]string
	for _, v := range ts.Samples {
		//fmt.Println(v)
		s = append(s, strings.Split(v, " "))
	}
	return s
}

// Finds all matching works between the two samples and calculates the similarity. Returns a percentage.
func (ts testStrings) matchStrings() string {
	sampleOne := ts.splitStrings()[0]
	sampleTwo := ts.splitStrings()[1]
	var matches [][]int
	// Set string one (sampleOne) to be the reference string to compare string two (sampleTwo) against
	// Iterate through sampleOne
	for k, v := range sampleOne {
		// Compare reference word to string two word.
		for k2, v2 := range sampleTwo {
			// If words match, add matched indexes to matches array and remove matched word from string two.
			if v == v2 {
				matches = append(matches, []int{k, k2})
				copy(sampleTwo[k2:], sampleTwo[k2+1:])
				sampleTwo[len(sampleTwo)-1] = ""
				sampleTwo = sampleTwo[:len(sampleTwo)-1]
				// Break from loop so not to double match words.
				break
			}
		}
	}

	//calculate percentage of total unmatched words
	return fmt.Sprintf("Samples have "+"%.2f%% similatrity.", 100*(1-(float64(len(sampleTwo)+(ts.length()[0]-len(matches)))/float64((ts.length()[0]+ts.length()[1])))))
}

// function to handle http post requests.
func matchHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	r.ParseForm()

	var ts testStrings

	// Validate there are two samples.
	if len(r.Form) == 2 {
		for k := range r.Form {
			ts.Samples = append(ts.Samples, r.PostFormValue(k))
		}
		// Write result as a http response
		fmt.Fprintln(w, ts.matchStrings())
		// Log the result
		log.Println(ts.matchStrings())
	} else {
		fmt.Fprintln(w, "Not the correct number of parameters")
		log.Println("Not the correct number of parameters")
	}
}
