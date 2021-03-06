package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/xerrors"

	"github.com/mpppk/anacondaaaa"
)

func main() {
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe(":1323", nil))
}

func httpHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		query := req.URL.Query()
		parameters, ok := query["crc_token"]
		if !ok || len(parameters) == 0 {
			http.Error(w, "invalid query parameters", http.StatusInternalServerError)
			return
		}

		crcToken := parameters[0]
		twitterConsumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
		response := &anacondaaaa.CRCResponse{
			ResponseToken: anacondaaaa.CreateCRCToken(crcToken, twitterConsumerSecret),
		}

		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(res)
		return
	}

	if req.Method != "POST" {
		http.Error(w, "invalid HTTP Method", http.StatusBadRequest)
		return
	}

	if req.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
		return
	}

	events, err := parseJsonBody(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if events.GetEventName() == anacondaaaa.TweetCreateEventsEventName {
		w.WriteHeader(http.StatusOK)
		retText := fmt.Sprintf("tweet event is arrived. first tweet content: %#v", events.TweetCreateEvents[0])
		_, _ = w.Write([]byte(retText))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func parseJsonBody(req *http.Request) (*anacondaaaa.AccountActivityEvent, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, xerrors.Errorf("failed to read request body")
	}
	var accountActivityEvent anacondaaaa.AccountActivityEvent
	err = json.Unmarshal(body, &accountActivityEvent)
	if err != nil {
		return nil, xerrors.Errorf("failed to unmarshal request body to json: %w", err)
	}
	return &accountActivityEvent, nil
}
