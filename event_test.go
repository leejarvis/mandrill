package mandrill

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testEvents Events

func init() {
	var err error
	testEvents, err = EventsFromReader(strings.NewReader(eventsJSON))

	if err != nil {
		panic(err)
	}
}

func TestMetadata(t *testing.T) {
	event := testEvents[0]
	assert.Equal(t, event.Metadata("user_id"), float64(111))
	assert.Nil(t, event.Metadata("unknown"))
}

var eventsJSON = `
[
   {
      "event" : "send",
      "_id" : "exampleaaaaaaaaaaaaaaaaaaaaaaaaa",
      "msg" : {
         "metadata" : {
            "user_id" : 111
         },
         "clicks" : [],
         "sender" : "example.sender@mandrillapp.com",
         "opens" : [],
         "ts" : 1365109999,
         "subject" : "This an example webhook message",
         "email" : "example.webhook@mandrillapp.com",
         "state" : "sent",
         "tags" : [
            "webhook-example"
         ],
         "_id" : "exampleaaaaaaaaaaaaaaaaaaaaaaaaa",
         "_version" : "exampleaaaaaaaaaaaaaaa"
      },
      "ts" : 1418683821
   },
   {
      "_id" : "exampleaaaaaaaaaaaaaaaaaaaaaaaaa1",
      "event" : "deferral",
      "ts" : 1418683821,
      "msg" : {
         "sender" : "example.sender@mandrillapp.com",
         "metadata" : {
            "user_id" : 111
         },
         "clicks" : [],
         "opens" : [],
         "email" : "example.webhook@mandrillapp.com",
         "subject" : "This an example webhook message",
         "ts" : 1365109999,
         "_version" : "exampleaaaaaaaaaaaaaaa",
         "smtp_events" : [
            {
               "source_ip" : "127.0.0.1",
               "size" : 0,
               "type" : "deferred",
               "diag" : "451 4.3.5 Temporarily unavailable, try again later.",
               "destination_ip" : "127.0.0.1",
               "ts" : 1365111111
            }
         ],
         "_id" : "exampleaaaaaaaaaaaaaaaaaaaaaaaaa1",
         "tags" : [
            "webhook-example"
         ],
         "state" : "deferred"
      }
   }
]`
