package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func main(){

  //3 lines \/ handles root path GET->calling fn handler,
  // serves js file for html script,
  // starts web server 
  http.HandleFunc("/", handler)
  http.Handle("/js/", http.FileServer(http.Dir("../dist")))
  fmt.Println(http.ListenAndServe(":8080", nil))
}

// handler* retrieves and parses index html && checks for errors
func handler(w http.ResponseWriter, r *http.Request) {
  templates := template.Must(template.ParseFiles("../dist/index.html"))
  if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
