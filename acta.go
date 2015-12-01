// golang web server 
// test for google cloud deploy
// AQ <aq@okaq.com>
// 2015-12-02
package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "acta.html")
    })
    http.ListenAndServe(":8008", nil)
}
