package sportsData

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
)

func GetTeamData() {
    var uri = "https://api.sportradar.us/nfl-ot1/games/2016/reg/schedule.json?api_key=j6hpwhu26eqnczatrr2vers9"
    client := &http.Client{}

    request, err := http.NewRequest("GET", uri, nil)
    request.Header.Set("Content-Type", "application/json")
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

    // file := createFile("teams.txt")
    // defer closeFile(file)
    // writeToFile(file, b)
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
