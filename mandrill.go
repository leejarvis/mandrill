package mandrill

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type EventProcessor func(Event)

type Events []Event

// EventsFromReader reads and decodes JSON from r and
// returns a slice of Event types.
func EventsFromReader(r io.Reader) (Events, error) {
	var events Events
	return events, json.NewDecoder(r).Decode(&events)
}

// EventHandler implements the net/http handler interface. It
// responds with a 200 for any HEAD request (as per the Mandrill
// docs). ep is called for every event received.
func EventHandler(ep EventProcessor) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// HEAD is sent to verify the endpoint exists
		if r.Method == "HEAD" {
			w.WriteHeader(200)
			return
		}

		eventData := r.PostFormValue("mandrill_events")
		if len(eventData) == 0 {
			http.Error(w, "no events submitted", 400)
			return
		}

		// we have events, decode them
		events, err := EventsFromReader(strings.NewReader(eventData))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		for _, e := range events {
			ep(e)
		}

		w.WriteHeader(200)
	})
}
