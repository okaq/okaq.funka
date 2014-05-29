// golang web server for GAE
// AQ <aq@okaq.com>
// 2014-05-28
package main

import (
	"net/http"
)

const (
	PATH = "index.html"
	// GUKA = "guka.html"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, PATH)
	})
}
