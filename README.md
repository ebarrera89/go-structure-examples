# Examples for my talk on structuring Go apps.

Note: I created an updated version of this repo [here](https://github.com/katzien/app-structure-examples)! These talks are 3+ years old now and I've changed my opinion on some of this stuff since then, so don't treat them as gospel. 🙂

## Links to video and slides:

_Note: these talks refer to [release (git tag) v1.0](https://github.com/katzien/go-structure-examples/releases/tag/v.10) of this repository._

2018-09-28 GoWayFest 2.0 [video](https://www.youtube.com/watch?v=Qtk9FFOoT5M&t=11s) | [slides](https://github.com/katzien/talks/blob/master/how-do-you-structure-your-apps/gowayfest2.0-2018-09-28/slides.pdf)

2018-09-19 London Gophers [video](https://www.youtube.com/watch?v=B5oQnECDJ8g) | [slides](https://github.com/katzien/talks/blob/master/how-do-you-structure-your-apps/londongophers-2018-09-19/slides.pdf)

2018-08-28 GopherCon 2018 [video](https://www.youtube.com/watch?v=oL6JBUk6tj0&t=245s) | [slides](https://github.com/katzien/talks/blob/master/how-do-you-structure-your-apps/gophercondenver-2018-08-28/slides.pdf)

### Older versions of this talk:

2018-06-02 GopherCon EU [video](https://www.youtube.com/watch?v=IjnfmYuFZss) | [slides](https://github.com/katzien/talks/blob/master/how-do-you-structure-your-apps/gopherconiceland-2018-06-02/slides.pdf)

2018-06-02 GopherCon UK [video](https://www.youtube.com/watch?v=VQym87o91f8&t=481s) | [slides](https://github.com/katzien/talks/blob/master/how-do-you-structure-your-apps/gopherconuk-2018-08-03/slides.pdf)

## How to run

Pick the example (directory) you want to run, then run main.go:

`go run main.go` or `go run cmd/beer-server/main.go`

### [Release v1.0 notes]((https://github.com/katzien/go-structure-examples/releases/tag/v.10))
Not all examples will compile. Only the flat, domain, domain-hex and domain-hex-actor will compile.
For more information about why that is, please refer to the talk.
