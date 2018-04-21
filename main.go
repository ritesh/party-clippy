package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
)

const (
	clippy = `
 _________________________________
/ It looks like you're building a \
\ microservice.                   /
 ---------------------------------
 \
  \
     __
    /  \
    |  |
    @  @
    |  |
    || |/
    || ||
    |\_/|
    \___/
`
)

var (
	colorOptions = []*color.Color{
		color.New(color.FgHiRed),
		color.New(color.FgHiGreen),
		color.New(color.FgHiYellow),
		color.New(color.FgHiBlue),
		color.New(color.FgHiMagenta),
		color.New(color.FgHiCyan),
	}
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Go to top of terminal and clear
		i := 0
		for i < len(colorOptions) {
			// Clear the terminal.
			fmt.Fprintf(w, "\033c")
			// Print clippy with a color.
			colorOptions[i].Fprintf(w, clippy)
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			time.Sleep(time.Second / 4)
			if i == len(colorOptions)-1 {
				i = 0
				continue
			}
			i++
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}