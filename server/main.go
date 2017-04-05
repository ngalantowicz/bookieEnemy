package main

import (
  //"bookieEnemy/sportsData"
  "fmt"
  "html"
  "net/http"
)

func main(){
  http.HandleFunc("/", handler)
  http.ListenAndServe("localhost:8181", nil)

  // resp, err := http.Get("http://www.google.com")
  // if err != nil {
  //   panic(err)
  // }
}

func handler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
