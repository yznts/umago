
<p align="center">
    <img width="200" src="LOGO.png" />
</p>

<h1 align="center">umago</h1>

<p align="center">
    Umami tracking inside of your Go application.
</p>

```go
import "github.com/yuriizinets/umago"
```

This library allows to track events into [Umami](https://umami.is) inside of your Go application.

Use cases:

- Server-side events tracking
- Bring-your-own-tracker option
- Bypass Google _"malicious"_ tracking protection
- Track completely custom/insane things (f.e. use Umami to track CLI app usage)

## Usage

Here are the simplest examples of how to use umami-go.
This examples are not covering all provided features,
so feel free to explore sources and godoc.

### Middleware

Example using middleware:

```go
package main

import (
	"net/http"

	"github.com/yuriizinets/umago"
)

var (
	// Define umami configuration
	umamiConfiguration = umago.Configuration{
		Href:    "https://umami.foo.bar", // Provide your Umami instance url here
		Website: "test",                  // Provide your website id here
	}
	// Initialize umami middleware
	umamiMiddleware = umago.NewMiddleware(umamiConfiguration)
)

// PageToTrack is a simple page handler that we want to track.
func page(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`It's a simple text page!`))
}

func main() {
	// Register page
	http.HandleFunc("/", umamiMiddleware(page))

	// Serve
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
```

### Handler (JS fetch)

Example using handler and JS fetch:

```go
package umami_test

import (
	"net/http"

	"github.com/yuriizinets/umago"
)

var (
	// Define umami configuration
	umamiConfiguration = umago.Configuration{
		Href:    "https://umami.foo.bar", // Provide your Umami instance url here
		Website: "test",                  // Provide your website id here
	}
	// Initialize umami handler
	umamiHandler = umago.NewHandler(umamiConfiguration)
)

// page is a simple page handler that we want to track.
// Please note that page must include some tracking code to send event.
func page(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<html>
		<head><title>This is example page!</title></head>
		<body>
			<script>
				fetch('/track', {
					method: 'POST',
					body: JSON.stringify({
						r: document.referrer,
					})
				})
			</script>
		</body>
		</html>
	`))
}

func main() {
	// Register page
	http.HandleFunc("/", page)
	// Register umami handler
	http.HandleFunc("/track", umamiHandler)

	// Serve
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
```

### Handler (script)

Example using handler and script reference.
Please note, this method is not using referenced script (it's empty),
but uses script request to track page view.

```go
// Souce is the same as for previous example,
// only page handler and track url are different.

...

func page(w http.ResponseWriter, r *http.Request) {
	ref := r.Referer()
	w.Write([]byte(fmt.Sprintf(`
		<html>
		<head><title>This is example page!</title></head>
		<body>
			<script src="/track.js?r=%s"></script>
		</body>
		</html>
	`, ref)))
}

func main() {

	...

	http.HandleFunc("/track.js", umamiHandler)
}

...

```

### Handler (pixel)

Example using handler and pixel reference.
In this case, handler will track page view and return 1x1 transparent png image.

```go
// Souce is the same as for previous example,
// only page handler and track url are different.

...

func page(w http.ResponseWriter, r *http.Request) {
	ref := r.Referer()
	w.Write([]byte(fmt.Sprintf(`
		<html>
		<head><title>This is example page!</title></head>
		<body>
			<img src="/track.png?r=%s" />
		</body>
		</html>
	`, ref)))
}

func main() {

	...

	http.HandleFunc("/track.png", umamiHandler)
}

...

```

## Saying thanks to Google

Because of Google's _"malicious"_ tracking protection
my company got an issue with paid ads,
with no details about what's happening.
We had to spend our time to find out actual reason
and prototype a workaround solution.
And, here we are, we've got a separate small library!

Google acts very anticompetitive to other analytics.
In the same time, Google Analytics sucks a lot.
Yes, it provides rich functionality for reporing.
But look at the tracking script, it's huge!
Even PageSpeed Insights complains about it.
GA UI requires a steep learning curve to use it,
even for simple things.

## Credits

- [Umami](https://umami.is) - awesome open-source analytics
- [Alexandra Metifieva](https://t.me/rossskosh) - thanks for the logo!
- [Google Analytics](https://analytics.google.com) - thanks for wasting our time!
