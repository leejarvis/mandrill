package mandrill

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadRequest(t *testing.T) {
	req, err := http.NewRequest("HEAD", "http://example.com/events", nil)
	assert.Nil(t, err)

	ep := func(e Event) {}

	w := httptest.NewRecorder()
	EventHandler(ep).ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
}

func TestProcessing(t *testing.T) {
	var events Events
	ep := func(e Event) {
		events = append(events, e)
	}

	ts := httptest.NewServer(EventHandler(ep))
	defer ts.Close()

	res, err := http.PostForm(ts.URL, url.Values{"mandrill_events": []string{eventsJSON}})
	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, 2, len(events))
}
