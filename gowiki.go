package main

import (
    "fmt"
    "os"
    "net/http"
    "log"
    "net/url"
    "strings"
    "io/ioutil"
    "encoding/json"
    )

const API = "http://en.wikipedia.org/w/api.php"

type Response struct {
    Query struct {
        Pages map[string]PageID `json:"pages`
    } `json:"query"`
}

type PageID struct {
    Extract string  `json:"extract"`
    ns      float64 `json:"ns"`
    pageid  float64 `json:"pageid"`
    title   string  `json:"title"`
}



func main() {

    if len(os.Args) < 2 {
        fmt.Println("\nGowiki - Wikipedia search in your terminal.\n\nUsage: gowiki '<search string>'\n\nExample: gowiki 'abraham lincoln'")
        os.Exit(1)
    }

    var searchQuery string

    if len(os.Args) > 2 {
        searchQuery = strings.Join(os.Args[1:], "_")
    } else {
        searchQuery = strings.Join(os.Args[1:], "")
    }

    // API URL builder
    u, err := url.Parse(API)
    if err != nil {
        log.Fatal(err)
    }

    parameters := url.Values{}
    parameters.Add("titles", searchQuery)
    parameters.Add("redirects", "")
    parameters.Add("action", "query")
    parameters.Add("prop", "extracts")
    parameters.Add("format", "json")
    parameters.Add("exintro", "")
    parameters.Add("explaintext", "")
    parameters.Add("exsectionformat", "plain")
    parameters.Add("indexpageids", "true")
    u.RawQuery = parameters.Encode()


    // API http request
    res, err := http.Get(u.String())
    if err != nil {
        log.Fatal(err)
    }
    if res.StatusCode != 200 {
        log.Fatal("Unexpected response. (Status code received: ", res.StatusCode)
    }

    // Read JSON response
    jsonResponse, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatalln(err)
    }
    defer res.Body.Close()

    // Unmarshal JSON
    var jsonData []Response
    err = json.Unmarshal([]byte(jsonResponse), &jsonData)
    q := Response{}
    err = json.Unmarshal([]byte(jsonResponse), &q)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print resulting summary
    for _, p := range q.Query.Pages {
        fmt.Println("\nWikipedia: \n\n", strings.Replace(p.Extract, "\n", "\n\n", -1))
        return
    }
}
