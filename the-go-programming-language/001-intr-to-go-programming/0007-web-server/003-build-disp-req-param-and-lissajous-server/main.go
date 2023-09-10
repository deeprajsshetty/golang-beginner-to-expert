// Server is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
)

var palette = []color.Color{color.White, color.Black}

const (
	//whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// Register handlers for root ("/") and "/image" paths
	http.HandleFunc("/", handler)
	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})

	// Start the HTTP server on port 8000 and handle requests.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Display information about the incoming HTTP request
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// Parse and display the form data if available
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// lissajous generates a Lissajous curve animation and writes it to the provided http.ResponseWriter.
func lissajous(out http.ResponseWriter) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// Handle encoding errors properly here
	if err := gif.EncodeAll(out, &anim); err != nil {
		// You can add error handling logic here, e.g., log the error or return it.
	}
}

// Input:
// go run the-go-programming-language/001-intr-to-go-programming/0007-web-server/003-build-disp-req-param-and-lissajous-server/main.go

// Server Running:
// 1. curl http://localhost:8000
// 2. http://localhost:8000/image

// Output:

//1.
//GET / HTTP/1.1
//Header["User-Agent"] = ["curl/8.1.2"]
//Header["Accept"] = ["*/*"]
//Host = "localhost:8000"
//RemoteAddr = "127.0.0.1:56022"

//2.
// Displayed image in the browser
