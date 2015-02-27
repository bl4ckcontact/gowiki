package main

import (
    "fmt"
    "os"
    "net/http"
    "log"
    "net/url"
    "strings"
//    "encoding/json"
    "io/ioutil"
    "encoding/json")

const API = "http://en.wikipedia.org/w/api.php"

type Response struct {
    Query struct {
        Pages struct {
            pageID struct {
                extract string  `json:"extract"`
                ns      float64 `json:"ns"`
                pageid  float64 `json:"pageid"`
                title   string  `json:"title"`
            } `json:"pageID"`
        } `json:"pages"`
    } `json:"query"`
    Warnings struct {
        Query struct {
            _ string `json:"*"`
        } `json:"query"`
    } `json:"warnings"`
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
    u.RawQuery = parameters.Encode()

    // API http request
    res, err := http.Get(u.String())
    if err != nil {
        log.Fatal(err)
    }

    jsonResponse, err := ioutil.ReadAll(res.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var jsonData []Response

    err = json.Unmarshal([]byte(jsonResponse), &jsonData)

    var r map[string]interface{}
    json.Unmarshal([]byte(jsonResponse), &r)
    fmt.Println(r["query"].(map[string]interface{})["pages"])

    if res.StatusCode != 200 {
        log.Fatal("Unexpected response. (Status code received: ", res.StatusCode)
    }
}
