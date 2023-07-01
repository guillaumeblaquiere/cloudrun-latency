package main

import (
	"fmt"
	"google.golang.org/api/idtoken"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/pong", pong)
	http.ListenAndServe(":8080", nil)
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ping(w http.ResponseWriter, r *http.Request) {
	//Get the url to call
	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "a target URL must be provided")
		return
	}

	//Get the number of call
	nbCall := r.URL.Query().Get("nbcall")
	nbCallValue, err := strconv.Atoi(nbCall)
	if nbCall == "" {
		nbCallValue = 10
	} else if err != nil || nbCallValue <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "nbcall must be a positive integer")
		return
	}

	//Get the mode (secure or insecure)
	secured := r.URL.Query().Get("useGoogleClient")
	useGoogleClient := true
	if strings.ToLower(secured) == "false" {
		useGoogleClient = false
	}

	client := http.DefaultClient
	if useGoogleClient {
		client, err = idtoken.NewClient(r.Context(), url)
	}
	fmt.Println(client, err)

	totalDuration, _ := time.ParseDuration("0s")

	for i := 0; i < nbCallValue; i++ {
		targetUrl := fmt.Sprintf("%s?nbcall=%d", url, i)
		timestart := time.Now()
		resp, err := client.Get(targetUrl)
		timeend := time.Now()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error during the query: %s. Exit now", err)
			return
		}
		if resp.StatusCode != 200 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "invalid status code: %s. Exit now", resp.Status)
			return
		}
		fmt.Printf("the request nb %d took %s\n", i, timeend.Sub(timestart))
		fmt.Fprintf(w, "the request nb %d took %s\n", i, timeend.Sub(timestart))
		totalDuration += timeend.Sub(timestart)
	}

	fmt.Fprintf(w, "\ntotal duration for %dms request is %s\n", totalDuration.Milliseconds(), nbCall)
	fmt.Fprintf(w, "average duration is %dms\n", totalDuration.Milliseconds()/int64(nbCallValue))
}
