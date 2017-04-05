package sportsData

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "os"
)

func GetTeamData(query string) (types.Teams, error ) {
    var uri = ""
    client := &http.Client{}

    request, err := http.NewRequest("GET", uri, nill)
    request.Header.Set("Content-Type", "https://api.github.com/user")
    request.Header.Set("Authorization", "e06c38d27add44d4da0d6c314429c584b9ebe401")
    request.Header.Set("Accept", "application/json")
    resp, err := client.Do(request)

    if err != nil {
          fmt.Println("error", err)
    }

    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("parse error", err)
    }

    fmt.Printf("%s", b)

    file := createFile("teams.txt")
    defer closeFile(file)
    writeToFile(file, b)
}

func createFile(path string)*os.File {
  fmt.Println("creating file")
  file, err := os.Create(path)
  if err != nil {
    panic(err)
  }
  return file
}

func writeToFile (file*os.File, dealerName string) {
  fmt.Println("writing to file")
  fmt.Fprintln(file, dealerName)
}

func closeFile(file*os.File) {
  fmt.Println("closing file")
  file.Close()
}
