mandrill
========

A Go library for handling incoming events from the [Mandrill Webhook
service](http://help.mandrill.com/entries/21738186-Introduction-to-Webhooks).

[Docs](https://godoc.org/github.com/leejarvis/mandrill)

Install
-------

    go get github.com/leejarvis/mandrill

Example
-------

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/leejarvis/mandrill"
)

func main() {
	  http.Handle("/", mandrill.EventHandler(func(e mandrill.Event) {
        fmt.Println(e.Type) // sent/click/deferred
        // insert the event into our database or something
	  }))

	  http.ListenAndServe(":3000", nil)
}
```

The `mandrill.EventHandler` function returns a `http.Handler` to be
compliant with `net/http`. This allows you to easily mount the handler
as middleware to your existing web app.

This function takes an argument of type `func(mandrill.Event)` and will
be called for every event parsed from the webhook request.
