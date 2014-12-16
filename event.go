package mandrill

import "time"

// Timestamp is a wrapper around int64 to represent a Unix
// timestamp.
type Timestamp int64

// Time returns the local time from the Unix timestamp.
func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}

type Event struct {
	Type      string    `json:"event"`
	ID        string    `json:"_id"`
	Timestamp Timestamp `json:"ts"`

	// for click events
	UserAgent       string
	IP              string `json:"ip"`
	Location        location
	UserAgentParsed userAgent `json:"user_agent_parsed"`

	// email message attributes
	Msg msg `json:"msg"`
}

// Metadata fetches the value for `key` for the Msg
// Metadata map. Purely convenience.
func (e Event) Metadata(key string) interface{} {
	return e.Msg.Metadata[key]
}

type msg struct {
	Metadata  map[string]interface{}
	Timestamp Timestamp `json:"ts"`
	Subject   string
	Email     string
	State     string
	Sender    string
	Tags      []string
	Clicks    []struct {
		URL       string    `json:"url"`
		Timestamp Timestamp `json:"ts"`
	}
	Opens []struct {
		Timestamp Timestamp `json:"ts"`
	}
}

type location struct {
	City        string
	Country     string
	Lon         float64 `json:"longitude"`
	Lat         float64 `json:"latitude"`
	CountryCode string  `json:"country_short"`
	Timezone    string
	PostalCode  string `json:"postal_code"`
	Region      string
}

// TODO implement this
type userAgent struct {
}
